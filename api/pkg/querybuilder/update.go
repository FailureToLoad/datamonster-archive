package querybuilder

import (
	"fmt"
	"sort"
	"strings"
)

type UpdateQuery struct {
	table     Table
	sets      map[string]interface{}
	where     *WhereClause
	returning []string
}

func Update(table Table) *UpdateQuery {
	return &UpdateQuery{
		table: table,
		sets:  make(map[string]interface{}),
	}
}

func (q *UpdateQuery) Set(column string, value interface{}) *UpdateQuery {
	q.sets[column] = value
	return q
}

func (q *UpdateQuery) Where(w WhereClause) *UpdateQuery {
	q.where = &w
	return q
}

func (q *UpdateQuery) Returning(cols ...string) *UpdateQuery {
	q.returning = cols
	return q
}

func (q *UpdateQuery) Build() (string, []interface{}) {
	if len(q.sets) == 0 {
		return "", nil
	}

	var params []interface{}
	query := &strings.Builder{}

	query.WriteString("UPDATE ")
	query.WriteString(q.table.TableName())
	query.WriteString(" SET ")

	setStrings := make([]string, 0, len(q.sets))
	paramCounter := 1

	columns := make([]string, 0, len(q.sets))
	for col := range q.sets {
		columns = append(columns, col)
	}
	sort.Strings(columns)

	for _, col := range columns {
		setStrings = append(setStrings, fmt.Sprintf("%s = $%d", col, paramCounter))
		params = append(params, q.sets[col])
		paramCounter++
	}
	query.WriteString(strings.Join(setStrings, ", "))

	if q.where != nil {
		if whereClause, whereArgs := q.where.Build(); whereClause != "" {
			// Adjust the parameter numbers in the where clause
			adjustedWhereClause, adjustedArgs := adjustWhereClauseParams(whereClause, whereArgs, paramCounter)
			query.WriteString(" ")
			query.WriteString(adjustedWhereClause)
			params = append(params, adjustedArgs...)
		}
	}

	if len(q.returning) > 0 {
		query.WriteString(" RETURNING ")
		query.WriteString(strings.Join(q.returning, ", "))
	}

	return query.String(), params
}

func adjustWhereClauseParams(whereClause string, args []interface{}, startFrom int) (string, []interface{}) {
	for i := len(args); i > 0; i-- {
		oldParam := fmt.Sprintf("$%d", i)
		newParam := fmt.Sprintf("$%d", i+startFrom-1)
		whereClause = strings.Replace(whereClause, oldParam, newParam, -1)
	}
	return whereClause, args
}

package querybuilder

import (
	"strings"
)

type DeleteCommand struct {
	table     Table
	where     *WhereClause
	returning []string
}

func Delete(table Table) *DeleteCommand {
	return &DeleteCommand{
		table: table,
	}
}

func (q *DeleteCommand) Where(w WhereClause) *DeleteCommand {
	q.where = &w
	return q
}

func (q *DeleteCommand) Returning(cols ...string) *DeleteCommand {
	q.returning = cols
	return q
}

// Build constructs the final SQL query and parameters
func (q *DeleteCommand) Build() (string, []interface{}) {
	var params []interface{}
	query := &strings.Builder{}

	query.WriteString("DELETE FROM ")
	query.WriteString(q.table.TableName())

	if q.where != nil {
		if whereClause, whereArgs := q.where.Build(); whereClause != "" {
			query.WriteString(" ")
			query.WriteString(whereClause)
			params = append(params, whereArgs...)
		}
	}

	if len(q.returning) > 0 {
		query.WriteString(" RETURNING ")
		query.WriteString(strings.Join(q.returning, ", "))
	}

	return query.String(), params
}

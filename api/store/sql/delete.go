package sql

import (
	"strings"
)

type DeleteCommand struct {
	TableName string
	where     *WhereClause
	returning []string
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
	builder := &strings.Builder{}

	builder.WriteString("DELETE FROM ")
	builder.WriteString(q.TableName)

	if q.where != nil {
		if whereClause, whereArgs := q.where.Build(); whereClause != "" {
			builder.WriteString(" ")
			builder.WriteString(whereClause)
			params = append(params, whereArgs...)
		}
	}

	if len(q.returning) > 0 {
		builder.WriteString(" RETURNING ")
		builder.WriteString(strings.Join(q.returning, ", "))
	}

	return builder.String(), params
}

package sql

import (
	"fmt"
	"strings"
)

type InsertCommand struct {
	TableName string
	columns   []string
	values    []interface{}
	returning []string
}

func (q *InsertCommand) Columns(cols ...string) *InsertCommand {
	q.columns = cols
	return q
}

func (q *InsertCommand) Values(values ...interface{}) *InsertCommand {
	q.values = values
	return q
}

func (q *InsertCommand) Returning(cols ...string) *InsertCommand {
	q.returning = cols
	return q
}

func (q *InsertCommand) Build() (string, []interface{}) {
	if len(q.columns) == 0 || len(q.values) == 0 {
		return "", nil
	}

	if len(q.columns) != len(q.values) {
		return "", nil
	}

	query := &strings.Builder{}

	query.WriteString("INSERT INTO ")
	query.WriteString(q.TableName)
	query.WriteString(" (")
	query.WriteString(strings.Join(q.columns, ", "))
	query.WriteString(")")

	placeholders := make([]string, len(q.values))
	for i := range q.values {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query.WriteString(" VALUES (")
	query.WriteString(strings.Join(placeholders, ", "))
	query.WriteString(")")

	if len(q.returning) > 0 {
		query.WriteString(" RETURNING ")
		query.WriteString(strings.Join(q.returning, ", "))
	}

	return query.String(), q.values
}

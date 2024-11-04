package columns_test

import (
	columns2 "github.com/failuretoload/datamonster/store/sql/columns"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumericColumn(t *testing.T) {
	tests := []struct {
		name      string
		column    columns2.Integer
		operation func(columns2.Integer) columns2.Condition
		want      string
		args      []interface{}
		paramEnd  int
	}{
		{
			name:   "equals operation",
			column: columns2.NewIntegerColumn("age"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.Equals(int32(25))
			},
			want:     "age = $1",
			args:     []interface{}{int32(25)},
			paramEnd: 2,
		},
		{
			name:   "greater than operation",
			column: columns2.NewIntegerColumn("count"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.GreaterThan(int32(10))
			},
			want:     "count > $1",
			args:     []interface{}{int32(10)},
			paramEnd: 2,
		},
		{
			name:   "less than operation",
			column: columns2.NewIntegerColumn("quantity"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.LessThan(int32(50))
			},
			want:     "quantity < $1",
			args:     []interface{}{int32(50)},
			paramEnd: 2,
		},
		{
			name:   "greater than or equal operation",
			column: columns2.NewIntegerColumn("price"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.GreaterOrEqual(int32(100))
			},
			want:     "price >= $1",
			args:     []interface{}{int32(100)},
			paramEnd: 2,
		},
		{
			name:   "less than or equal operation",
			column: columns2.NewIntegerColumn("stock"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.LessOrEqual(int32(1000))
			},
			want:     "stock <= $1",
			args:     []interface{}{int32(1000)},
			paramEnd: 2,
		},
		{
			name:   "in operation single value",
			column: columns2.NewIntegerColumn("id"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.In(int32(1))
			},
			want:     "id = ANY($1)",
			args:     []interface{}{[]int32{1}},
			paramEnd: 2,
		},
		{
			name:   "in operation multiple values",
			column: columns2.NewIntegerColumn("id"),
			operation: func(c columns2.Integer) columns2.Condition {
				return c.In(int32(1), int32(2), int32(3))
			},
			want:     "id = ANY($1)",
			args:     []interface{}{[]int32{1, 2, 3}},
			paramEnd: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args, paramEnd := tt.operation(tt.column).Build(1)
			assert.Equal(t, tt.want, sql, "SQL should match")
			assert.Equal(t, tt.args, args, "arguments should match")
			assert.Equal(t, tt.paramEnd, paramEnd, "param end should match")
		})
	}
}

func TestLexicColumn(t *testing.T) {
	tests := []struct {
		name      string
		column    columns2.Text
		operation func(columns2.Text) columns2.Condition
		want      string
		args      []interface{}
		paramEnd  int
	}{
		{
			name:   "equals operation",
			column: columns2.NewTextColumn("name"),
			operation: func(c columns2.Text) columns2.Condition {
				return c.Equals("John")
			},
			want:     "name = $1",
			args:     []interface{}{"John"},
			paramEnd: 2,
		},
		{
			name:   "case sensitive LIKE operation",
			column: columns2.NewTextColumn("email"),
			operation: func(c columns2.Text) columns2.Condition {
				return c.Like("%@example.com", false)
			},
			want:     "email LIKE $1",
			args:     []interface{}{"%@example.com"},
			paramEnd: 2,
		},
		{
			name:   "case insensitive ILIKE operation",
			column: columns2.NewTextColumn("title"),
			operation: func(c columns2.Text) columns2.Condition {
				return c.Like("%manager%", true)
			},
			want:     "title ILIKE $1",
			args:     []interface{}{"%manager%"},
			paramEnd: 2,
		},
		{
			name:   "IN operation with single value",
			column: columns2.NewTextColumn("status"),
			operation: func(c columns2.Text) columns2.Condition {
				return c.In("active")
			},
			want:     "status = ANY($1)",
			args:     []interface{}{[]string{"active"}},
			paramEnd: 2,
		},
		{
			name:   "IN operation with multiple values",
			column: columns2.NewTextColumn("category"),
			operation: func(c columns2.Text) columns2.Condition {
				return c.In("books", "movies", "games")
			},
			want:     "category = ANY($1)",
			args:     []interface{}{[]string{"books", "movies", "games"}},
			paramEnd: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args, paramEnd := tt.operation(tt.column).Build(1)
			assert.Equal(t, tt.want, sql, "SQL should match")
			assert.Equal(t, tt.args, args, "arguments should match")
			assert.Equal(t, tt.paramEnd, paramEnd, "param end should match")
		})
	}
}

func TestComposite(t *testing.T) {
	age := columns2.NewIntegerColumn("age")
	name := columns2.NewTextColumn("name")

	tests := []struct {
		name      string
		condition columns2.Condition
		want      string
		args      []interface{}
		paramEnd  int
	}{
		{
			name: "single AND condition",
			condition: columns2.And(
				age.GreaterThan(int32(20)),
			),
			want:     "age > $1",
			args:     []interface{}{int32(20)},
			paramEnd: 2,
		},
		{
			name: "multiple AND Conditions",
			condition: columns2.And(
				age.GreaterThan(int32(20)),
				age.LessThan(int32(30)),
			),
			want:     "(age > $1 AND age < $2)",
			args:     []interface{}{int32(20), int32(30)},
			paramEnd: 3,
		},
		{
			name: "single OR condition",
			condition: columns2.Or(
				name.Equals("John"),
			),
			want:     "name = $1",
			args:     []interface{}{"John"},
			paramEnd: 2,
		},
		{
			name: "multiple OR Conditions",
			condition: columns2.Or(
				name.Equals("John"),
				name.Equals("Jane"),
			),
			want:     "(name = $1 OR name = $2)",
			args:     []interface{}{"John", "Jane"},
			paramEnd: 3,
		},
		{
			name: "nested AND/OR Conditions",
			condition: columns2.And(
				age.GreaterThan(int32(20)),
				columns2.Or(
					name.Equals("John"),
					name.Equals("Jane"),
				),
			),
			want:     "(age > $1 AND (name = $2 OR name = $3))",
			args:     []interface{}{int32(20), "John", "Jane"},
			paramEnd: 4,
		},
		{
			name:      "empty composite",
			condition: columns2.And(),
			want:      "",
			args:      nil,
			paramEnd:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args, paramEnd := tt.condition.Build(1)
			assert.Equal(t, tt.want, sql, "SQL should match")
			assert.Equal(t, tt.args, args, "arguments should match")
			assert.Equal(t, tt.paramEnd, paramEnd, "param end should match")
		})
	}
}

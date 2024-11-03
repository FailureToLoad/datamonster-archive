package querybuilder_test

import (
	"testing"

	"github.com/failuretoload/datamonster/pkg/querybuilder"
	"github.com/failuretoload/datamonster/pkg/querybuilder/columns"
	"github.com/stretchr/testify/assert"
)

type mockTable struct {
	ID    columns.Integer
	Name  columns.Text
	Value columns.Integer
}

func newMockTable() mockTable {
	return mockTable{
		ID:    columns.NewIntegerColumn("id"),
		Name:  columns.NewTextColumn("name"),
		Value: columns.NewIntegerColumn("value"),
	}
}

func (m mockTable) TableName() string {
	return "test_table"
}

func (m mockTable) AllColumns() []string {
	return []string{"id", "name", "value"}
}

func TestSelect(t *testing.T) {
	table := newMockTable()

	t.Run("select all columns", func(t *testing.T) {
		sql, args := querybuilder.Select(table).Build()
		expected := "SELECT id, name, value FROM test_table"
		assert.Equal(t, expected, sql)
		assert.Empty(t, args)
	})

	t.Run("select specific columns", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name", "value").Build()
		assert.Equal(t, "SELECT name, value FROM test_table", sql)
		assert.Empty(t, args)
	})

	t.Run("select with where", func(t *testing.T) {
		sql, args := querybuilder.Select(table).
			Where(querybuilder.Where(table.Name.Equals("test"))).
			Build()
		expected := "SELECT id, name, value FROM test_table WHERE name = $1"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test"}, args)
	})

	t.Run("select with complex where", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name", "value").
			Where(querybuilder.Where(columns.And(
				table.Name.Equals("test"),
				table.Value.GreaterThan(int32(10)),
			))).
			OrderBy("value DESC").
			Limit(10).
			Build()

		expected := "SELECT name, value FROM test_table WHERE (name = $1 AND value > $2) ORDER BY value DESC LIMIT 10"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("select with order by", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name").
			OrderBy("name ASC").
			Build()
		assert.Equal(t, "SELECT name FROM test_table ORDER BY name ASC", sql)
		assert.Empty(t, args)
	})

	t.Run("select with multiple order by", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name").
			OrderBy("value DESC", "name ASC").
			Build()
		assert.Equal(t, "SELECT name FROM test_table ORDER BY value DESC, name ASC", sql)
		assert.Empty(t, args)
	})

	t.Run("select with limit and offset", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name").
			Limit(10).
			Offset(5).
			Build()
		assert.Equal(t, "SELECT name FROM test_table LIMIT 10 OFFSET 5", sql)
		assert.Empty(t, args)
	})

	t.Run("select with everything", func(t *testing.T) {
		sql, args := querybuilder.Select(table, "name", "value").
			Where(querybuilder.Where(columns.And(
				table.Name.Equals("test"),
				columns.Or(
					table.Value.GreaterThan(int32(10)),
					table.Value.LessThan(int32(5)),
				),
			))).
			OrderBy("value DESC", "name ASC").
			Limit(10).
			Offset(5).
			Build()

		expected := "SELECT name, value FROM test_table " +
			"WHERE (name = $1 AND (value > $2 OR value < $3)) " +
			"ORDER BY value DESC, name ASC LIMIT 10 OFFSET 5"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10), int32(5)}, args)
	})
}

func TestWhere(t *testing.T) {
	table := newMockTable()

	t.Run("single equals condition", func(t *testing.T) {
		sql, args := querybuilder.Where(table.Name.Equals("test")).Build()
		assert.Equal(t, "WHERE name = $1", sql)
		assert.Equal(t, []interface{}{"test"}, args)
	})

	t.Run("single comparison conditions", func(t *testing.T) {
		cases := []struct {
			name     string
			where    querybuilder.WhereClause
			wantSQL  string
			wantArgs []interface{}
		}{
			{
				name:     "greater than",
				where:    querybuilder.Where(table.Value.GreaterThan(int32(5))),
				wantSQL:  "WHERE value > $1",
				wantArgs: []interface{}{int32(5)},
			},
			{
				name:     "less than",
				where:    querybuilder.Where(table.Value.LessThan(int32(5))),
				wantSQL:  "WHERE value < $1",
				wantArgs: []interface{}{int32(5)},
			},
			{
				name:     "greater or equal",
				where:    querybuilder.Where(table.Value.GreaterOrEqual(int32(5))),
				wantSQL:  "WHERE value >= $1",
				wantArgs: []interface{}{int32(5)},
			},
			{
				name:     "less or equal",
				where:    querybuilder.Where(table.Value.LessOrEqual(int32(5))),
				wantSQL:  "WHERE value <= $1",
				wantArgs: []interface{}{int32(5)},
			},
		}

		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				sql, args := tc.where.Build()
				assert.Equal(t, tc.wantSQL, sql)
				assert.Equal(t, tc.wantArgs, args)
			})
		}
	})

	t.Run("AND conditions", func(t *testing.T) {
		sql, args := querybuilder.Where(columns.And(
			table.Name.Equals("test"),
			table.Value.GreaterThan(int32(5)),
		)).Build()
		assert.Equal(t, "WHERE (name = $1 AND value > $2)", sql)
		assert.Equal(t, []interface{}{"test", int32(5)}, args)
	})

	t.Run("OR conditions", func(t *testing.T) {
		sql, args := querybuilder.Where(columns.Or(
			table.Name.Equals("test1"),
			table.Name.Equals("test2"),
		)).Build()
		assert.Equal(t, "WHERE (name = $1 OR name = $2)", sql)
		assert.Equal(t, []interface{}{"test1", "test2"}, args)
	})

	t.Run("LIKE conditions", func(t *testing.T) {
		t.Run("case sensitive", func(t *testing.T) {
			sql, args := querybuilder.Where(table.Name.Like("%test%", false)).Build()
			assert.Equal(t, "WHERE name LIKE $1", sql)
			assert.Equal(t, []interface{}{"%test%"}, args)
		})

		t.Run("case insensitive", func(t *testing.T) {
			sql, args := querybuilder.Where(table.Name.Like("%test%", true)).Build()
			assert.Equal(t, "WHERE name ILIKE $1", sql)
			assert.Equal(t, []interface{}{"%test%"}, args)
		})
	})

	t.Run("IN conditions", func(t *testing.T) {
		t.Run("multiple values", func(t *testing.T) {
			sql, args := querybuilder.Where(table.Name.In("test1", "test2", "test3")).Build()
			assert.Equal(t, "WHERE name = ANY($1)", sql)
			assert.Equal(t, []interface{}{[]string{"test1", "test2", "test3"}}, args)
		})

		t.Run("empty IN clause", func(t *testing.T) {
			sql, args := querybuilder.Where(table.Name.In()).Build()
			assert.Equal(t, "WHERE name = ANY($1)", sql)
			assert.Equal(t, []interface{}{[]string(nil)}, args)
		})
	})

	t.Run("edge cases", func(t *testing.T) {
		t.Run("nil condition", func(t *testing.T) {
			sql, args := querybuilder.Where(nil).Build()
			assert.Empty(t, sql)
			assert.Empty(t, args)
		})

		t.Run("empty AND", func(t *testing.T) {
			sql, args := querybuilder.Where(columns.And()).Build()
			assert.Empty(t, sql)
			assert.Empty(t, args)
		})

		t.Run("empty OR", func(t *testing.T) {
			sql, args := querybuilder.Where(columns.Or()).Build()
			assert.Empty(t, sql)
			assert.Empty(t, args)
		})
	})
}

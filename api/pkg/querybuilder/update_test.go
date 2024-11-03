package querybuilder_test

import (
	"testing"

	"github.com/failuretoload/datamonster/pkg/querybuilder"
	"github.com/failuretoload/datamonster/pkg/querybuilder/columns"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	table := newMockTable()

	t.Run("basic update", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "test").
			Build()

		expected := "UPDATE test_table SET name = $1"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test"}, args)
	})

	t.Run("update multiple columns", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "test").
			Set("value", int32(10)).
			Build()

		expected := "UPDATE test_table SET name = $1, value = $2"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("update with where clause", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "test").
			Where(querybuilder.Where(table.ID.Equals(int32(1)))).
			Build()

		expected := "UPDATE test_table SET name = $1 WHERE id = $2"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(1)}, args)
	})

	t.Run("update with complex where clause", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "new_name").
			Where(querybuilder.Where(columns.And(
				table.Value.GreaterThan(int32(10)),
				table.Name.Equals("old_name"),
			))).
			Build()

		expected := "UPDATE test_table SET name = $1 WHERE (value > $2 AND name = $3)"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"new_name", int32(10), "old_name"}, args)
	})

	t.Run("update with returning", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "test").
			Where(querybuilder.Where(table.ID.Equals(int32(1)))).
			Returning("id", "name").
			Build()

		expected := "UPDATE test_table SET name = $1 WHERE id = $2 RETURNING id, name"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(1)}, args)
	})

	t.Run("update with no set clauses", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Where(querybuilder.Where(table.ID.Equals(int32(1)))).
			Build()

		assert.Empty(t, sql)
		assert.Empty(t, args)
	})

	t.Run("update same column multiple times", func(t *testing.T) {
		sql, args := querybuilder.Update(table).
			Set("name", "test1").
			Set("name", "test2"). // This should override the previous value
			Build()

		expected := "UPDATE test_table SET name = $1"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test2"}, args)
	})
}

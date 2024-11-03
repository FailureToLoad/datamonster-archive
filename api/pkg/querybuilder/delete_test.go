package querybuilder_test

import (
	"testing"

	"github.com/failuretoload/datamonster/pkg/querybuilder"
	"github.com/failuretoload/datamonster/pkg/querybuilder/columns"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	table := newMockTable()

	t.Run("basic delete", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).Build()
		expected := "DELETE FROM test_table"
		assert.Equal(t, expected, sql)
		assert.Empty(t, args)
	})

	t.Run("delete with where clause", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).
			Where(querybuilder.Where(table.ID.Equals(int32(1)))).
			Build()

		expected := "DELETE FROM test_table WHERE id = $1"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{int32(1)}, args)
	})

	t.Run("delete with complex where clause", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).
			Where(querybuilder.Where(columns.And(
				table.Name.Equals("test"),
				table.Value.GreaterThan(int32(10)),
			))).
			Build()

		expected := "DELETE FROM test_table WHERE (name = $1 AND value > $2)"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("delete with returning", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).
			Where(querybuilder.Where(table.ID.Equals(int32(1)))).
			Returning("id", "name").
			Build()

		expected := "DELETE FROM test_table WHERE id = $1 RETURNING id, name"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{int32(1)}, args)
	})

	t.Run("delete with multiple conditions and returning", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).
			Where(querybuilder.Where(columns.Or(
				table.Name.Equals("test1"),
				table.Name.Equals("test2"),
			))).
			Returning("id").
			Build()

		expected := "DELETE FROM test_table WHERE (name = $1 OR name = $2) RETURNING id"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test1", "test2"}, args)
	})

	t.Run("delete with empty where clause", func(t *testing.T) {
		sql, args := querybuilder.Delete(table).
			Where(querybuilder.Where(nil)).
			Build()

		expected := "DELETE FROM test_table"
		assert.Equal(t, expected, sql)
		assert.Empty(t, args)
	})
}

package querybuilder_test

import (
	"testing"

	"github.com/failuretoload/datamonster/pkg/querybuilder"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	table := newMockTable()

	t.Run("basic insert", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).
			Columns("name", "value").
			Values("test", int32(10)).
			Build()

		expected := "INSERT INTO test_table (name, value) VALUES ($1, $2)"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("insert with returning", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).
			Columns("name", "value").
			Values("test", int32(10)).
			Returning("id").
			Build()

		expected := "INSERT INTO test_table (name, value) VALUES ($1, $2) RETURNING id"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("insert with multiple returning columns", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).
			Columns("name", "value").
			Values("test", int32(10)).
			Returning("id", "name").
			Build()

		expected := "INSERT INTO test_table (name, value) VALUES ($1, $2) RETURNING id, name"
		assert.Equal(t, expected, sql)
		assert.Equal(t, []interface{}{"test", int32(10)}, args)
	})

	t.Run("insert with no columns", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).Build()
		assert.Empty(t, sql)
		assert.Empty(t, args)
	})

	t.Run("insert with mismatched columns and values", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).
			Columns("name", "value").
			Values("test").
			Build()

		assert.Empty(t, sql)
		assert.Empty(t, args)
	})

	t.Run("insert with no values", func(t *testing.T) {
		sql, args := querybuilder.Insert(table).
			Columns("name", "value").
			Build()

		assert.Empty(t, sql)
		assert.Empty(t, args)
	})
}

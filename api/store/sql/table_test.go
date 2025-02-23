package sql_test

import (
	"testing"

	"github.com/failuretoload/datamonster/store/sql"
	"github.com/failuretoload/datamonster/store/sql/columns"
	"github.com/stretchr/testify/assert"
)

func TestTable_Select(t *testing.T) {
	table := sql.Table{
		Name: "users",
	}

	tests := []struct {
		name           string
		setupQuery     func() *sql.SelectQuery
		expectedSQL    string
		expectedParams []interface{}
	}{
		{
			name: "select all columns",
			setupQuery: func() *sql.SelectQuery {
				return table.Select()
			},
			expectedSQL:    "SELECT * FROM users",
			expectedParams: nil,
		},
		{
			name: "select specific columns",
			setupQuery: func() *sql.SelectQuery {
				return table.Select("name", "age")
			},
			expectedSQL:    "SELECT name, age FROM users",
			expectedParams: nil,
		},
		{
			name: "select with where clause",
			setupQuery: func() *sql.SelectQuery {
				idCol := columns.NewIntegerColumn("id")
				return table.Select("id", "name", "age").Where(sql.Where(idCol.Equals(1)))
			},
			expectedSQL:    "SELECT id, name, age FROM users WHERE id = $1",
			expectedParams: []interface{}{int32(1)},
		},
		{
			name: "select with complex where clause",
			setupQuery: func() *sql.SelectQuery {
				ageCol := columns.NewIntegerColumn("age")
				nameCol := columns.NewTextColumn("name")
				where := sql.Where(columns.And(
					ageCol.GreaterThan(18),
					nameCol.Like("John%", false),
				))
				return table.Select("id", "name", "age").Where(where)
			},
			expectedSQL:    "SELECT id, name, age FROM users WHERE (age > $1 AND name LIKE $2)",
			expectedParams: []interface{}{int32(18), "John%"},
		},
		{
			name: "select with order by",
			setupQuery: func() *sql.SelectQuery {
				return table.Select("id", "name", "age").OrderBy("name", "age DESC")
			},
			expectedSQL:    "SELECT id, name, age FROM users ORDER BY name, age DESC",
			expectedParams: nil,
		},
		{
			name: "select with limit and offset",
			setupQuery: func() *sql.SelectQuery {
				return table.Select("id", "name", "age").Limit(10).Offset(20)
			},
			expectedSQL:    "SELECT id, name, age FROM users LIMIT 10 OFFSET 20",
			expectedParams: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := tt.setupQuery()
			actualSQL, args := query.Build()
			assert.Equal(t, tt.expectedSQL, actualSQL)
			assert.Equal(t, tt.expectedParams, args)
		})
	}
}

func TestTable_Insert(t *testing.T) {
	table := sql.Table{
		Name: "users",
	}

	tests := []struct {
		name           string
		setupQuery     func() (*sql.InsertCommand, []interface{})
		expectedSQL    string
		expectedParams []interface{}
	}{
		{
			name: "basic insert",
			setupQuery: func() (*sql.InsertCommand, []interface{}) {
				return table.Insert().
					Columns("name", "age").
					Values("John Doe", 30), nil
			},
			expectedSQL:    "INSERT INTO users (name, age) VALUES ($1, $2)",
			expectedParams: []interface{}{"John Doe", 30},
		},
		{
			name: "insert with returning",
			setupQuery: func() (*sql.InsertCommand, []interface{}) {
				return table.Insert().
					Columns("name", "age").
					Values("John Doe", 30).
					Returning("id"), nil
			},
			expectedSQL:    "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id",
			expectedParams: []interface{}{"John Doe", 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, _ := tt.setupQuery()
			actualSQL, args := query.Build()
			assert.Equal(t, tt.expectedSQL, actualSQL)
			assert.Equal(t, tt.expectedParams, args)
		})
	}
}

func TestTable_Update(t *testing.T) {
	table := sql.Table{
		Name: "users",
	}

	tests := []struct {
		name           string
		setupQuery     func() (*sql.UpdateQuery, []interface{})
		expectedSQL    string
		expectedParams []interface{}
	}{
		{
			name: "basic update",
			setupQuery: func() (*sql.UpdateQuery, []interface{}) {
				return table.Update().
					Set("name", "Jane Doe").
					Set("age", 31), nil
			},
			expectedSQL:    "UPDATE users SET age = $1, name = $2",
			expectedParams: []interface{}{31, "Jane Doe"},
		},
		{
			name: "update with where clause",
			setupQuery: func() (*sql.UpdateQuery, []interface{}) {
				idCol := columns.NewIntegerColumn("id")
				return table.Update().
					Set("name", "Jane Doe").
					Set("age", 31).
					Where(sql.Where(idCol.Equals(1))), nil
			},
			expectedSQL:    "UPDATE users SET age = $1, name = $2 WHERE id = $3",
			expectedParams: []interface{}{31, "Jane Doe", int32(1)},
		},
		{
			name: "update with returning",
			setupQuery: func() (*sql.UpdateQuery, []interface{}) {
				return table.Update().
					Set("name", "Jane Doe").
					Set("age", 31).
					Returning("id", "name"), nil
			},
			expectedSQL:    "UPDATE users SET age = $1, name = $2 RETURNING id, name",
			expectedParams: []interface{}{31, "Jane Doe"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, _ := tt.setupQuery()
			actualSQL, args := query.Build()
			assert.Equal(t, tt.expectedSQL, actualSQL)
			assert.Equal(t, tt.expectedParams, args)
		})
	}
}

func TestTable_Delete(t *testing.T) {
	table := sql.Table{
		Name: "users",
	}

	tests := []struct {
		name           string
		setupQuery     func() (*sql.DeleteCommand, []interface{})
		expectedSQL    string
		expectedParams []interface{}
	}{
		{
			name: "delete all",
			setupQuery: func() (*sql.DeleteCommand, []interface{}) {
				return table.Delete(), nil
			},
			expectedSQL:    "DELETE FROM users",
			expectedParams: nil,
		},
		{
			name: "delete with where clause",
			setupQuery: func() (*sql.DeleteCommand, []interface{}) {
				idCol := columns.NewIntegerColumn("id")
				return table.Delete().Where(sql.Where(idCol.Equals(1))), nil
			},
			expectedSQL:    "DELETE FROM users WHERE id = $1",
			expectedParams: []interface{}{int32(1)},
		},
		{
			name: "delete with returning",
			setupQuery: func() (*sql.DeleteCommand, []interface{}) {
				idCol := columns.NewIntegerColumn("id")
				return table.Delete().
					Where(sql.Where(idCol.Equals(1))).
					Returning("id", "name"), nil
			},
			expectedSQL:    "DELETE FROM users WHERE id = $1 RETURNING id, name",
			expectedParams: []interface{}{int32(1)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query, _ := tt.setupQuery()
			actualSQL, args := query.Build()
			assert.Equal(t, tt.expectedSQL, actualSQL)
			assert.Equal(t, tt.expectedParams, args)
		})
	}
}

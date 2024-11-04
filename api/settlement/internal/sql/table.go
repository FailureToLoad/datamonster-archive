package sql

import (
	"github.com/failuretoload/datamonster/store/sql"
)

func PostGres() sql.Table {
	tableName := "campaign.settlement"
	return sql.Table{Name: tableName, ColumnNames: columnNames}
}

func SQLite() sql.Table {
	tableName := "settlement"
	return sql.Table{Name: tableName, ColumnNames: columnNames}
}

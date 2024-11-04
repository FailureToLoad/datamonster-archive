package sql

import (
	"github.com/failuretoload/datamonster/store/sql"
)

func PostGres() sql.Table {
	return sql.Table{Name: "campaign.survivor"}
}

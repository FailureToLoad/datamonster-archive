package sql

import (
	"github.com/failuretoload/datamonster/store/sql"
	coldefs "github.com/failuretoload/datamonster/store/sql/columns"
)

func PostGres() sql.Table {
	return sql.Table{Name: "campaign.settlement"}
}

type SettlementColumns struct {
	ID                coldefs.Integer
	Owner             coldefs.Text
	Name              coldefs.Text
	SurvivalLimit     coldefs.Integer
	DepartingSurvival coldefs.Integer
	CollectiveCog     coldefs.Integer
	Year              coldefs.Integer
}

const (
	ID                = "id"
	Owner             = "owner"
	Name              = "name"
	SurvivalLimit     = "survival_limit"
	DepartingSurvival = "departing_survival"
	CC                = "collective_cognition"
	Year              = "year"
)

var (
	Columns = SettlementColumns{
		ID:                coldefs.NewIntegerColumn(ID),
		Owner:             coldefs.NewTextColumn(Owner),
		Name:              coldefs.NewTextColumn(Name),
		SurvivalLimit:     coldefs.NewIntegerColumn(SurvivalLimit),
		DepartingSurvival: coldefs.NewIntegerColumn(DepartingSurvival),
		CollectiveCog:     coldefs.NewIntegerColumn(CC),
		Year:              coldefs.NewIntegerColumn(Year),
	}
)

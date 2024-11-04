package sql

import coldefs "github.com/failuretoload/datamonster/store/sql/columns"

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
	ID                  = "id"
	Owner               = "owner"
	NameColumn          = "name"
	SurvivalLimitColumn = "survival_limit"
	DepartingSurvival   = "departing_survival"
	CC                  = "collective_cognition"
	Year                = "year"
)

var (
	Columns = SettlementColumns{
		ID:                coldefs.NewIntegerColumn(ID),
		Owner:             coldefs.NewTextColumn(Owner),
		Name:              coldefs.NewTextColumn(NameColumn),
		SurvivalLimit:     coldefs.NewIntegerColumn(SurvivalLimitColumn),
		DepartingSurvival: coldefs.NewIntegerColumn(DepartingSurvival),
		CollectiveCog:     coldefs.NewIntegerColumn(CC),
		Year:              coldefs.NewIntegerColumn(Year),
	}
	columnNames = []string{
		ID,
		Owner,
		NameColumn,
		SurvivalLimitColumn,
		DepartingSurvival,
		CC,
		Year,
	}
)

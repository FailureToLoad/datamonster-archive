package sql

import coldefs "github.com/failuretoload/datamonster/store/sql/columns"

type SurvivorColumns struct {
	ID               coldefs.Integer
	Settlement       coldefs.Integer
	Name             coldefs.Text
	Birth            coldefs.Integer
	Gender           coldefs.Text
	Status           coldefs.Text
	HuntXp           coldefs.Integer
	Survival         coldefs.Integer
	Movement         coldefs.Integer
	Accuracy         coldefs.Integer
	Strength         coldefs.Integer
	Evasion          coldefs.Integer
	Luck             coldefs.Integer
	Speed            coldefs.Integer
	Insanity         coldefs.Integer
	SystemicPressure coldefs.Integer
	Torment          coldefs.Integer
	Lumi             coldefs.Integer
	Courage          coldefs.Integer
	Understanding    coldefs.Integer
}

const (
	ID               = "id"
	Settlement       = "settlement"
	Name             = "name"
	Birth            = "birth"
	Gender           = "gender"
	Status           = "status"
	HuntXp           = "huntxp"
	SurvivalColumn   = "survival"
	Movement         = "movement"
	Accuracy         = "accuracy"
	Strength         = "strength"
	Evasion          = "evasion"
	Luck             = "luck"
	Speed            = "speed"
	Insanity         = "insanity"
	SystemicPressure = "systemic_pressure"
	Torment          = "torment"
	Lumi             = "lumi"
	Courage          = "courage"
	Understanding    = "understanding"
)

var (
	Columns = SurvivorColumns{
		ID:               coldefs.NewIntegerColumn(ID),
		Settlement:       coldefs.NewIntegerColumn(Settlement),
		Name:             coldefs.NewTextColumn(Name),
		Birth:            coldefs.NewIntegerColumn(Birth),
		Gender:           coldefs.NewTextColumn(Gender),
		Status:           coldefs.NewTextColumn(Status),
		HuntXp:           coldefs.NewIntegerColumn(HuntXp),
		Survival:         coldefs.NewIntegerColumn(SurvivalColumn),
		Movement:         coldefs.NewIntegerColumn(Movement),
		Accuracy:         coldefs.NewIntegerColumn(Accuracy),
		Strength:         coldefs.NewIntegerColumn(Strength),
		Evasion:          coldefs.NewIntegerColumn(Evasion),
		Luck:             coldefs.NewIntegerColumn(Luck),
		Speed:            coldefs.NewIntegerColumn(Speed),
		Insanity:         coldefs.NewIntegerColumn(Insanity),
		SystemicPressure: coldefs.NewIntegerColumn(SystemicPressure),
		Torment:          coldefs.NewIntegerColumn(Torment),
		Lumi:             coldefs.NewIntegerColumn(Lumi),
		Courage:          coldefs.NewIntegerColumn(Courage),
		Understanding:    coldefs.NewIntegerColumn(Understanding),
	}
)

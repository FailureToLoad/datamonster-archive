package domain

type Survivor struct {
	ID               int     `db:"id"`
	Settlement       int     `db:"settlement"`
	Name             string  `db:"name"`
	Birth            int     `db:"birth"`
	Gender           string  `db:"gender"`
	Status           *string `db:"status"`
	HuntXp           int     `db:"huntxp"`
	Survival         int     `db:"survival"`
	Movement         int     `db:"movement"`
	Accuracy         int     `db:"accuracy"`
	Strength         int     `db:"strength"`
	Evasion          int     `db:"evasion"`
	Luck             int     `db:"luck"`
	Speed            int     `db:"speed"`
	Insanity         int     `db:"insanity"`
	SystemicPressure int     `db:"systemic_pressure"`
	Torment          int     `db:"torment"`
	Lumi             int     `db:"lumi"`
	Courage          int     `db:"courage"`
	Understanding    int     `db:"understanding"`
}

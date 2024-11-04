package domain

type Settlement struct {
	ID                  int
	Owner               string
	Name                string
	SurvivalLimit       int
	DepartingSurvival   int
	CollectiveCognition int
	CurrentYear         int
}

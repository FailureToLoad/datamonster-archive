package internal

import (
	"context"
	"github.com/failuretoload/datamonster/store"
)

type Repo struct {
	pool store.Connection
}

type Settlement struct {
	ID                  int
	Owner               string
	Name                string
	SurvivalLimit       int
	DepartingSurvival   int
	CollectiveCognition int
	CurrentYear         int
}

func New(d store.Connection) *Repo {
	return &Repo{pool: d}
}

func (r Repo) Select(ctx context.Context, userID string) ([]Settlement, error) {
	query := "SELECT * FROM campaign.settlement WHERE owner=$1"
	rows, queryErr := r.pool.Query(ctx, query, userID)
	if queryErr != nil {
		return []Settlement{}, queryErr
	}
	defer rows.Close()

	settlements := []Settlement{}
	for rows.Next() {
		var s Settlement
		err := rows.Scan(&s.ID, &s.Owner, &s.Name, &s.SurvivalLimit, &s.DepartingSurvival, &s.CollectiveCognition, &s.CurrentYear)
		if err != nil {
			return settlements, err
		}
		settlements = append(settlements, s)
	}
	return settlements, nil
}

func (r Repo) Get(ctx context.Context, id, userID string) (Settlement, error) {
	query := "SELECT * FROM campaign.settlement WHERE id = $1 AND owner = $2 LIMIT 1"
	var s Settlement
	err := r.pool.QueryRow(ctx, query, id, userID).Scan(
		&s.ID,
		&s.Owner,
		&s.Name,
		&s.SurvivalLimit,
		&s.DepartingSurvival,
		&s.CollectiveCognition,
		&s.CurrentYear,
	)
	return s, err
}

func (r Repo) Insert(ctx context.Context, s Settlement) (int, error) {
	query := `
        INSERT INTO campaign.settlement 
            (owner, name, survival_limit, departing_survival, collective_cognition, year)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	var id int
	err := r.pool.QueryRow(ctx, query,
		s.Owner,
		s.Name,
		s.SurvivalLimit,
		s.DepartingSurvival,
		s.CollectiveCognition,
		s.CurrentYear,
	).Scan(&id)

	return id, err
}

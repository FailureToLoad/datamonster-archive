package repo

import (
	"context"
	"github.com/failuretoload/datamonster/settlement/domain"
	"github.com/failuretoload/datamonster/settlement/internal/sql"
	"github.com/failuretoload/datamonster/store"
	builder "github.com/failuretoload/datamonster/store/sql"
	"github.com/failuretoload/datamonster/store/sql/columns"
)

type Repo struct {
	conn  store.Connection
	table builder.Table
}

func New(c store.Connection) *Repo {
	return &Repo{conn: c, table: sql.PostGres()}
}

func (r Repo) Select(ctx context.Context, userID string) ([]domain.Settlement, error) {
	query, args := r.table.Select().
		Where(builder.Where(sql.Columns.Owner.Equals(userID))).
		Build()

	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settlements []domain.Settlement
	for rows.Next() {
		var s SettlementRow
		err := rows.Scan(
			&s.ID,
			&s.Owner,
			&s.Name,
			&s.SurvivalLimit,
			&s.DepartingSurvival,
			&s.CollectiveCognition,
			&s.CurrentYear,
		)
		if err != nil {
			return nil, err
		}
		settlements = append(settlements, s.ToDomain())
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return settlements, nil
}

func (r Repo) Get(ctx context.Context, id int, userID string) (domain.Settlement, error) {
	query, args := r.table.Select().
		Where(builder.Where(
			columns.And(
				sql.Columns.ID.Equals(int32(id)),
				sql.Columns.Owner.Equals(userID),
			),
		)).
		Build()

	var s SettlementRow
	err := r.conn.QueryRow(ctx, query, args...).Scan(
		&s.ID,
		&s.Owner,
		&s.Name,
		&s.SurvivalLimit,
		&s.DepartingSurvival,
		&s.CollectiveCognition,
		&s.CurrentYear,
	)
	return s.ToDomain(), err
}

func (r Repo) Insert(ctx context.Context, s domain.Settlement) (int, error) {
	query, args := r.table.Insert().
		Columns(
			sql.Owner,
			sql.Name,
			sql.SurvivalLimit,
			sql.DepartingSurvival,
			sql.CC,
			sql.Year,
		).
		Values(
			s.Owner,
			s.Name,
			sql.Columns.SurvivalLimit.Convert(s.SurvivalLimit),
			sql.Columns.DepartingSurvival.Convert(s.DepartingSurvival),
			sql.Columns.CollectiveCog.Convert(s.CollectiveCognition),
			sql.Columns.Year.Convert(s.CurrentYear),
		).
		Returning("id").
		Build()

	var id int32
	err := r.conn.QueryRow(ctx, query, args...).Scan(&id)
	return int(id), err
}

type SettlementRow struct {
	ID                  int32
	Owner               string
	Name                string
	SurvivalLimit       int32
	DepartingSurvival   int32
	CollectiveCognition int32
	CurrentYear         int32
}

func (r SettlementRow) ToDomain() domain.Settlement {
	return domain.Settlement{
		ID:                  int(r.ID),
		Owner:               r.Owner,
		Name:                r.Name,
		SurvivalLimit:       int(r.SurvivalLimit),
		DepartingSurvival:   int(r.DepartingSurvival),
		CollectiveCognition: int(r.CollectiveCognition),
		CurrentYear:         int(r.CurrentYear),
	}
}

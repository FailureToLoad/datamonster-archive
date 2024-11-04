package repo

import (
	"context"
	"github.com/failuretoload/datamonster/settlement/domain"
	"github.com/failuretoload/datamonster/settlement/internal/sql"
	"github.com/failuretoload/datamonster/store"
	builder "github.com/failuretoload/datamonster/store/sql"
	"github.com/failuretoload/datamonster/store/sql/operators"
)

type Repo struct {
	dao   store.DAO
	table builder.Table
}

func New(d store.DAO, t builder.Table) *Repo {
	return &Repo{dao: d, table: t}
}

func (r Repo) Select(ctx context.Context, userID string) ([]domain.Settlement, error) {
	query, args := r.table.Select().
		Where(builder.Where(sql.Columns.Owner.Equals(userID))).
		Build()

	rows, err := r.dao.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var settlements []domain.Settlement
	for rows.Next() {
		var s domain.Settlement
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
		settlements = append(settlements, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return settlements, nil
}

func (r Repo) Get(ctx context.Context, id int32, userID string) (domain.Settlement, error) {
	query, args := r.table.Select().
		Where(builder.Where(
			operators.And(
				sql.Columns.ID.Equals(id),
				sql.Columns.Owner.Equals(userID),
			),
		)).
		Build()

	var s domain.Settlement
	err := r.dao.QueryRow(ctx, query, args...).Scan(
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

func (r Repo) Insert(ctx context.Context, s domain.Settlement) (int32, error) {
	query, args := r.table.Insert().
		Columns(
			sql.Owner,
			sql.NameColumn,
			sql.SurvivalLimitColumn,
			sql.DepartingSurvival,
			sql.CC,
			sql.Year,
		).
		Values(
			s.Owner,
			s.Name,
			s.SurvivalLimit,
			s.DepartingSurvival,
			s.CollectiveCognition,
			s.CurrentYear,
		).
		Returning("id").
		Build()

	var id int32
	err := r.dao.QueryRow(ctx, query, args...).Scan(&id)
	return id, err
}

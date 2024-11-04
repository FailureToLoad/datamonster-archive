package repo

import (
	"context"
	"fmt"
	"github.com/failuretoload/datamonster/store"
	builder "github.com/failuretoload/datamonster/store/sql"
	"github.com/failuretoload/datamonster/survivor/domain"
	"github.com/failuretoload/datamonster/survivor/internal/sql"
	"log"
	"strings"
)

type Repo struct {
	conn  store.Connection
	table builder.Table
}

func New(c store.Connection) *Repo {
	return &Repo{conn: c, table: sql.PostGres()}
}

func (r Repo) CreateSurvivor(ctx context.Context, d domain.Survivor) error {
	s := FromDomain(d)
	query, args := r.table.Insert().
		Columns(
			sql.Settlement,
			sql.Name,
			sql.Birth,
			sql.Gender,
			sql.HuntXp,
			sql.SurvivalColumn,
			sql.Movement,
			sql.Accuracy,
			sql.Strength,
			sql.Evasion,
			sql.Luck,
			sql.Speed,
			sql.Insanity,
			sql.SystemicPressure,
			sql.Torment,
			sql.Lumi,
			sql.Courage,
			sql.Understanding,
		).
		Values(
			s.Settlement,
			s.Name,
			s.Birth,
			s.Gender,
			s.HuntXp,
			s.Survival,
			s.Movement,
			s.Accuracy,
			s.Strength,
			s.Evasion,
			s.Luck,
			s.Speed,
			s.Insanity,
			s.SystemicPressure,
			s.Torment,
			s.Lumi,
			s.Courage,
			s.Understanding,
		).
		Build()

	_, err := r.conn.Exec(ctx, query, args...)
	if err != nil {
		logString := fmt.Errorf("unable to create survivor in settlement %d: %w", s.Settlement, err)
		log.Default().Println(logString)

		if IsDuplicateKeyError(err) {
			return NewDuplicateNameError(s.Name)
		}
		return err
	}
	return nil
}

func (r Repo) GetAllSurvivorsForSettlement(ctx context.Context, settlementID int) ([]domain.Survivor, error) {
	query, args := r.table.Select().
		Where(builder.Where(sql.Columns.Settlement.Equals(int32(settlementID)))).
		Build()

	rows, err := r.conn.Query(ctx, query, args...)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var survivors []domain.Survivor
	for rows.Next() {
		var s SurvivorRow
		err := rows.Scan(
			&s.ID,
			&s.Settlement,
			&s.Name,
			&s.Birth,
			&s.Gender,
			&s.Status,
			&s.HuntXp,
			&s.Survival,
			&s.Movement,
			&s.Accuracy,
			&s.Strength,
			&s.Evasion,
			&s.Luck,
			&s.Speed,
			&s.Insanity,
			&s.SystemicPressure,
			&s.Torment,
			&s.Lumi,
			&s.Courage,
			&s.Understanding,
		)
		if err != nil {
			log.Default().Println(err.Error())
			return survivors, err
		}
		survivors = append(survivors, s.ToDomain())
	}
	return survivors, rows.Err()
}

func IsDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}

type SurvivorRow struct {
	ID               int32
	Settlement       int32
	Name             string
	Birth            int32
	Gender           string
	Status           *string
	HuntXp           int32
	Survival         int32
	Movement         int32
	Accuracy         int32
	Strength         int32
	Evasion          int32
	Luck             int32
	Speed            int32
	Insanity         int32
	SystemicPressure int32
	Torment          int32
	Lumi             int32
	Courage          int32
	Understanding    int32
}

func (r SurvivorRow) ToDomain() domain.Survivor {
	return domain.Survivor{
		ID:               int(r.ID),
		Settlement:       int(r.Settlement),
		Name:             r.Name,
		Birth:            int(r.Birth),
		Gender:           r.Gender,
		Status:           r.Status,
		HuntXp:           int(r.HuntXp),
		Survival:         int(r.Survival),
		Movement:         int(r.Movement),
		Accuracy:         int(r.Accuracy),
		Strength:         int(r.Strength),
		Evasion:          int(r.Evasion),
		Luck:             int(r.Luck),
		Speed:            int(r.Speed),
		Insanity:         int(r.Insanity),
		SystemicPressure: int(r.SystemicPressure),
		Torment:          int(r.Torment),
		Lumi:             int(r.Lumi),
		Courage:          int(r.Courage),
		Understanding:    int(r.Understanding),
	}
}

func FromDomain(s domain.Survivor) SurvivorRow {
	return SurvivorRow{
		ID:               int32(s.ID),
		Settlement:       int32(s.Settlement),
		Name:             s.Name,
		Birth:            int32(s.Birth),
		Gender:           s.Gender,
		Status:           s.Status,
		HuntXp:           int32(s.HuntXp),
		Survival:         int32(s.Survival),
		Movement:         int32(s.Movement),
		Accuracy:         int32(s.Accuracy),
		Strength:         int32(s.Strength),
		Evasion:          int32(s.Evasion),
		Luck:             int32(s.Luck),
		Speed:            int32(s.Speed),
		Insanity:         int32(s.Insanity),
		SystemicPressure: int32(s.SystemicPressure),
		Torment:          int32(s.Torment),
		Lumi:             int32(s.Lumi),
		Courage:          int32(s.Courage),
		Understanding:    int32(s.Understanding),
	}
}

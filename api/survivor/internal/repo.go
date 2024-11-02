package repo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/failuretoload/datamonster/store"
)

type PGRepo struct {
	pool store.Connection
}

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

func NewRepo(d store.Connection) *PGRepo {
	return &PGRepo{pool: d}
}

func (r PGRepo) CreateSurvivor(ctx context.Context, s Survivor) error {
	query := `
        INSERT INTO campaign.survivor (
            settlement, name, birth, huntxp, gender, survival, 
            movement, accuracy, strength, evasion, luck, speed, 
            insanity, systemic_pressure, torment, lumi, courage, understanding
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`

	_, err := r.pool.Exec(ctx, query,
		s.Settlement,
		s.Name,
		s.Birth,
		s.HuntXp,
		s.Gender,
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
	)

	if err != nil {
		logString := fmt.Errorf("unable to create survivor in settlement %d: %w", s.Settlement, err)
		log.Default().Println(logString)
		if strings.Contains(err.Error(), "duplicate key value") {
			return NewDuplicateNameError(s.Name)
		}
	}
	return err
}

func (r PGRepo) GetAllSurvivorsForSettlement(ctx context.Context, settlementID int) ([]Survivor, error) {
	query := "SELECT * FROM campaign.survivor WHERE settlement = $1"
	survivors, err := r.find(ctx, query, settlementID)
	return survivors, err
}

func (r PGRepo) find(ctx context.Context, query string, args ...interface{}) ([]Survivor, error) {
	log.Default().Println(query)
	rows, queryErr := r.pool.Query(ctx, query, args...)
	if queryErr != nil {
		log.Default().Println(queryErr.Error())
		return nil, queryErr
	}
	defer rows.Close()
	survivors := []Survivor{}
	for rows.Next() {
		var s Survivor
		err := rows.Scan(&s.ID,
			&s.Settlement,
			&s.Name,
			&s.Gender,
			&s.Birth,
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
			&s.Status,
		)
		if err != nil {
			log.Default().Println(err.Error())
			return survivors, err
		}
		survivors = append(survivors, s)
	}
	return survivors, nil
}

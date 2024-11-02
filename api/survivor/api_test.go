package survivor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*Controller, pgxmock.PgxPoolIface, *chi.Mux) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("failed to create mock: %v", err)
	}

	controller := NewController(mock)
	router := chi.NewRouter()
	controller.RegisterRoutes(router)

	return controller, mock, router
}

func TestGetSurvivors_ReturnsSurvivorList(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	survivorCols := []string{
		"id", "settlement", "name", "gender", "birth", "huntxp", "survival",
		"movement", "accuracy", "strength", "evasion", "luck", "speed",
		"insanity", "systemic_pressure", "torment", "lumi", "courage", "understanding",
		"status",
	}

	values := [][]any{
		{1, 1, "Zach", "M", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, nil},
		{2, 1, "Lucy", "M", 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, nil},
	}

	rows := pgxmock.NewRows(survivorCols).AddRows(values...)

	db.ExpectQuery(`SELECT \* FROM campaign\.survivor WHERE settlement = 1`).
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", "/settlements/1/survivors", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode, "200 response should be returned")

	body, _ := io.ReadAll(resp.Body)
	dtoList := []DTO{}
	_ = json.Unmarshal(body, &dtoList)
	assert.Equal(t, 2, len(dtoList), "2 survivors should be returned")
}

func TestCreateSurvivor_ReturnsNoContent(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec(`INSERT INTO campaign\.survivor \(settlement, name, birth, huntxp, gender, survival, 
        movement, accuracy, strength, evasion, luck, speed, insanity, systemic_pressure, torment, lumi, courage, understanding\) 
        VALUES \(1, 'Zach', 1, 1, 'M', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1\)`).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	survivor := DTO{
		Settlement:       1,
		Name:             "Zach",
		Birth:            1,
		Gender:           "M",
		HuntXp:           1,
		Survival:         1,
		Movement:         1,
		Accuracy:         1,
		Strength:         1,
		Evasion:          1,
		Luck:             1,
		Speed:            1,
		Insanity:         1,
		SystemicPressure: 1,
		Torment:          1,
		Lumi:             1,
		Understanding:    1,
		Courage:          1,
	}
	reqBody, err := json.Marshal(survivor)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest("POST", "/settlements/1/survivors", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 204, resp.StatusCode, "204 response should be returned")
}

func TestCreateSurvivor_RequiresAValidSettlementId(t *testing.T) {
	_, _, router := setupTest(t)

	req := httptest.NewRequest("POST", "/settlements/z/survivors", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 500, resp.StatusCode, "500 should be returned if the param is invalid")
}

func TestCreateSurvivor_RequiresAUniqueName(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec(`INSERT INTO campaign\.survivor \(settlement, name, birth, huntxp, gender, survival, 
        movement, accuracy, strength, evasion, luck, speed, insanity, systemic_pressure, torment, lumi, courage, understanding\) 
        VALUES \(1, 'Zach', 1, 1, 'M', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1\)`).
		WillReturnError(fmt.Errorf("duplicate key value"))

	survivor := DTO{
		Settlement:       1,
		Name:             "Zach",
		Birth:            1,
		Gender:           "M",
		HuntXp:           1,
		Survival:         1,
		Movement:         1,
		Accuracy:         1,
		Strength:         1,
		Evasion:          1,
		Luck:             1,
		Speed:            1,
		Insanity:         1,
		SystemicPressure: 1,
		Torment:          1,
		Lumi:             1,
		Understanding:    1,
		Courage:          1,
	}

	reqBody, err := json.Marshal(survivor)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest("POST", "/settlements/1/survivors", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 400, resp.StatusCode, "400 should be returned if the survivor already exists")
}

func TestCreateSurvivor_RequiresAValidBody(t *testing.T) {
	_, _, router := setupTest(t)

	wrongBody := struct {
		a, b, c int
	}{
		a: 1,
		b: 1,
		c: 1,
	}

	reqBody, err := json.Marshal(wrongBody)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest("POST", "/settlements/1/survivors", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 500, resp.StatusCode, "500 should be returned if the body is invalid")
}

func TestCreateSurvivor_CommunicatesDbIssues(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec(`INSERT INTO campaign\.survivor \(settlement, name, birth, huntxp, gender, survival, 
        movement, accuracy, strength, evasion, luck, speed, insanity, systemic_pressure, torment, lumi, courage, understanding\) 
        VALUES \(1, 'Zach', 1, 1, 'M', 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1\)`).
		WillReturnError(fmt.Errorf("well that ain't right"))

	survivor := DTO{
		Settlement:       1,
		Name:             "Zach",
		Birth:            1,
		Gender:           "M",
		HuntXp:           1,
		Survival:         1,
		Movement:         1,
		Accuracy:         1,
		Strength:         1,
		Evasion:          1,
		Luck:             1,
		Speed:            1,
		Insanity:         1,
		SystemicPressure: 1,
		Torment:          1,
		Lumi:             1,
		Understanding:    1,
		Courage:          1,
	}

	reqBody, err := json.Marshal(survivor)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	req := httptest.NewRequest("POST", "/settlements/1/survivors", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, 500, resp.StatusCode, "500 should be returned as the default for DB issues")
}

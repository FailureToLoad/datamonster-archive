package survivor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/failuretoload/datamonster/request"
	"github.com/go-chi/chi/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testSettlementID = 1
	testUserID       = "userId"
	survivorCols     = []string{
		"id", "settlement", "name", "birth", "gender",
		"status", "huntxp", "survival", "movement",
		"accuracy", "strength", "evasion", "luck",
		"speed", "insanity", "systemic_pressure",
		"torment", "lumi", "courage", "understanding",
	}
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

	values := [][]any{
		{int32(1), int32(testSettlementID), "Survivor One", int32(1), "F", nil, int32(0), int32(1), int32(5), int32(5), int32(5), int32(5), int32(5), int32(5), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0)},
		{int32(2), int32(testSettlementID), "Survivor Two", int32(1), "M", nil, int32(0), int32(1), int32(5), int32(5), int32(5), int32(5), int32(5), int32(5), int32(0), int32(0), int32(0), int32(0), int32(0), int32(0)},
	}
	rows := pgxmock.NewRows(survivorCols).AddRows(values...)
	db.ExpectQuery("SELECT .* FROM campaign.survivor WHERE").
		WithArgs(int32(testSettlementID)).
		WillReturnRows(rows)

	req := httptest.NewRequest("GET", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode, "200 response should be returned")
	body, _ := io.ReadAll(resp.Body)
	var dto []DTO
	_ = json.Unmarshal(body, &dto)
	assert.Equal(t, 2, len(dto), "2 survivors should be returned")
}

func TestGetSurvivors_ReportsScanErrors(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectQuery("SELECT .* FROM campaign.survivor WHERE").
		WithArgs(int32(testSettlementID)).
		WillReturnError(fmt.Errorf("scan error"))

	req := httptest.NewRequest("GET", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 500, resp.StatusCode, "scan errors should result in server error")
}

func TestGetSurvivors_ReportsConnectionErrors(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectQuery("SELECT .* FROM campaign.survivor WHERE").
		WithArgs(int32(testSettlementID)).
		WillReturnError(fmt.Errorf("connection error"))

	req := httptest.NewRequest("GET", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), nil)
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 500, resp.StatusCode, "connection issues should result in server error")
}

func TestCreateSurvivor_Success(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec("INSERT INTO campaign.survivor").
		WithArgs(
			int32(testSettlementID), "New Survivor", int32(1), "F",
			int32(0), int32(1), int32(5), int32(5), int32(5),
			int32(5), int32(5), int32(5), int32(0), int32(0),
			int32(0), int32(0), int32(0), int32(0),
		).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	survivorRequest := DTO{
		Settlement: testSettlementID,
		Name:       "New Survivor",
		Birth:      1,
		Gender:     "F",
		HuntXp:     0,
		Survival:   1,
		Movement:   5,
		Accuracy:   5,
		Strength:   5,
		Evasion:    5,
		Luck:       5,
		Speed:      5,
	}
	reqBody, _ := json.Marshal(survivorRequest)
	req := httptest.NewRequest("POST", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), bytes.NewReader(reqBody))
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 204, resp.StatusCode, "successful creation should return 204")
}

func TestCreateSurvivor_RequiresValidSettlementID(t *testing.T) {
	_, _, router := setupTest(t)

	survivorRequest := DTO{
		Name:   "New Survivor",
		Gender: "F",
	}
	reqBody, _ := json.Marshal(survivorRequest)
	req := httptest.NewRequest("POST", "/settlements/invalid/survivors", bytes.NewReader(reqBody))
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 500, resp.StatusCode, "invalid settlement ID should return 500")
}

func TestCreateSurvivor_RequiresName(t *testing.T) {
	_, _, router := setupTest(t)

	survivorRequest := DTO{
		Settlement: testSettlementID,
		Gender:     "F",
	}
	reqBody, _ := json.Marshal(survivorRequest)
	req := httptest.NewRequest("POST", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), bytes.NewReader(reqBody))
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 500, resp.StatusCode, "missing name should return 500")
}

func TestCreateSurvivor_RequiresUniqueName(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec("INSERT INTO campaign.survivor").
		WithArgs(
			int32(testSettlementID), "Duplicate Name", int32(1), "F",
			int32(0), int32(1), int32(5), int32(5), int32(5),
			int32(5), int32(5), int32(5), int32(0), int32(0),
			int32(0), int32(0), int32(0), int32(0),
		).
		WillReturnError(fmt.Errorf("UNIQUE constraint failed"))

	survivorRequest := DTO{
		Settlement: testSettlementID,
		Name:       "Duplicate Name",
		Birth:      1,
		Gender:     "F",
		HuntXp:     0,
		Survival:   1,
		Movement:   5,
		Accuracy:   5,
		Strength:   5,
		Evasion:    5,
		Luck:       5,
		Speed:      5,
	}
	reqBody, _ := json.Marshal(survivorRequest)
	req := httptest.NewRequest("POST", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), bytes.NewReader(reqBody))
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 400, resp.StatusCode, "duplicate name should return 400")
}

func TestCreateSurvivor_ReportsCreationErrors(t *testing.T) {
	_, db, router := setupTest(t)
	defer db.Close()

	db.ExpectExec("INSERT INTO campaign.survivor").
		WithArgs(
			int32(testSettlementID), "Error Survivor", int32(1), "F",
			int32(0), int32(1), int32(5), int32(5), int32(5),
			int32(5), int32(5), int32(5), int32(0), int32(0),
			int32(0), int32(0), int32(0), int32(0),
		).
		WillReturnError(fmt.Errorf("database error"))

	survivorRequest := DTO{
		Settlement: testSettlementID,
		Name:       "Error Survivor",
		Birth:      1,
		Gender:     "F",
		HuntXp:     0,
		Survival:   1,
		Movement:   5,
		Accuracy:   5,
		Strength:   5,
		Evasion:    5,
		Luck:       5,
		Speed:      5,
	}
	reqBody, _ := json.Marshal(survivorRequest)
	req := httptest.NewRequest("POST", fmt.Sprintf("/settlements/%d/survivors", testSettlementID), bytes.NewReader(reqBody))
	ctx := req.Context()
	ctx = context.WithValue(ctx, request.UserIDKey, testUserID)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req.WithContext(ctx))
	resp := w.Result()

	assert.Equal(t, 500, resp.StatusCode, "database errors should return 500")
}

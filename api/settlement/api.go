package settlement

import (
	"github.com/failuretoload/datamonster/request"
	"github.com/failuretoload/datamonster/response"
	"net/http"

	repo "github.com/failuretoload/datamonster/settlement/internal"
	"github.com/failuretoload/datamonster/store"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	db *repo.Repo
}

func NewController(conn store.Connection) *Controller {
	db := repo.New(conn)
	return &Controller{db: db}
}

type DTO struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	SurvivalLimit       int    `json:"limit"`
	DepartingSurvival   int    `json:"departing"`
	CollectiveCognition int    `json:"cc"`
	Year                int    `json:"year"`
}

func (c Controller) RegisterRoutes(r chi.Router) {
	r.Get("/settlements", c.getSettlements)
	r.Post("/settlements", c.createSettlement)
	r.Route("/settlements/{id}", func(r chi.Router) {
		r.Get("/", c.getSettlement)
	})
}

func (c Controller) getSettlements(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(request.UserIDKey).(string)
	settlements, repoErr := c.db.Select(r.Context(), userID)
	if repoErr != nil {
		response.InternalServerError(r.Context(), w, "unable to retrieve settlements", repoErr)
		return
	}
	data := domainListToDto(settlements)
	response.OK(r.Context(), w, data)
}

type CreateSettlementRequest struct {
	Name string `json:"name"`
}

func (c Controller) createSettlement(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(request.UserIDKey).(string)
	if !ok {
		response.BadRequest(ctx, w, "no user id provided", nil)
		return
	}
	var body CreateSettlementRequest
	err := request.DecodeJSONRequest(r.Body, &body)
	if err != nil {
		response.BadRequest(ctx, w, "invalid request body", err)
		return
	}
	if body.Name == "" {
		response.BadRequest(ctx, w, "name is required", nil)
		return
	}
	settlement := repo.Settlement{
		Owner:               userID,
		Name:                body.Name,
		SurvivalLimit:       1,
		DepartingSurvival:   0,
		CollectiveCognition: 0,
		CurrentYear:         1,
	}
	newID, insertErr := c.db.Insert(r.Context(), settlement)
	if insertErr != nil {
		response.InternalServerError(ctx, w, "unable to create settlement", insertErr)
		return
	}

	settlement.ID = newID
	dto := domainToDto(settlement)
	response.OK(r.Context(), w, dto)
}

func (c Controller) getSettlement(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value(request.UserIDKey).(string)
	if !ok {
		response.BadRequest(ctx, w, "no user id provided", nil)
		return
	}
	settlementID := chi.URLParam(r, "id")
	settlement, repoErr := c.db.Get(r.Context(), settlementID, userID)
	if repoErr != nil {
		response.InternalServerError(r.Context(), w, "unable to retrieve settlement", repoErr)
		return
	}
	dto := domainToDto(settlement)
	response.OK(r.Context(), w, dto)
}

func domainListToDto(settlements []repo.Settlement) []DTO {
	dtos := []DTO{}
	for _, s := range settlements {
		dto := DTO{
			ID:                  s.ID,
			Name:                s.Name,
			SurvivalLimit:       s.SurvivalLimit,
			DepartingSurvival:   s.DepartingSurvival,
			CollectiveCognition: s.CollectiveCognition,
			Year:                s.CurrentYear,
		}
		dtos = append(dtos, dto)
	}
	return dtos
}

func domainToDto(s repo.Settlement) DTO {
	return DTO{
		ID:                  s.ID,
		Name:                s.Name,
		SurvivalLimit:       s.SurvivalLimit,
		DepartingSurvival:   s.DepartingSurvival,
		CollectiveCognition: s.CollectiveCognition,
		Year:                s.CurrentYear,
	}
}

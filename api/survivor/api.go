package survivor

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/failuretoload/datamonster/survivor/domain"

	"github.com/failuretoload/datamonster/response"
	"github.com/failuretoload/datamonster/store"
	"github.com/failuretoload/datamonster/survivor/internal/repo"

	"github.com/go-chi/chi/v5"
)

type Controller struct {
	db *repo.Repo
}

func NewController(conn store.Connection) *Controller {
	r := repo.New(conn)
	return &Controller{db: r}
}

func (c Controller) RegisterRoutes(r chi.Router) {
	r.Get("/settlements/{id}/survivors", c.getSurvivors)
	r.Post("/settlements/{id}/survivors", c.createSurvivor)
}

func (c Controller) getSurvivors(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	settlementID, convErr := strconv.Atoi(param)
	if convErr != nil {
		response.InternalServerError(r.Context(), w, "unable to convert query param", convErr)
		return
	}
	survivors, err := c.db.GetAllSurvivorsForSettlement(r.Context(), settlementID)
	if err != nil {
		response.InternalServerError(r.Context(), w, "error retrieving survivors", err)
		return
	}
	data := dtoListFromDomain(survivors)
	response.OK(r.Context(), w, data)
}

func (c Controller) createSurvivor(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	settlementID, convErr := strconv.Atoi(param)
	if convErr != nil {
		response.InternalServerError(r.Context(), w, "settlement id must be a positive integer", convErr)
		return
	}
	survivorDTO := DTO{}
	decodeErr := json.NewDecoder(r.Body).Decode(&survivorDTO)
	if decodeErr != nil || survivorDTO.Name == "" {
		response.InternalServerError(r.Context(), w, "unable to decode request body", decodeErr)
		return
	}
	survivorDTO.Settlement = settlementID
	err := c.db.CreateSurvivor(r.Context(), domainFromDTO(survivorDTO))
	if err != nil {
		dupError := repo.DuplicateNameError{}
		if errors.As(err, &dupError) {
			response.BadRequest(r.Context(), w, fmt.Sprintf("survivor with name %s already exists", survivorDTO.Name), dupError)
			return
		}
		response.InternalServerError(r.Context(), w, "error creating survivor", err)
		return
	}
	response.NoContent(r.Context(), w)
}

type DTO struct {
	ID               int     `json:"id"`
	Settlement       int     `json:"settlement"`
	Name             string  `json:"name"`
	Birth            int     `json:"birth"`
	Gender           string  `json:"gender"`
	Status           *string `json:"status,omitempty"`
	HuntXp           int     `json:"huntXp"`
	Survival         int     `json:"survival"`
	Movement         int     `json:"movement"`
	Accuracy         int     `json:"accuracy"`
	Strength         int     `json:"strength"`
	Evasion          int     `json:"evasion"`
	Luck             int     `json:"luck"`
	Speed            int     `json:"speed"`
	Insanity         int     `json:"insanity"`
	SystemicPressure int     `json:"systemicPressure"`
	Torment          int     `json:"torment"`
	Lumi             int     `json:"lumi"`
	Courage          int     `json:"courage"`
	Understanding    int     `json:"understanding"`
}

func dtoFromDomain(s domain.Survivor) DTO {
	return DTO(s)
}

func domainFromDTO(s DTO) domain.Survivor {
	return domain.Survivor(s)
}

func dtoListFromDomain(s []domain.Survivor) []DTO {
	survivors := make([]DTO, len(s))
	for i, v := range s {
		survivors[i] = dtoFromDomain(v)
	}
	return survivors
}

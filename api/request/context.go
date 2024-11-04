package request

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type ctxUserIDKey string
type SettlementID string

const (
	UserIDKey ctxUserIDKey = "userId"
)

func IDParam(r *http.Request) (int, error) {
	id := chi.URLParam(r, "id")
	val, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

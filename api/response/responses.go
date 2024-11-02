package response

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func writeJSON(rw http.ResponseWriter, status int, data any) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, err = rw.Write(js)
	if err != nil {
		return err
	}
	return nil
}

func BadRequest(ctx context.Context, rw http.ResponseWriter, reason string, err error) {
	writeError := writeJSON(rw, http.StatusBadRequest, reason)
	if writeError != nil {
		slog.Log(ctx, slog.LevelWarn, "error producing BadRequest response", slog.Any("error", writeError))
	}
	var content any
	if err != nil {
		content = slog.Any("error", err)
	}
	slog.Log(ctx, slog.LevelError, "bad request", content)
}

func InternalServerError(ctx context.Context, rw http.ResponseWriter, reason string, err error) {
	writeError := writeJSON(rw, http.StatusInternalServerError, reason)
	if writeError != nil {
		slog.Log(ctx, slog.LevelWarn, "error producing InternalServerError response", slog.Any("error", writeError))
	}
	var content any
	if err != nil {
		content = slog.Any("error", err)
	}
	slog.Log(ctx, slog.LevelError, "internal server error", content)
}

func Unauthorized(ctx context.Context, rw http.ResponseWriter, err error) {
	writeError := writeJSON(rw, http.StatusUnauthorized, err.Error())
	if writeError != nil {
		slog.Log(ctx, slog.LevelWarn, "error producing InternalServerError response", slog.Any("error", writeError))
	}
	slog.Log(ctx, slog.LevelError, "unauthorized", slog.Any("error", err))
}

func NoContent(ctx context.Context, rw http.ResponseWriter) {
	writeError := writeJSON(rw, http.StatusNoContent, nil)
	if writeError != nil {
		slog.Log(ctx, slog.LevelWarn, "error producing InternalServerError response", slog.Any("error", writeError))
	}
	slog.Log(ctx, slog.LevelInfo, "no content")
}

func OK(ctx context.Context, rw http.ResponseWriter, data interface{}) {
	writeError := writeJSON(rw, http.StatusOK, data)
	if writeError != nil {
		slog.Log(ctx, slog.LevelError, "error producing OK response", slog.Any("error", writeError))
	}
	slog.Log(ctx, slog.LevelInfo, "ok")
}

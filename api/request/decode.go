package request

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
)

func DecodeJSONRequest(rc io.ReadCloser, data interface{}) error {
	defer func(rc io.ReadCloser) {
		err := rc.Close()
		if err != nil {
			slog.Log(context.Background(), slog.LevelWarn, "error closing ReadCloser", slog.Any("error", err))
		}
	}(rc)
	decoder := json.NewDecoder(rc)
	return decoder.Decode(data)
}

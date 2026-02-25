package logger

import "log/slog"

func WithRequestID(l *slog.Logger, id string) *slog.Logger {
	return l.With(slog.String(KeyRequestID, id))
}

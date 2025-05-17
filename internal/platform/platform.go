package platform

import "log/slog"

type AppDeps struct {
	Logger *slog.Logger
}

func New(logger *slog.Logger) AppDeps {
	if logger == nil {
		logger = initLogger()
	}
	return AppDeps{Logger: logger}
}

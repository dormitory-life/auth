package logger

import (
	"errors"
	"log/slog"
	"os"

	"github.com/dormitory-life/auth/internal/config"
)

const (
	EnvLocal = "local"
	EnvDegub = "debug"
	EvnProd  = "production"
)

func New(cfg *config.Config) (*slog.Logger, error) {
	switch cfg.Env {
	case EnvLocal:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})), nil

	case EnvDegub:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		})), nil

	case EvnProd:
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelError,
		})), nil

	default:
		return nil, errors.New("unknown log level")
	}
}

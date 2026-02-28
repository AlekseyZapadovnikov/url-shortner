package conf

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/AlekseyZapadovnikov/url-shortner/internal/validation"
	"github.com/ilyakaznacheev/cleanenv"
)

var valid = validation.Valid

type Format string

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

type LoggerConfig struct {
	AppName    string     `env:"APP_NAME" env-default:"url-shortener" validate:"required"`
	Env        string     `env:"APP_ENV" env-default:"dev" validate:"required,oneof=dev prod"`
	Level      slog.Level `env:"LOG_LEVEL" env-default:"debug" validate:"required"`
	Format     Format     `env:"LOG_FORMAT" env-default:"text" validate:"required,oneof=json text"`
	AddSource  bool       `env:"LOG_ADD_SOURCE" env-default:"true"`
	TimeFormat string     `env:"LOG_TIME_FORMAT" env-default:""`
	Output     io.Writer  `env:"-"`
}

func MustLoadLogger() *LoggerConfig {
	var cfg LoggerConfig

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(fmt.Sprintf("read env: %v", err))
	}

	if err := valid.Struct(cfg); err != nil {
		panic(fmt.Sprintf("validate config: %v", err))
	}

	cfg.Output = os.Stdout
	return &cfg
}

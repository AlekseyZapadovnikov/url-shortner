package conf

type HttpConfig struct {
	Host       string `env:"HTTP_HOST" env-required:"true" validate:"required,hostname|ip"`
	Port       int    `env:"HTTP_PORT" env-required:"true" validate:"required,port"`
	StaticPath string `env:"HTTP_STATIC_PATH" env-required:"true" validate:"required,path"`
}

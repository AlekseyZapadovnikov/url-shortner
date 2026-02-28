package conf

import "strconv"

type HttpConfig struct {
	Host       string `env:"HTTP_HOST" env-required:"true" validate:"required,hostname|ip"`
	Port       int    `env:"HTTP_PORT" env-required:"true" validate:"required,port"`
	StaticPath string `env:"HTTP_STATIC_PATH" env-required:"true" validate:"required,path"`
}

func (c *HttpConfig) GetAdress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

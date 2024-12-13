package config

import "github.com/ilyakaznacheev/cleanenv"

type Mode string

const (
	Release Mode = "release"
	Dev     Mode = "dev"
)

var Cfg Config

type Config struct {
	PostgresURL            string `env:"postgres_url" env-required:"true"`
	Port                   string `env:"port" env-required:"true"`
	Mode                   Mode   `env:"mode" env-required:"true"`
	AccessTokenSecret      string `env:"access_token_secret" env-required:"true"`
	AdminAccessTokenSecret string `env:"admin_access_token_secret" env-required:"true"`
}

func (m Mode) IsRelease() bool {
	return m == Release
}

func Init() *Config {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		panic(err)
	}

	return &Cfg
}

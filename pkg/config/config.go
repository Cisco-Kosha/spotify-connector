package config

import (
	"flag"
	"os"
)

type Config struct {
	accessToken  string
	expiresAt    string
	refreshToken string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.accessToken, "accessToken", os.Getenv("ACCESS_TOKEN"), "Access Token")
	flag.StringVar(&conf.expiresAt, "expiresAt", os.Getenv("EXPIRES_AT"), "Expires At")
	flag.StringVar(&conf.refreshToken, "refreshToken", os.Getenv("REFRESH_TOKEN"), "Refresh Token")

	flag.Parse()

	return conf
}

func (c *Config) GetAccessToken() string {
	return c.accessToken
}

func (c *Config) GetExpiresAt() string {
	return c.expiresAt
}

func (c *Config) GetRefreshToken() string {
	return c.refreshToken
}

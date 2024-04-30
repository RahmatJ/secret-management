package config

type (
	EnvConfig struct {
		DbHost             string
		DbPort             string
		DbUser             string
		DbPass             string
		DbName             string
		ApiKeyTTL          int64
		ApiKeyRecreateTime int64
	}
)

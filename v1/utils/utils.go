package utils

type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB int `json:"db"`
	Password string `json:"password"`
	Prefix string `json:"prefix"`
}
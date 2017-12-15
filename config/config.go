package config

import "os"

func Get(key string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	defValue, ok := defaults[key]
	if !ok {
		return ""
	}

	return defValue
}

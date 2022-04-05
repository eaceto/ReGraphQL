package helpers

import (
	"github.com/spf13/viper"
	"os"
)

func GetEnvVar(key string, defValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	value, ok = viper.Get(key).(string)

	if ok {
		return value
	}

	return defValue
}

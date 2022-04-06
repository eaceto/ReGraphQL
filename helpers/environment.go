// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package helpers

import (
	"github.com/spf13/viper"
	"os"
)

// Returns an environment varible if available, if not its defaultValue
// Lookup table:
// - Environment
// - Argument (vipe)
func GetEnvVar(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	value, ok = viper.Get(key).(string)

	if ok {
		return value
	}

	return defaultValue
}

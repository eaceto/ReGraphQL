/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package helpers

import (
	"github.com/spf13/viper"
	"os"
)

// GetEnvVar Returns an environment variable's value if available, if not its defaultValue
// Lookup table:
// - Environment (os)
// - Argument (viper)
// - Default Value (param)
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

/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package app

const (
	EnvironmentVariablesFile = ".env"

	ServerHostKey          = "SERVER_HOST"
	ServerHostDefaultValue = ""

	ServerPortKey          = "SERVER_PORT"
	ServerPortDefaultValue = 8080

	ServicePathKey          = "SERVICE_PATH"
	ServicePathDefaultValue = "/graphql"

	ServerReadTimeoutKey      = "SERVER_READ_TIMEOUT"
	ServerWriteTimeoutKey     = "SERVER_WRITE_TIMEOUT"
	ServerTimeoutDefaultValue = "120"

	RouterConfigPathKey          = "ROUTER_CONFIG_PATH"
	RouterConfigPathDefaultValue = "/data/config/"

	TraceCallsKey          = "TRACE_CALLS"
	TraceCallsDefaultValue = "0"

	DebugKey          = "DEBUG"
	DebugDefaultValue = "0"

	PreAllocatedRoutesNumber = 10

	HealthPath  = "/health"
	MetricsPath = "/metrics"
)

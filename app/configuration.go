package app

import (
	"fmt"
	"github.com/eaceto/ReGraphQL/helpers"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type App struct {
	ServerAddr         string
	ServicePath        string
	ServerReadTimeout  time.Duration
	ServerWriteTimeout time.Duration
	RouterConfigsPath  string
	TraceCallsEnabled  bool
	DebugEnabled       bool
	HTTPClient         *http.Client
}

func NewApp() (*App, error) {

	viper.SetConfigFile(EnvironmentVariablesFile)
	_ = viper.ReadInConfig()

	// Parse Server host and port
	serverHost := helpers.GetEnvVar(ServerHostKey, ServerHostDefaultValue)
	if len(serverHost) > 0 && net.ParseIP(serverHost) != nil {
		return nil, fmt.Errorf("invalid %s value: '%v'", ServerHostKey, serverHost)
	}

	serverPort := helpers.GetEnvVar(ServerPortKey, ServerPortDefaultValue)
	if _, err := strconv.ParseUint(serverPort, 10, 32); err != nil {
		return nil, fmt.Errorf("invalid %s value: '%v'. %v", ServerPortKey, serverPort, err)
	}

	servicePath := helpers.GetEnvVar(ServicePathKey, ServicePathDefaultValue)
	if len(servicePath) == 0 {
		return nil, fmt.Errorf("invalid %s value: '%v'", ServicePathKey, servicePath)
	}
	if !strings.HasPrefix(servicePath, "/") {
		servicePath = "/" + servicePath
	}

	// Parse Server configuration
	serverReadTimeout, serverReadTimeoutError := strconv.Atoi(helpers.GetEnvVar(ServerReadTimeoutKey, ServerTimeoutDefaultValue))
	if serverReadTimeoutError != nil || serverReadTimeout < 1 {
		return nil, fmt.Errorf("invalid %s value: '%v'. %v", ServerReadTimeoutKey, serverReadTimeout, serverReadTimeoutError)
	}

	serverWriteTimeout, serverWriteTimeoutError := strconv.Atoi(helpers.GetEnvVar(ServerWriteTimeoutKey, ServerTimeoutDefaultValue))
	if serverWriteTimeoutError != nil || serverWriteTimeout < 1 {
		return nil, fmt.Errorf("invalid %s value: '%v'. %v", ServerWriteTimeoutKey, serverWriteTimeoutError, serverWriteTimeout)
	}

	traceCallsEnabled := helpers.GetEnvVar(TraceCallsKey, TraceCallsDefaultValue) == "1"
	debugEnabled := helpers.GetEnvVar(DebugKey, DebugDefaultValue) == "1"

	// Parse path for the router configuration files
	routerConfigPath := helpers.GetEnvVar(RouterConfigPathKey, RouterConfigPathDefaultValue)
	if _, err := os.Stat(routerConfigPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path not found: invalid %s value: '%v'", RouterConfigPathKey, routerConfigPath)
	}

	// Return application configuration
	return &App{
		ServerAddr:         serverHost + ":" + serverPort,
		ServicePath:        servicePath,
		RouterConfigsPath:  routerConfigPath,
		ServerReadTimeout:  time.Duration(serverReadTimeout) * time.Second,
		ServerWriteTimeout: time.Duration(serverWriteTimeout) * time.Second,
		TraceCallsEnabled:  traceCallsEnabled,
		DebugEnabled:       debugEnabled,
		HTTPClient:         &http.Client{Timeout: time.Duration(serverReadTimeout) * time.Second},
	}, nil
}

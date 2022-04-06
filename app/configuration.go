/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package app

import (
	"fmt"
	"github.com/eaceto/ReGraphQL/helpers"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Configuration struct {
	ServerAddr         string
	ServicePath        string
	ServerReadTimeout  time.Duration
	ServerWriteTimeout time.Duration
	RouterConfigsPath  string
	TraceCallsEnabled  bool
	DebugEnabled       bool
	HTTPClient         *http.Client
}

func NewConfiguration() (*Configuration, error) {

	viper.SetConfigFile(EnvironmentVariablesFile)
	_ = viper.ReadInConfig()

	// Parse Server host and port
	serverHost := helpers.GetEnvVar(ServerHostKey, ServerHostDefaultValue)
	if len(serverHost) > 0 && net.ParseIP(serverHost) != nil {
		return nil, fmt.Errorf("invalid %s value: '%v'", ServerHostKey, serverHost)
	}

	serverPort := helpers.GetEnvVar(ServerPortKey, fmt.Sprint(ServerPortDefaultValue))
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

	if servicePath == HealthPath {
		return nil, fmt.Errorf("invalid %s value: '%v' has conflicts with reserverd path: %v", ServicePathKey, servicePath, HealthPath)
	}
	if servicePath == MetricsPath {
		return nil, fmt.Errorf("invalid %s value: '%v' has conflicts with reserverd path: %v", ServicePathKey, servicePath, HealthPath)
	}

	// Parse Server Configuration
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

	// Parse path for the router Configuration files
	routerConfigPath := helpers.GetEnvVar(RouterConfigPathKey, RouterConfigPathDefaultValue)
	if _, err := os.Stat(routerConfigPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path not found: invalid %s value: '%v'", RouterConfigPathKey, routerConfigPath)
	}

	// Return application Configuration
	return &Configuration{
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

func (c *Configuration) log() {
	if !c.DebugEnabled {
		return
	}
	klog.Warningln("Debug Enabled")

	klog.Infof("Config files: %s\n", c.RouterConfigsPath)
	klog.Infof("Service path: %s\n", c.ServicePath)

	if c.TraceCallsEnabled {
		klog.Infoln("TraceCalls Enabled")
	}
}

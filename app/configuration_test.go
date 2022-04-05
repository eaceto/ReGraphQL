package app

import (
	"strconv"
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	testConfigPath := "../tests/files/starwars/"
	t.Setenv(RouterConfigPathKey, testConfigPath) // Avoid an error as config directory does not exists
	config, err := NewApp()

	if err != nil {
		t.Error(err)
		return
	}

	if config.DebugEnabled != false {
		t.Errorf("Debug should be disabled by default ")
	}
	if config.TraceCallsEnabled != false {
		t.Errorf("Trace Calls should be disabled by default ")
	}
	if config.RouterConfigsPath != testConfigPath {
		t.Errorf("`config.RouterConfigsPath = %s`, want: \"%s\"", config.RouterConfigsPath, testConfigPath)
	}
	if config.ServerAddr != ServerHostDefaultValue+":"+ServerPortDefaultValue {
		t.Errorf("`config.ServerAddr = %s`, want: \"%s\"", config.ServerAddr, ServerHostDefaultValue+":"+ServerPortDefaultValue)
	}

	timeout, _ := strconv.ParseUint(ServerTimeoutDefaultValue, 10, 32)
	if config.ServerReadTimeout != time.Duration(timeout)*time.Second {
		t.Errorf("`config.ServerReadTimeout = %s`, want: \"%s\"", config.ServerReadTimeout, time.Duration(timeout)*time.Second)
	}
	if config.ServerWriteTimeout != time.Duration(timeout)*time.Second {
		t.Errorf("`config.ServerWriteTimeout = %s`, want: \"%s\"", config.ServerWriteTimeout, time.Duration(timeout)*time.Second)
	}
}

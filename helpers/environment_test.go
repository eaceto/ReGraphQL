// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package helpers

import (
	"github.com/spf13/viper"
	"testing"
)

func TestGetEnvVar(t *testing.T) {
	value := "value"
	def := "def_value"
	key := "AN_ENV_VAR"

	t.Setenv(key, value)

	got := GetEnvVar(key, def)
	if value != got {
		t.Errorf("`value != %s`. want %s", got, value)
	}
}

func TestGetEnvVarDefault(t *testing.T) {
	def := "def_value"
	key := "AN_ENV_VAR"

	got := GetEnvVar(key, def)
	if def != got {
		t.Errorf("`def != %s`. want %s", got, def)
	}
}

func TestGetEnvVarViper(t *testing.T) {
	viper.SetConfigFile("../tests/viper/test.env")
	_ = viper.ReadInConfig()

	value := "value"
	def := "def_value"
	key := "AN_ENV_VAR"

	got := GetEnvVar(key, def)
	if value != got {
		t.Errorf("`value != %s`. want %s", got, value)
	}
}

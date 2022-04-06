// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package app

import (
	"errors"
	"io/fs"
	"testing"
)

func TestSingleFileRoutes(t *testing.T) {
	appConfiguration := App{
		ServerAddr:        ":8081",
		RouterConfigsPath: "../tests/files/starwars/",
		TraceCallsEnabled: true,
		DebugEnabled:      true,
	}
	routes, err := appConfiguration.LoadRoutesFromFiles()
	if err != nil {
		t.Errorf("Could not load routes. %v", err)
	}

	if len(routes) != 2 {
		t.Errorf("`len(routes) = %v`, want: 2 routes", len(routes))
	}

	route1 := routes[0]
	if route1.HTTP.URI != "/persons/{person}" {
		t.Errorf("`route1.HTTP.URI == %s`, want: \"/persons/{person}\"", route1.HTTP.URI)
	}
	if route1.HTTP.Method != "GET" {
		t.Errorf("`route1.HTTP.Method == %s`, want: \"GET\"", route1.HTTP.Method)
	}
	if route1.GraphQL.Endpoint != "https://swapi.skyra.pw/" {
		t.Errorf("`route1.GraphQL.Endpoint == %s`, want: \"https://swapi.skyra.pw/\"", route1.GraphQL.Endpoint)
	}

	route2 := routes[1]
	if route2.HTTP.URI != "/films/{film}" {
		t.Errorf("`route1.HTTP.URI == %s`, want: \"/films/{film}\"", route1.HTTP.URI)
	}
	if route2.HTTP.Method != "GET" {
		t.Errorf("`route1.HTTP.Method == %s`, want: \"GET\"", route1.HTTP.Method)
	}
	if route2.GraphQL.Endpoint != "https://swapi.skyra.pw/" {
		t.Errorf("`route1.GraphQL.Endpoint == %s`, want: \"https://swapi.skyra.pw/\"", route1.GraphQL.Endpoint)
	}
}

func TestMultipleFilesRoutes(t *testing.T) {
	appConfiguration := App{
		ServerAddr:        ":8081",
		RouterConfigsPath: "../tests/files/multiple_files",
		TraceCallsEnabled: true,
		DebugEnabled:      true,
	}
	routes, err := appConfiguration.LoadRoutesFromFiles()
	if err != nil {
		t.Errorf("Could not load routes. %v", err)
	}

	if len(routes) != 2 {
		t.Errorf("`len(routes) = %v`, want: 2 routes", len(routes))
	}
}

func TestEmptyRoutesPath(t *testing.T) {
	appConfiguration := App{
		ServerAddr:        ":8081",
		RouterConfigsPath: "../tests/files/empty",
		TraceCallsEnabled: true,
		DebugEnabled:      true,
	}
	routes, err := appConfiguration.LoadRoutesFromFiles()
	if err != nil {
		t.Errorf("Could not load routes. %v", err)
	}

	if len(routes) != 0 {
		t.Errorf("`len(routes) = %v`, want: 0 routes", len(routes))
	}
}

func TestInvalidRoutesPath(t *testing.T) {
	appConfiguration := App{
		ServerAddr:        ":8081",
		RouterConfigsPath: "../tests/files/not_found",
		TraceCallsEnabled: true,
		DebugEnabled:      true,
	}
	_, err := appConfiguration.LoadRoutesFromFiles()

	var pathErr *fs.PathError
	if err != nil && errors.As(err, &pathErr) {
		return
	}
	t.Errorf("Could not load routes. %v", err)
}

func TestInvalidYaml(t *testing.T) {
	appConfiguration := App{
		ServerAddr:        ":8081",
		RouterConfigsPath: "../tests/files/invalid_yaml",
		TraceCallsEnabled: true,
		DebugEnabled:      true,
	}
	routes, _ := appConfiguration.LoadRoutesFromFiles()

	if len(routes) != 0 {
		t.Errorf("`len(routes) = %v`, want: 0 routes", len(routes))
	}
}

func TestIsYamlFile(t *testing.T) {
	if isYaml("../tests/files/empty/empty.yaml") == false {
		t.Errorf("`isYaml(fileInfo) == false`, want: true")
	}
}

func TestIsYmlFile(t *testing.T) {
	if isYaml("../tests/files/starwars/starwars.yml") == false {
		t.Errorf("`isYaml(fileInfo) == false`, want: true")
	}
}

func TestIsNotYamlFile(t *testing.T) {
	if isYaml("../tests/files/json/file.json") == true {
		t.Errorf("`isYaml(fileInfo) == true`, want: false")
	}
}

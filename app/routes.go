// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/klog/v2"
	"os"
	"path/filepath"
)

type Route struct {
	HTTP struct {
		URI    string `yaml:"uri"`
		Method string `yaml:"method"`
	} `yaml:"http"`
	GraphQL struct {
		Endpoint string            `yaml:"endpoint"`
		Query    string            `yaml:"query"`
		Types    map[string]string `yaml:"types"`
	} `yaml:"graphql"`
}

type routesConfig struct {
	Routes []Route `yaml:"routes"`
}

func isYaml(path string) bool {
	ext := filepath.Ext(path)
	hasYamlExt := ext == ".yaml" || ext == ".yml"
	return hasYamlExt
}

func (a *App) LoadRoutesFromFiles() ([]Route, error) {
	routes := make([]Route, 0, PreAllocatedRoutesNumber)

	if a.DebugEnabled {
		klog.Infof("Walking config files path: `%s`", a.RouterConfigsPath)
	}
	err := filepath.Walk(a.RouterConfigsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if a.DebugEnabled {
				klog.Errorf("Error walking config files path: `%s`", a.RouterConfigsPath)
			}
			return err
		}

		if info != nil && !info.IsDir() && len(path) > 0 && isYaml(path) {
			if a.DebugEnabled {
				klog.Infof("Reading file: `%s`", path)
			}
			file, fileErr := ioutil.ReadFile(path)
			if fileErr != nil {
				return fmt.Errorf("error reading file: %s. %v", path, fileErr)
			}

			var fileConfig routesConfig
			ymlErr := yaml.Unmarshal(file, &fileConfig)
			if ymlErr != nil {
				return fmt.Errorf("error decoding yaml @ file: %s. %v", path, fileErr)
			}

			routes = append(routes, fileConfig.Routes...)
		} else if info != nil && info.IsDir() {
			if a.DebugEnabled {
				klog.Infof("Found directory: `%s`", info.Name())
			}
		} else if !isYaml(path) {
			if a.DebugEnabled {
				klog.Warningf("Found non-yaml file: `%s`", path)
			}
		}
		return nil
	})

	if err != nil && a.DebugEnabled {
		klog.Errorf("Error walking config files path: `%s`", a.RouterConfigsPath)
	}

	return routes, err
}

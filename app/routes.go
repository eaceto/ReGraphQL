/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
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
	Errors struct {
		HidePath      bool `yaml:"hidePath" default:"false"`
		HideLocations bool `yaml:"hideLocations" default:"false"`
		Extensions    struct {
			Hide        bool           `yaml:"hide" default:"false"`
			CodeMapping map[string]int `yaml:"codeMapping"`
		} `yaml:"extensions"`
	} `yaml:"errors"`
}

func (route *Route) shouldModifyResponse() bool {
	return route.Errors.HideLocations || route.Errors.HidePath || route.Errors.Extensions.Hide
}

type routesConfig struct {
	Routes []Route `yaml:"routes"`
}

func isYaml(path string) bool {
	ext := filepath.Ext(path)
	hasYamlExt := ext == ".yaml" || ext == ".yml"
	return hasYamlExt
}

func (c *Configuration) loadRoutesFromFiles() ([]Route, error) {
	routes := make([]Route, 0, PreAllocatedRoutesNumber)

	if c.DebugEnabled {
		klog.Infof("Walking config files path: `%s`", c.RouterConfigsPath)
	}
	err := filepath.Walk(c.RouterConfigsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			if c.DebugEnabled {
				klog.Errorf("Error walking config files path: `%s`", c.RouterConfigsPath)
			}
			return err
		}

		if info != nil && !info.IsDir() && len(path) > 0 && isYaml(path) {
			if c.DebugEnabled {
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
			if c.DebugEnabled {
				klog.Infof("Found directory: `%s`", info.Name())
			}
		} else if !isYaml(path) {
			if c.DebugEnabled {
				klog.Warningf("Found non-yaml file: `%s`", path)
			}
		}
		return nil
	})

	if err != nil && c.DebugEnabled {
		klog.Errorf("Error walking config files path: `%s`", c.RouterConfigsPath)
	}

	return routes, err
}

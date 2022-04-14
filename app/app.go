/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package app

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Application struct {
	Router        *mux.Router
	Routes        []Route
	Configuration *Configuration
}

func NewApplication(rootRouter *mux.Router) (*Application, error) {

	configuration, err := NewConfiguration()
	if err != nil {
		return nil, err
	}

	configuration.log()

	routes, err := configuration.loadRoutesFromFiles()
	if err != nil {
		return nil, err
	}

	if len(routes) == 0 {
		return nil, fmt.Errorf("no routes available in config path: %s", configuration.RouterConfigsPath)
	}

	router, err := configuration.addServiceHTTPRouter(rootRouter, routes)
	if router == nil {
		return nil, fmt.Errorf("could not create HTTP router")
	}

	return &Application{
		Router:        router,
		Routes:        routes,
		Configuration: configuration,
	}, err
}

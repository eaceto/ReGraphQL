/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package services

import (
	"github.com/eaceto/ReGraphQL/app"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Service struct {
	application *app.Application
}

func AddServiceEndpoints(application *app.Application, baseRouter *mux.Router) {

	a := Service{application: application}

	const path = app.HealthPath

	healthRouter := baseRouter.PathPrefix(path).Subrouter().StrictSlash(true)
	healthRouter.
		Path("/liveness").
		Name("Liveness").
		Methods("GET").
		HandlerFunc(a.liveness)

	healthRouter.
		Path("/readiness").
		Name("Readiness").
		Methods("GET").
		HandlerFunc(a.readiness)

	baseRouter.
		Path(app.MetricsPath).
		Name("Metrics").
		Methods("GET").
		Handler(promhttp.Handler())

	baseRouter.NotFoundHandler = http.NotFoundHandler()
}

/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eaceto/ReGraphQL/middlewares"
	"github.com/gorilla/mux"
	"io"
	"k8s.io/klog/v2"
	"net/http"
	"strconv"
)

type graphQLPayload struct {
	OperationName *string                `json:"operationName,omitempty"`
	Query         string                 `json:"query,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
}

func (c *Configuration) addServiceHTTPRouter(router *mux.Router, routes []Route) (*mux.Router, error) {
	r := router.PathPrefix(c.ServicePath).Subrouter()

	if c.TraceCallsEnabled {
		r.Use(middlewares.TraceCallsMiddleware)
	}

	for idx, route := range routes {
		if c.DebugEnabled {
			klog.Infof("Loading route #%v - [%s] %s\n", idx, route.HTTP.Method, route.HTTP.URI)
		}

		var handler http.HandlerFunc
		handler = middlewares.PrometheusMiddleware(c.createHandlerFunc(route))

		r.HandleFunc(route.HTTP.URI, handler).
			Name("HTTP(" + route.HTTP.URI + ")").
			Methods(route.HTTP.Method)
	}

	return r, nil
}

func (c *Configuration) createHandlerFunc(route Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := c.createQueryParams(r, route)
		if err != nil {
			http.Error(w, "could not create request", http.StatusBadRequest)
			return
		}

		data := graphQLPayload{
			OperationName: nil,
			Query:         route.GraphQL.Query,
			Variables:     params,
		}

		payloadBytes, _ := json.Marshal(data)
		graphQLBody := bytes.NewReader(payloadBytes)

		graphQLRequest, err := http.NewRequest("POST", route.GraphQL.Endpoint, graphQLBody)
		if err != nil {
			http.Error(w, "could not create request", http.StatusBadRequest)
			return
		}

		for key, vv := range r.Header {
			for _, v := range vv {
				graphQLRequest.Header.Add(key, v)
			}
		}

		graphQLRequest.Header.Add("Accept", "*/*")
		graphQLRequest.Header.Add("Content-Type", "application/json")
		graphQLRequest.Header.Set("X-Forwarded-For", r.RemoteAddr)

		response, err := c.HTTPClient.Do(graphQLRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Step 4: copy payload to response writer
		copyHeader(w.Header(), response.Header)
		w.WriteHeader(response.StatusCode)
		_, _ = io.Copy(w, response.Body)
		_ = response.Body.Close()
	}
}

func (c *Configuration) createQueryParams(r *http.Request, route Route) (map[string]interface{}, error) {
	params := make(map[string]interface{})

	reqVars := mux.Vars(r)
	for k, v := range reqVars {
		params[k] = v
	}

	for k, t := range route.GraphQL.Types {
		if val, ok := params[k]; ok {
			switch t {
			case "Bool":
				f, err := strconv.ParseBool(reqVars[k])
				if err == nil {
					params[k] = f
				} else {
					return nil, fmt.Errorf("cannot convert value '%s' for key '%s' into type '%s'", val, k, t)
				}
			case "Float":
				f, err := strconv.ParseFloat(reqVars[k], 64)
				if err == nil {
					params[k] = f
				} else {
					return nil, fmt.Errorf("cannot convert value '%s' for key '%s' into type '%s'", val, k, t)
				}
			case "Int":
				f, err := strconv.ParseInt(reqVars[k], 10, 64)
				if err == nil {
					params[k] = f
				} else {
					return nil, fmt.Errorf("cannot convert value '%s' for key '%s' into type '%s'", val, k, t)
				}
			default:
				continue
			}
		}
	}

	return params, nil
}

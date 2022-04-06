/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package middlewares

import (
	"k8s.io/klog/v2"
	"net/http"
)

func TraceCallsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		klog.InfoS("Handling request", "uri", r.RequestURI, "method", r.Method, "#headers", len(r.Header), "content-type", r.Header["Content-Type"], "content-length", r.ContentLength)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

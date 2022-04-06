/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package helpers

import (
	"github.com/gorilla/mux"
	"k8s.io/klog/v2"
)

func LogEndpoints(r *mux.Router) {
	klog.Info("Exposed endpoints")
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		name := route.GetName()
		klog.InfoS("Found HTTP endpoint", "name", name, "methods", methods, "path", path)
		return nil
	})
}

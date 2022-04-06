/*
 * ReGraphQL - Proxy
 * This is the proxy service of project ReGraphQL
 *
 * Contact: ezequiel.aceto+regraphql@gmail.com
 */

package services

import (
	"encoding/json"
	"net/http"
	"os"
)

func (s *Service) liveness(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	hostname, ok := s.getHostname(w)
	if !ok {
		return
	}

	s.responseWithStatus(w, http.StatusOK, "up", hostname)
}

func (s *Service) readiness(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	hostname, ok := s.getHostname(w)
	if !ok {
		return
	}

	if len(s.application.Routes) == 0 {
		s.responseWithStatus(w, http.StatusPreconditionFailed, "waiting", hostname)
		return
	}

	s.responseWithStatus(w, http.StatusOK, "ready", hostname)
}

func (s *Service) getHostname(w http.ResponseWriter) (string, bool) {
	hostname, hostnameErr := os.Hostname()
	if hostnameErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return "", false
	}
	return hostname, true
}

func (s *Service) responseWithStatus(w http.ResponseWriter, statusCode int, status string, hostname string) {
	data := map[string]string{"status": status, "hostname": hostname}
	jsonData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

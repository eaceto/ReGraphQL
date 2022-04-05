// Copyright 2021 Ezequiel (Kimi) Aceto. All rights reserved.

package app

import (
	"net/http"
)

// copyHeader from "/net/http/httputil/reverseproxy.go"
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

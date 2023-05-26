/******************************************************************************
*
*  Copyright 2020 SAP SE
*
*  Licensed under the Apache License, Version 2.0 (the "License");
*  you may not use this file except in compliance with the License.
*  You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
*
******************************************************************************/

package test

import (
	"net/http"
	"net/http/httptest"
)

// RoundTripper is a http.RoundTripper that redirects some domains to
// http.Handler instances.
type RoundTripper struct {
	Handlers map[string]http.Handler
}

var originalDefaultTransport http.RoundTripper

// WithRoundTripper sets up a RoundTripper instance as the default HTTP
// transport for the duration of the given action.
func WithRoundTripper(action func(*RoundTripper)) {
	if originalDefaultTransport != nil {
		panic("WithRoundTripper calls may not be nested")
	}

	t := RoundTripper{Handlers: make(map[string]http.Handler)}
	originalDefaultTransport = http.DefaultTransport
	http.DefaultTransport = &t
	//The cleanup is in a defer, rather than just at the end of the function,
	//in order to work correctly even if action() does a t.Fatal() or panic().
	defer func() {
		http.DefaultTransport = originalDefaultTransport
		originalDefaultTransport = nil
	}()

	action(&t)
}

// WithoutRoundTripper can be used during WithRoundTripper() to temporarily revert back to the
func WithoutRoundTripper(action func()) {
	if originalDefaultTransport == nil {
		panic("WithoutRoundTripper must be called from within WithRoundTripper")
	}

	prevTransport := http.DefaultTransport
	http.DefaultTransport = originalDefaultTransport
	action()
	http.DefaultTransport = prevTransport
}

// RoundTrip implements the http.RoundTripper interface.
func (t *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	//only intercept requests when the target host is known to us
	h := t.Handlers[req.URL.Host]
	if h == nil {
		return originalDefaultTransport.RoundTrip(req)
	}

	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	resp := w.Result()

	//in practice, most HTTP handlers for GET/HEAD requests write into the
	//response body regardless of whether the method was GET or HEAD; strip the
	//response body from HEAD responses to align with net/http's actual behavior
	if req.Method == http.MethodHead {
		resp.Body = nil
	}

	return resp, nil
}

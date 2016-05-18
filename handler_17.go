// +build go1.7

package httptreemux

import (
	"net/http"
)

type InternalHandlerFunc http.Handler
type InternalPanicHandler http.Handler

func createHandlerWrapper(f HandlerFunc) InternalHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.Context.Get("params").(map[string]string)
		f(w, r, params)
	}
}

func createPanicWrapper(f PanicHandler) InternalPanicHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.Context.Get("params")
		f(w, r, params)
	}
}

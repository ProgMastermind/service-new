// Package debug provides handler supprort for the debugging endpoints
package debug

import (
	"expvar"
	"net/http"
	"net/http/pprof"

	"github.com/arl/statsviz"
)

// Mux registers all the debug routes from the standard libarary into a new mux
// bypassing the use of the DefaultServeMux. Using the DefaultServeMux would
// be a security risk since a dependency could inject a handler into our service
// without us knowing it.

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	mux.Handle("/debug/vars/", expvar.Handler())

	statsviz.Register(mux)

	return mux
}

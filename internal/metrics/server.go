package metrics

import (
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New() http.Handler {
	//use separated ServeMux to prevent handling on the global Mux
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/debug/pprof", pprof.Index)
	mux.HandleFunc("/debug/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/trace", pprof.Trace)

	mux.HandleFunc("/debug/profile", pprof.Profile)
	mux.Handle("/debug/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/threadcreate", pprof.Handler("threadcreate"))
	mux.Handle("/debug/heap", pprof.Handler("heap"))
	mux.Handle("/debug/block", pprof.Handler("block"))
	mux.Handle("/debug/mutex", pprof.Handler("mutex"))
	mux.Handle("/debug/allocs", pprof.Handler("allocs"))

	return mux
}

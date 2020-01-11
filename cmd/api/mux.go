package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type MyRouter struct {
	*mux.Router
}

func (r *MyRouter) ServerHttp(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.Write([]byte(err.(error).Error()))
		}
	}()
	r.Router.ServeHTTP(w, req)
}

func NewApiMux() http.Handler {
	r := MyRouter{mux.NewRouter()}
	s := r.PathPrefix("/api").Subrouter()
	initUserApi(s)

	r.HandleFunc("/", DefaultHadle)

	return r
}

func DefaultHadle(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusNotFound)
	//w.Write([]byte("defalut..."))
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

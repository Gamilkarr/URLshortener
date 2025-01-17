package router

import "net/http"

type IHandler interface {
	Shorten(res http.ResponseWriter, req *http.Request)
	Resolve(res http.ResponseWriter, req *http.Request)
}

func New(h IHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, h.Shorten)
	mux.HandleFunc(`/{id}`, h.Resolve)
	return mux
}

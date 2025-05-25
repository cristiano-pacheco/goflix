package router

import (
	"github.com/julienschmidt/httprouter"

	"github.com/cristiano-pacheco/goflix/internal/shared/modules/httpserver"
)

type Router struct {
	server *httpserver.HTTPServer
}

func NewRouter(server *httpserver.HTTPServer) *Router {
	return &Router{server: server}
}

func (r *Router) Router() *httprouter.Router {
	return r.server.Router()
}

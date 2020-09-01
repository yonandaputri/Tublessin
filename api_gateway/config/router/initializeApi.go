package router

import (
	"fmt"
	"net/http"
	"tublessin/api_gateway/domains/login"

	"github.com/gorilla/mux"
)

type ConfigRouter struct {
	Router *mux.Router
}

// NewAppRouter for creating new Route
func NewAppRouter(r *mux.Router) *ConfigRouter {
	return &ConfigRouter{Router: r}
}

func (ar *ConfigRouter) InitRouter() {
	login.InitLoginRoute(LOGIN_MAIN_ROUTE, ar.Router)
	ar.Router.NotFoundHandler = http.HandlerFunc(notFound)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, `<h1>404 Status Not Found</h1>`)
}

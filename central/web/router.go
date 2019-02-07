package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router struct
type Router struct {
	*mux.Router
}

// Route is a structure of Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter router
func NewRouter(dataCtrl *DataController) *Router {
	router := Router{mux.NewRouter()}
	router.StrictSlash(false)
	AddDataRoutes(dataCtrl, router)
	return &router
}

// AddDataRoutes add
func AddDataRoutes(dataCtrl *DataController, router Router) {
	for _, route := range dataCtrl.Routes {
		router.
			Methods(route.Method).
			Path(dataCtrl.Prefix + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

package router

import (
	"github.com/BIQDev/perpus-api/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var BIQRouter *mux.Router
var initialized bool

type BIQRoute struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func init() {
	if initialized == true {
		return
	}

	BIQRouter = mux.NewRouter()
	BIQRouter.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	BIQRouter.PathPrefix("/assets").Handler(http.FileServer(http.Dir("./web")))
	initialized = true
}

func MiddlewareInit() {
	BIQRouter.Use(middleware.BIQMiddleware.Logger(BIQRouter))
}

func BIQRouteApply(routes []BIQRoute) {
	for _, route := range routes {
		handler := route.HandlerFunc

		BIQRouter.
			Methods(
				route.Methods...,
			).Path(
			route.Pattern,
		).Name(
			route.Name,
		).Handler(
			handler,
		)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

	if os.Getenv("GO_ENV") == "development" {
		log.Println("Error 404, at URL:", r.RequestURI)
	}
}

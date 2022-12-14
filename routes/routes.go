package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoute(r)
	AuthRoutes(r)
	ProductRoutes(r)
	TopingRoutes(r)
	OrderRoutes(r)
	TransactionRoutes(r)
}

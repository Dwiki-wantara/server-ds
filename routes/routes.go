package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	MusicRoutes(r)
	ArtisRoutes(r)
	TransactionRoutes(r)
}

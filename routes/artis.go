package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtisRoutes(r *mux.Router) {
	artisRepository := repositories.RepositoryArtis(mysql.DB)
	h := handlers.HandlerArtis(artisRepository)

	r.HandleFunc("/artis", h.FindArtis).Methods("GET")
	r.HandleFunc("/artis/{id}", h.GetArtis).Methods("GET")
	r.HandleFunc("/artis", middleware.Auth(h.CreateArtis)).Methods("POST")
	r.HandleFunc("/artis/{id}", middleware.Auth(h.UpdateArtis)).Methods("PATCH")
	r.HandleFunc("/artis/{id}", middleware.Auth(h.DeleteArtis)).Methods("DELETE")
}

package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/the-arcade-01/anime-poll-app/internal/service"
)

type Server struct {
	Router *chi.Mux
}

func (s *Server) mountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Heartbeat("/ping"))
}

func (s *Server) mountHandlers() {
	apiService := service.NewApiService()
	s.Router.Get("/greet", apiService.Greet)
	s.Router.Get("/start/ingestion", apiService.StartDBAnimeIngestion)
	s.Router.Get("/db/flush", apiService.FlushAnimeDB)
	s.Router.Get("/db/{id}", apiService.DeleteAnimeById)
	s.Router.Get("/db/animes", apiService.FetchAllAnimes)
	s.Router.Get("/anime/fight", apiService.GetAnimesForFaceOff)
	s.Router.Post("/anime/vote/{id}", apiService.VoteAnime)
}

func NewServer() *Server {
	server := &Server{
		Router: chi.NewRouter(),
	}
	server.mountMiddlewares()
	server.mountHandlers()
	return server
}

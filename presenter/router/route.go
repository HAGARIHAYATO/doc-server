package router

import (
	"doc-server/interactor"
	serverMiddleware "doc-server/presenter/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Route *chi.Mux
}

func NewRouter() *Server {
	return &Server {
		Route: chi.NewRouter(),
	}
}

func (s *Server) Router(h *interactor.Handler, m serverMiddleware.ServerMiddleware) {
	s.Route.Use(middleware.Logger)
	s.Route.Use(middleware.Recoverer)
	s.Route.Use(m.CORS)
	s.Route.Route("/api/v1", func(r chi.Router) {
		r.Route("/docs", func(r chi.Router) {
			r.Get("/", h.DocHandler.DocIndex)
			r.Post("/", h.DocHandler.DocCreate)
			r.Get("/{id}", h.DocHandler.DocShow)
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.UserHandler.UserIndex)
			r.Post("/create", h.UserHandler.UserCreate)
			r.Post("/session", h.UserHandler.UserSession)
			r.Get("/auth", h.UserHandler.VerifyAccess)
			r.Route("/{user_id}/bundles", func(r chi.Router) {
				r.Get("/", h.BundleHandler.BundleIndex)
				r.Get("/{id}", h.BundleHandler.BundleShow)
			})
		})
	})
}
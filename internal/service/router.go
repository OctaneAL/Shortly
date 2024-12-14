package service

import (
	"context"

	"github.com/OctaneAL/Shortly/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *Service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(context.Background(), s.db),
		),
	)

	r.Route("/integrations/Shortly", func(r chi.Router) {
		r.Post("/decode", handlers.Decode)
		r.Get("/encode", handlers.Encode)
	})

	return r
}

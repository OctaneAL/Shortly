package service

import (
	"github.com/OctaneAL/Shortly/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
		),
	)

	r.Get("/hello", handlers.HelloWorld)
	r.Post("/decode", handlers.Decode)

	// r.Route("/integrations/Shortly", func(r chi.Router) {
	// 	// configure endpoints here
	// })

	return r
}

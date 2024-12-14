package service

import (
	"net"
	"net/http"

	"github.com/OctaneAL/Shortly/internal/config"
	"github.com/OctaneAL/Shortly/internal/db"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	db       *db.DB
}

func (s *Service) run() error {
	s.log.Info("Service started")
	r := s.router()

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *Service {
	database := db.NewDB(cfg.DatabaseURL())

	return &Service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		db:       database,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}

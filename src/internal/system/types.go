package system

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type Service interface {
	Db() *sql.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
}

type Module interface {
	Start()
}

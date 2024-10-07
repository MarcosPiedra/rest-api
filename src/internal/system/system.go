package system

import (
	"database/sql"
	"doctor-search-engine/internal/config"
	"doctor-search-engine/internal/logger"
	"io/fs"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
)

type System struct {
	cfg    config.AppConfig
	db     *sql.DB
	mux    *chi.Mux
	logger zerolog.Logger
}

func NewSystem(cfg config.AppConfig) (*System, error) {
	var system = &System{}

	system.cfg = cfg

	system.initLog()
	system.initDb()
	system.initMux()

	return system, nil
}

func (s *System) initLog() {
	s.logger = logger.NewLogger(logger.LogConfig{
		Environment: s.cfg.Environment,
		LogLevel:    logger.Level(s.cfg.LogLevel),
	})
}

func (s *System) initMux() {
	s.mux = chi.NewMux()
	s.mux.Use(middleware.URLFormat)
	s.mux.Use(render.SetContentType(render.ContentTypeJSON))
	s.mux.Use(middleware.Recoverer)
}

func (s *System) initDb() {
	var err error
	s.db, err = sql.Open("pgx", s.cfg.Postgres.Conn)
	if err != nil {
		s.logger.Err(err).Msg("init db")

		panic("Error in DB!")
	}
}

func (s *System) MigrateDb(fs fs.FS) error {
	goose.SetBaseFS(fs)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(s.db, "."); err != nil {
		return err
	}
	return nil
}

func (s *System) Db() *sql.DB {
	return s.db
}

func (s *System) Mux() *chi.Mux {
	return s.mux
}

func (s *System) Cfg() config.AppConfig {
	return s.cfg
}

func (s *System) Logger() zerolog.Logger {
	return s.logger
}

func (s *System) StartWebServer() {
	webServer := &http.Server{
		Addr:    s.cfg.Web.Address(),
		Handler: s.mux,
	}

	s.logger.Info().Msgf("** web server started; listening at http://localhost%s\n", s.cfg.Web.Port)
	defer s.logger.Info().Msg("** web server shutdown")
	if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Err(err)

		return
	}
}

func (s *System) Shutdown() {
	s.logger.Info().Msg("** closing db")
	s.db.Close()
}

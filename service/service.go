package service

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Service - handle microservice bootstrapping
type Service interface {
	Config() *Config
	API() http.Handler
	Logger() *logrus.Entry
	RunMigration(fn Migrate, schemaFs fs.FS, schema string) error
}
type Migrate func(db *sql.DB, schemaFs fs.FS, schema string) error

type Config struct {
	Address  string
	LogLevel uint32
	tlsCfg   *tls.Config
}
type service struct {
	config *Config
	logger *logrus.Entry
	api    http.Handler
	db     *sql.DB
}

func NewService(config *Config, logger *logrus.Entry, db *sql.DB, api http.Handler) Service {
	return &service{
		config: config,
		logger: logger,
		api:    api,
		db:     db,
	}
}

func (s *service) API() http.Handler {
	return s.api
}
func (s *service) Config() *Config {
	return s.config
}
func (s *service) Logger() *logrus.Entry {
	return s.logger
}
func (s *service) RunMigration(fn Migrate, schemaFs fs.FS, schema string) error {
	return fn(s.db, schemaFs, schema)
}

// Run it's a blocking method that runs the http server
func Run(ctx context.Context, s Service) error {
	mainCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	srv := http.Server{
		Addr:      s.Config().Address,
		Handler:   s.API(),
		TLSConfig: s.Config().tlsCfg,
	}
	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		s.Logger().Infof("starting go server on port: %s ", s.Config().Address)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger().Errorf("Failed to listen and serve: %+v", err)
			return fmt.Errorf("failed to Listen and serve %v", err)
		}
		return nil
	})
	g.Go(func() error {
		<-gCtx.Done()
		s.Logger().Infof("Shutting down the http server on %s", s.Config().Address)
		if err := srv.Shutdown(context.Background()); err != nil {
			s.Logger().Errorf("Failed to shutdown the http server: %v\n", err)
			return fmt.Errorf("failed to shutdown the http server %v", err)
		}
		s.Logger().Info("Server shutdown gracefully")
		return nil
	})
	if err := g.Wait(); err != nil {
		s.Logger().Infof("Exiting : %+v\n", err)
		return err
	}
	return nil
}

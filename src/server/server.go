package server

import (
	"catalog/src/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const (
	AppShutdownTimeout = 15
	SysShutdownTimeout = 3
)

type Server struct {
	Multiplexer *http.ServeMux
	Log         *zap.Logger
    Cache       *redis.Client
	config *config.ServerConfig

	panicHook func(w http.ResponseWriter, req *http.Request, err interface{})
}

func NewServer(config *config.ServerConfig, log *zap.Logger, cacheClient *redis.Client) *Server {
	s := &Server{
		Log:    log,
		config: config,
		Cache: cacheClient,

	}

	s.Log.Info(fmt.Sprintf("Server config: %+v", config))

	return s
}

func (s *Server) Run(router *mux.Router) {
	halt := make(chan os.Signal, 1)

	signal.Notify(halt, syscall.SIGTERM, syscall.SIGINT)

	systemServer := s.initSystemServer()

	go func() {
		if err := systemServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Log.Fatal("system server failed", zap.Error(err))
		}
	}()

	appServer := s.initAppServer(router)

	go func() {
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Log.Fatal("app server failed", zap.Error(err))
		}
	}()

	<-halt
	s.Log.Info("Server shutdown started")

	ctx, cancelFunc := context.WithTimeout(context.Background(), AppShutdownTimeout*time.Second)
	if err := appServer.Shutdown(ctx); err != nil {
		s.Log.Error("Application server shutdown error", zap.Error(err))
	}
	cancelFunc()

	ctx, cancelFunc = context.WithTimeout(context.Background(), SysShutdownTimeout*time.Second)
	if err := systemServer.Shutdown(ctx); err != nil {
		s.Log.Error("System server shutdown error", zap.Error(err))
	}
	cancelFunc()
}

func (s *Server) initAppServer(router *mux.Router) *http.Server {
	return &http.Server{
		Addr:    s.config.AppAddr,
		Handler: router,
	}
}

func (s *Server) initSystemServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	return &http.Server{
		Addr:    s.config.SystemAddr,
		Handler: mux,
	}
}

func (s *Server) SetOnPanic(f func(w http.ResponseWriter, req *http.Request, err interface{})) {
	s.panicHook = f
}

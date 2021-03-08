package server

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	traceId := generateTraceId()
	ctx := contextWithMetadata(req.Context(), &Metadata{
		Log: s.Log.With(zap.Any("trace_id", traceId)),
	})
	req = req.WithContext(ctx)

	defer func() {
		if err := recover(); err != nil {

			if s.panicHook != nil {
				s.panicHook(w, req, err)
			}
		}
	}()
	s.Multiplexer.ServeHTTP(w, req)
}

func generateTraceId() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return "uuid-err"
	}
	return u.String()
}

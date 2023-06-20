package middlewares

import (
	"context"
	"log"
	"net/http"
)

type traceIDKey struct{}

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func newResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := newTraceID()
		log.Printf("[%d]%s %s\n", traceID, r.RequestURI, r.Method)

		ctx := SetTraceID(r.Context(), traceID)
		r = r.WithContext(ctx)
		rlw := newResLoggingWriter(w)

		next.ServeHTTP(rlw, r)
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}

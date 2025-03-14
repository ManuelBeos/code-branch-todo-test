package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

type responseLogger struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rw *responseLogger) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseLogger) Write(body []byte) (int, error) {
	rw.body.Write(body)
	return rw.ResponseWriter.Write(body)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var requestBody bytes.Buffer
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			requestBody.Write(bodyBytes)
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		respLogger := &responseLogger{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(respLogger, r)

		log.Printf("→ %s %s | Body: %s", r.Method, r.URL.Path, requestBody.String())
		log.Printf("← %d %s | Response: %s | Duration: %s", respLogger.statusCode, http.StatusText(respLogger.statusCode), respLogger.body.String(), time.Since(start))
	})
}

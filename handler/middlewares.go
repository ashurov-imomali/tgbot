package handler

import (
	"github.com/rs/zerolog"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"time"
)

type MiddleWare func(http.Handler) http.Handler

var z zerolog.Logger

func Use(h http.Handler, middlewares ...MiddleWare) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func init() {
	z = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}

func CheckAuthKey(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Method = http.MethodPost
		log.Println("ok")
		ip, p, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err)
		}
		i := net.ParseIP(ip)
		log.Println(i, p)
		handler.ServeHTTP(w, r)
	})

}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Установка заголовков CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovery: panic recovered: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode   int
	errorMessage string
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func LoggerWithFormatter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknow"
		}
		path := r.URL.Path
		start := time.Now()

		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rr, r)

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := rr.statusCode
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		clientIP := net.ParseIP(ip)
		clientUserAgent := r.Header.Get("User-Agent")
		z.Info().
			Interface("hostname", hostname).
			Interface("statusCode", rr.statusCode).
			Interface("latency", latency).
			Interface("clientIP", clientIP).
			Interface("method", r.Method).
			Interface("path", path).
			Interface("userAgent", clientUserAgent).Send()
		format := "ClientIP: %s | Hostname: %s | Time: %v | Method: %s | Path: %s | Status: %d | UserAgent: %s | Latency: %d"
		if statusCode > 499 {
			z.Error().Msgf(format, clientIP, hostname, stop, r.Method, path, statusCode, r.UserAgent(), latency)
		} else if statusCode > 399 {
			z.Warn().Msgf(format, clientIP, hostname, stop, r.Method, path, statusCode, r.UserAgent(), latency)
		} else {
			z.Info().Msgf(format, clientIP, hostname, stop, r.Method, path, statusCode, r.UserAgent(), latency)
		}
	})
}

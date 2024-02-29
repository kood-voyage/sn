package server

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"os"
	"social-network/pkg/jwttoken"
	"time"
)

func (s *Server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{w, http.StatusOK}
		if r.Method == http.MethodOptions {
			next.ServeHTTP(rw, r)
			return
		}

		s.logger.Printf("started %s %s ----- remote_addr:%s request_id:%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			r.Context().Value(ctxKeyRequestID),
		)
		start := time.Now()
		next.ServeHTTP(rw, r)
		s.logger.Printf("completed in %s with %d %s ----- remote_addr:%s  request_id:%s",
			time.Now().Sub(start),
			rw.code,
			http.StatusText(rw.code),
			r.RemoteAddr,
			r.Context().Value(ctxKeyRequestID),
		)
	})
}

func (s *Server) CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers dynamically based on the request's Origin header
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Allow only specific methods for actual requests
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *Server) jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(sessionName)
		if err != nil {
			s.error(w, http.StatusUnauthorized, err)
			return
		}

		// Parse the token
		alg := jwttoken.HmacSha256(os.Getenv(jwtKey))
		claims, err := alg.DecodeAndValidate(cookie.Value)
		if err != nil {
			s.error(w, http.StatusUnauthorized, err)
			return
		}

		id, err := claims.Get("user_id")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxUserID, id)))
		}
	})
}

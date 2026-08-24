package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ThanhNV121097/project-ed9b1c8e/backend/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("connect database: ", err)
	}
	defer pool.Close()

	if err := db.Migrate(ctx, pool); err != nil {
		log.Fatal("migrate database: ", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthHandler(pool))
	mux.HandleFunc("GET /v1/greeting", greetingHandler(pool))

	server := &http.Server{
		Addr:              ":" + port(),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func port() string {
	if value := os.Getenv("PORT"); value != "" {
		return value
	}
	if value := os.Getenv("APP_PORT"); value != "" {
		return value
	}
	return "8080"
}

func healthHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		var one int
		if err := pool.QueryRow(ctx, "select 1").Scan(&one); err != nil || one != 1 {
			writeError(w, http.StatusServiceUnavailable, "service_unavailable", "Service unavailable")
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte("ok\n"))
	}
}

func greetingHandler(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		var text string
		err := pool.QueryRow(ctx, "select text from greetings where id = $1", 1).Scan(&text)
		if db.IsNoRows(err) {
			writeError(w, http.StatusNotFound, "not_found", "Greeting not found")
			return
		}
		if err != nil {
			log.Println("read greeting:", err)
			writeError(w, http.StatusInternalServerError, "internal_error", "Request failed")
			return
		}

		writeJSON(w, http.StatusOK, map[string]map[string]string{
			"greeting": {"text": text},
		})
	}
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, map[string]map[string]string{
		"error": {"code": code, "message": message},
	})
}

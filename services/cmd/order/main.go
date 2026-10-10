package main

import (
	"log/slog"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
}


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)

	srv := &http.Server{
		Addr: ":8080", 
		Handler: mux,
	}

	slog.Info("server started", "addr", srv.Addr) 
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
	
}
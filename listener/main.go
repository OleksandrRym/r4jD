package main

import (
	"context"
	"listener/cfg"
	"listener/db"
	"listener/kafka"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := cfg.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Postgres
	repo, err := db.NewPostgresRepo(ctx, cfg.PostgresURL)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer repo.Close()

	// Kafka consumer
	consumer := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID, repo)
	defer consumer.Close()

	go consumer.Start(ctx)

	// health endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	go func() {
		log.Println("HTTP server listening on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("http server error: %v", err)
		}
	}()

	// graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down gracefully...")
	cancel()
}

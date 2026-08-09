package db

import (
	"context"
	"fmt"
	"listener/model"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(ctx context.Context, connString string) (*PostgresRepo, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	repo := &PostgresRepo{pool: pool}

	if err := repo.migrate(ctx); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	log.Println("Connected to Postgres successfully")
	return repo, nil
}

func (r *PostgresRepo) migrate(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS kafka_messages (
		id SERIAL PRIMARY KEY,
		msg_key VARCHAR(255),
		msg_value TEXT NOT NULL,
		topic VARCHAR(255) NOT NULL,
		partition INT NOT NULL,
		"offset" BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		UNIQUE(topic, partition, "offset")
	);`
	_, err := r.pool.Exec(ctx, query)
	return err
}

func (r *PostgresRepo) SaveMessage(ctx context.Context, msg model.KafkaMessage) error {
	query := `
	INSERT INTO kafka_messages (msg_key, msg_value, topic, partition, "offset", created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (topic, partition, "offset") DO NOTHING`

	_, err := r.pool.Exec(ctx, query,
		msg.Key, msg.Value, msg.Topic, msg.Partition, msg.Offset, time.Now())

	if err != nil {
		return fmt.Errorf("failed to insert message: %w", err)
	}
	return nil
}

func (r *PostgresRepo) Close() {
	r.pool.Close()
}

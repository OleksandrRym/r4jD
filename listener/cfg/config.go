package cfg

import (
	"os"
)

type Config struct {
	KafkaBrokers []string
	KafkaTopic   string
	KafkaGroupID string
	PostgresURL  string
}

func Load() *Config {
	return &Config{
		KafkaBrokers: []string{getEnv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")},
		KafkaTopic:   getEnv("KAFKA_TOPIC", "my-topic"),
		KafkaGroupID: getEnv("KAFKA_GROUP_ID", "go-consumer-group"),
		PostgresURL:  getEnv("POSTGRES_URL", "postgres://appuser:apppass@localhost:5432/appdb"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

package kafka

import (
	"context"
	"listener/db"
	"listener/model"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	repo   *db.PostgresRepo
}

func NewConsumer(brokers []string, topic, groupID string, repo *db.PostgresRepo) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
	})

	return &Consumer{
		reader: reader,
		repo:   repo,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	log.Println("Kafka consumer started, waiting for messages...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer stopping due to context cancellation")
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error reading message: %v", err)
				continue
			}

			kafkaMsg := model.KafkaMessage{
				Key:       string(msg.Key),
				Value:     string(msg.Value),
				Topic:     msg.Topic,
				Partition: msg.Partition,
				Offset:    msg.Offset,
			}

			if err := c.repo.SaveMessage(ctx, kafkaMsg); err != nil {
				log.Printf("failed to persist message (partition=%d offset=%d): %v",
					msg.Partition, msg.Offset, err)
				continue
			}

			log.Printf("Persisted message: key=%s partition=%d offset=%d",
				kafkaMsg.Key, kafkaMsg.Partition, kafkaMsg.Offset)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

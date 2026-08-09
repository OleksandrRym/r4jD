package model

import "time"

type KafkaMessage struct {
	ID        int64     `json:"id"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Partition int       `json:"partition"`
	Offset    int64     `json:"offset"`
	Topic     string    `json:"topic"`
	CreatedAt time.Time `json:"created_at"`
}

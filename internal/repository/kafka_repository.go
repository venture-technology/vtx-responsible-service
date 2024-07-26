package repository

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type IKafkaRepository interface {
	PublishKafkaMessage(ctx context.Context, msg string) error
}

type KafkaRepository struct {
	producer *kafka.Writer
}

func NewKafkaRepository(producer *kafka.Writer) *KafkaRepository {
	return &KafkaRepository{
		producer: producer,
	}
}

func (k *KafkaRepository) PublishKafkaMessage(ctx context.Context, msg string) error {

	message := []byte(msg)

	err := k.producer.WriteMessages(ctx, kafka.Message{
		Key:     nil,
		Value:   message,
		Time:    time.Now(),
		Headers: nil,
	})

	if err != nil {
		log.Fatalf("error to writing message on Kafka: %v", err)
	}

	return err

}

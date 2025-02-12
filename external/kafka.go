package external

import (
	"context"
	"ecommerce-order/helpers"
	"fmt"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

func (e External) ProductKafkaMessage(ctx context.Context, data []byte) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	brokers := strings.Split(helpers.GetEnv("KAFKA_BROKERS", "localhost:29092,localhost:29093,localhost:29094"), ",")
	topic := helpers.GetEnv("KAFKA_TOPIC_PAYMENT_INITIATE", "payment-initiation-topic")

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return errors.Wrap(err, "failed to comunicate with kafka brokers")
	}

	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)

	if err != nil {
		return errors.Wrap(err, "failed to produce message to kafka")
	}

	helpers.Logger.Info(fmt.Sprintf("successfuly to produce on topic %s, partition %d, offset %d", topic, partition, offset))

	return nil
}

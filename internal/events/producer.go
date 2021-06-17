package events

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type IssueEventType = uint8

const (
	Created IssueEventType = iota
	Updated
	Removed
)

type IssueEvent struct {
	Type IssueEventType         `json:"type"`
	Body map[string]interface{} `json:"body"`
}

type EventProducer interface {
	Produce(event IssueEvent) error
	Close()
}

type producer struct {
	kafkaProducer *kafka.Producer
	topic         string
}

func (p *producer) Produce(event IssueEvent) error {
	value, err := json.Marshal(event)

	if err != nil {
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          value,
	}

	return p.kafkaProducer.Produce(message, nil)
}

func (p *producer) Close() {
	p.kafkaProducer.Close()
}

func NewProducer(brokers string, topic string) (EventProducer, error) {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})

	if err != nil {
		return nil, err
	}

	return &producer{
		kafkaProducer: kafkaProducer,
		topic:         topic,
	}, nil
}

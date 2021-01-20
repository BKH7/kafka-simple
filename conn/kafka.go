package conn

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

// KafkaClient ...
type KafkaClient struct {
	producer *kafka.Producer
}

var kc = &KafkaClient{}

// NewKafkaConnection ...
func NewKafkaConnection(host string) error {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		panic(err)
	}
	kc.producer = producer
	return nil
}

// Producer ...
func Producer(topic string, msg []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := kc.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, deliveryChan)
	if err != nil {
		logrus.Error(err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		logrus.Infof("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		logrus.Infof("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
	return m.TopicPartition.Error
}

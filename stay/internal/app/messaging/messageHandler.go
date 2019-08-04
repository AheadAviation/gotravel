package messaging

import (
  "fmt"
  "github.com/confluentinc/confluent-kafka-go/kafka"
  "os"
)

type KConsumer struct {
  consumer *kafka.Consumer
  group    string
}

func (kc *KConsumer) OnMessage(callback func(*kafka.Message)) {
  for {
    select {
    case ev := <-kc.consumer.Events():
      switch e := ev.(type) {
      case kafka.AssignedPartitions:
        fmt.Fprint(os.Stderr, "%% %v\n", e)
        kc.consumer.Assign(e.Partitions)
      case kafka.RevokedPartitions:
        fmt.Fprint(os.Stderr, "%% %v\n", e)
        kc.consumer.Unassign()
      case *kafka.Message:
        fmt.Printf("%% Message on %s:\n%s\n",
          e.TopicPartition, string(e.Value))
      case kafka.PartitionEOF:
        fmt.Printf("%% Reached %v\n", e)
      case kafka.Error:
        fmt.Fprint(os.Stderr, "%% Error: %v\n", e)
      }
    }
  }
}

func New(broker, group string) (*KConsumer, error) {
  c, err := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servrers": broker,
    "group.id": group,
    "session.timeout.ms": 6000,
    "go.events.channel.enable": true,
    "go.application.rebalance.enable": true,
    "default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"}})

  if err != nil {
    fmt.Fprint(os.Stderr, "failed to create consomer: %s\n", err)
  }
  return &KConsumer{consumer: c, group: group}, nil
}

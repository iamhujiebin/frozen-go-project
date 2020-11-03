package svc

import (
	"frozen-go-project/rpc/event-rpc/internal/config"
	"github.com/Shopify/sarama"
	"github.com/tal-tech/go-zero/core/logx"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	KafkaProducer sarama.SyncProducer
}

func NewServiceContext(c config.Config) *ServiceContext {
	kc := sarama.NewConfig()
	switch c.Kafka.RequiredAcks {
	case "no_response":
		kc.Producer.RequiredAcks = sarama.NoResponse
	case "wait_for_local":
		kc.Producer.RequiredAcks = sarama.WaitForLocal
	case "wait_for_all":
		kc.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	default:
		panic("not support kafka requires")
	}
	kc.Producer.Retry.Max = c.Kafka.RetryMax // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = c.Kafka.ReturnSuccess
	kc.Producer.Return.Errors = c.Kafka.ReturnError
	kc.Net.DialTimeout = time.Second * 3

	var err error
	producer, err := sarama.NewSyncProducer(c.Kafka.Brokers, kc)
	if err != nil {
		logx.Errorf("kafka fail:%s", err.Error())
		//panic(err)
	}
	return &ServiceContext{
		Config:        c,
		KafkaProducer: producer,
	}
}

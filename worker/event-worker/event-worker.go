package main

import (
	"context"
	"flag"
	"frozen-go-project/worker/event-worker/internal/svc"
	"frozen-go-project/worker/event-worker/internal/user_event"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/zrpc"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

var configFile = flag.String("f", "etc/event-worker.yaml", "the config file")

var c = struct {
	BaseRpc  zrpc.RpcClientConf
	UserRpc  zrpc.RpcClientConf
	Log      logx.LogConf
	Brokers  []string
	Group    string
	Topics   []string
	Assignor string
	Oldest   bool
	Verbose  bool
}{}

func Init() {
	conf.MustLoad(*configFile, &c)
	if len(c.Brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}
	if len(c.Topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}
	if len(c.Group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	flag.Parse()
	logx.MustSetup(c.Log)
	Init()
	svc.InitServiceContext(c.BaseRpc, c.UserRpc)
	logx.Info("Starting a new Sarama consumer")
	if c.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	config := sarama.NewConfig()

	switch c.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", c.Assignor)
	}

	if c.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(c.Brokers, c.Group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, c.Topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	logx.Infof("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		//msg should not be handle long or easily fail
		//or should store the offset in db
		success := user_event.HandleUserEvent(message)
		if success {
			session.MarkMessage(message, "")
		} else {
			logx.Errorf("UserEvent Message fail,should store to some where: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		}
	}

	return nil
}

package main

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-worker/config"
	"github.com/witalok2/test-dev-golang-worker/internal/entity"
	gateway "github.com/witalok2/test-dev-golang-worker/internal/gateway"
	"github.com/witalok2/test-dev-golang-worker/internal/repository"
	"github.com/witalok2/test-dev-golang-worker/internal/service"
)

var serv service.Service

func init() {
	logger.New().WithContext(context.WithValue(context.Background(), entity.SERVICE, entity.SERVICE_NAME))
}

func main() {
	ctx := context.Background()
	logger.Infof("Starting execution: %v", time.Now().Format("2006-01-02 15:04:05"))

	env, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("instantiating database")
	dbWriter, err := repository.NewReaderConnection(env.Database.URI)
	if err != nil {
		logger.Fatal(err)
	}

	defer dbWriter.Close()

	logger.Info("instantiating service")
	serv = service.NewService(ctx, dbWriter)

	logger.Info("instantiating message queue")
	rabbitMQ, err := gateway.NewQueueClient(env.Queue)
	if err != nil {
		logger.Fatal(err)
	}
	defer rabbitMQ.Conn.Close()
	defer rabbitMQ.Channel.Close()

	msgs, err := rabbitMQ.Channel.Consume(
		env.QueueName, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		logger.WithError(err).Error("failed to consume a message")
		return
	}

	done := make(chan bool)

	go consumirMensagem(ctx, rabbitMQ.Channel, msgs, done)

	<-done

	logger.Info("Finished execution")
}

func consumirMensagem(ctx context.Context, ch *amqp.Channel, entrega <-chan amqp.Delivery, done chan<- bool) {
	for msg := range entrega {
		go func(m amqp.Delivery) {
			payload := entity.QueueRequest{}
			err := json.Unmarshal(m.Body, &payload)
			if err != nil {
				return
			}

			switch payload.Param {
			case entity.CREATE_CLIENT:
				err := serv.CreateClient(ctx, payload.Data)
				if err != nil {
					logger.Errorf("error on create client: %v", strings.Join([]string{payload.Data.Name, payload.Data.LastName}, " "))
				}
			case entity.UPDATE_CLIENT:
				err := serv.UpdateClient(ctx, payload.Data)
				if err != nil {
					logger.Errorf("error on update client: %v", strings.Join([]string{payload.Data.Name, payload.Data.LastName}, " "))
				}
			}
		}(msg)
	}

	done <- true
}

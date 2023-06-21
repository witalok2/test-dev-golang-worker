package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

type (
	Environment struct {
		Port string
		Database
		Queue
	}
	Database struct {
		URI string
	}
	Queue struct {
		URI       string
		QueueName string
	}
)

func LoadConfig() (*Environment, error) {
	logger.Info("Loading environment variables")

	gaeEnv := os.Getenv("GAE_ENV")
	if gaeEnv == "" {
		//Case it's running locally
		err := godotenv.Load("./config/.env")
		if err != nil {
			return nil, errors.New("error load env locally")
		}
	}

	defaultPort, ok := os.LookupEnv("PORT")
	if !ok {
		defaultPort = "8081"
	}

	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return nil, errors.New("env var isn't set: DATABASE_URL")
	}

	rabbitURI, ok := os.LookupEnv("RABBITMQ_URI")
	if !ok {
		return nil, errors.New("env var isn't set: RABBITMQ_URI")
	}

	queueName, ok := os.LookupEnv("QUEUE_NAME")
	if !ok {
		return nil, errors.New("env var isn't set: QUEUE_NAME")
	}

	logger.Info("Successfully loaded all environment variables")

	return &Environment{
		Port: defaultPort,
		Database: Database{
			URI: databaseURL,
		},
		Queue: Queue{
			URI:       rabbitURI,
			QueueName: queueName,
		},
	}, nil
}

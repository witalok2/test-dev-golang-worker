package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// Simulate the environment variables
	os.Setenv("GAE_ENV", "TEST")
	os.Setenv("PORT", "8081")
	os.Setenv("DATABASE_URL", "postgres://user:password@localhost:5432/dbname")
	os.Setenv("RABBITMQ_URI", "amqp://guest:guest@localhost:5672/")
	os.Setenv("QUEUE_NAME", "my_queue")

	defer func() {
		// Clean up the environment variables after the test
		os.Unsetenv("GAE_ENV")
		os.Unsetenv("PORT")
		os.Unsetenv("DATABASE_URL")
		os.Unsetenv("RABBITMQ_URI")
		os.Unsetenv("QUEUE_NAME")
	}()

	config, err := LoadConfig()
	assert.NoError(t, err)

	expectedConfig := &Environment{
		Port: "8081",
		Database: Database{
			URI: "postgres://user:password@localhost:5432/dbname",
		},
		Queue: Queue{
			URI:       "amqp://guest:guest@localhost:5672/",
			QueueName: "my_queue",
		},
	}

	assert.Equal(t, expectedConfig, config)
}

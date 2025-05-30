package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aftosmiros/moodu/user-service/pkg/mongo"
	"github.com/joho/godotenv"
)

type Config struct {
	Mongo      mongo.Config
	GRPCServer GRPCServer
	Redis      Redis
	Nats       Nats
	JWT        JWT
	Version    string
}

type Redis struct {
	Host string
	DB   int
}

type GRPCServer struct {
	Port                  int
	MaxRecvMsgSizeMiB     int
	MaxConnectionAge      time.Duration
	MaxConnectionAgeGrace time.Duration
}

type JWT struct {
	Secret string
}

type Nats struct {
	Host string
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	port, _ := strconv.Atoi(getEnv("GRPC_PORT", "50052"))
	maxMsg, _ := strconv.Atoi(getEnv("GRPC_MAX_MESSAGE_SIZE_MIB", "12"))
	age, _ := time.ParseDuration(getEnv("GRPC_MAX_CONNECTION_AGE", "30s"))
	ageGrace, _ := time.ParseDuration(getEnv("GRPC_MAX_CONNECTION_AGE_GRACE", "10s"))

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	cfg := &Config{
		Mongo: mongo.Config{
			Database:     getEnv("MONGO_DB", "minimoodle"),
			URI:          getEnv("MONGO_DB_URI", "mongodb://localhost:27017"),
			Username:     getEnv("MONGO_USERNAME", ""),
			Password:     getEnv("MONGO_PWD", ""),
			ReplicaSet:   getEnv("MONGO_DB_REPLICA_SET", ""),
			WriteConcern: getEnv("MONGO_WRITE_CONCERN", "majority"),
			TLSFilePath:  getEnv("MONGO_TLS_FILE_PATH", ""),
			TLSEnable:    getEnv("MONGO_TLS_ENABLE", "false") == "true",
		},
		GRPCServer: GRPCServer{
			Port:                  port,
			MaxRecvMsgSizeMiB:     maxMsg,
			MaxConnectionAge:      age,
			MaxConnectionAgeGrace: ageGrace,
		},
		Redis: Redis{
			Host: getEnv("REDIS_HOST", "localhost:6379"),
			DB:   redisDB,
		},
		Nats: Nats{
			Host: getEnv("NATS_HOST", "nats://localhost:4222"),
		},
		JWT: JWT{
			Secret: getEnv("JWT_SECRET", ""),
		},
		Version: getEnv("VERSION", "dev"),
	}

	return cfg
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

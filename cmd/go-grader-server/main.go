package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MShilenko/go-grader-server/configs"
	"github.com/MShilenko/go-grader-server/internal/postgre"
	"github.com/MShilenko/go-grader-server/internal/redis"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	logger := log.New(os.Stdout, "logger: ", log.Lshortfile)

	cfg, err := configs.NewConfig()
	if err != nil {
		logger.Fatalf("get config %v")
		os.Exit(1)
	}

	if err := cfg.Print(); err != nil {
		logger.Fatalf("read config %v", err)
		os.Exit(1)
	}
	logger.Println("Config read successfull")

	pgxConnect, err := postgre.GetPostgreConnect(cfg)
	if err != nil {
		logger.Fatalf("postgres start %v", err)
		os.Exit(1)
	}
	defer pgxConnect.Close()

	redisClient, err := redis.GetRedisClient(cfg)
	if err != nil {
		logger.Fatalf("redis start %v", err)
		os.Exit(1)
	}
	defer redisClient.Close()
}

func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

package main

import (
	"log"
	"os"

	"github.com/MShilenko/go-grader-server/configs"
	"github.com/MShilenko/go-grader-server/internal/postgre"
)

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	logger := log.New(os.Stdout, "logger: ", log.Lshortfile)

	cfg, err := configs.NewConfig(
	if err != nil {
		logger.Fatalf("get config %v", err)
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
}

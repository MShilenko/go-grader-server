package postgre

import (
	"database/sql"
	"fmt"

	"github.com/MShilenko/go-grader-server/configs"
)

func GetPostgreConnect(cfg *configs.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.DatabaseName,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Secure,
	))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.Postgres.MaxConnsPool)

	return db, nil
}

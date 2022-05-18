package pg

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4"
	"log"
)

var pgCfg struct {
	Port     uint   `env:"PG_PORT" envDefault:"5432"`
	Host     string `env:"PG_HOST" envDefault:"localhost"`
	Username string `env:"PG_USER" envDefault:"postgres"`
	Password string `env:"PG_PASS" envDefault:"password"`
	Db       string `env:"DB_NAME" envDefault:"postgres"`
}

func init() {
	err := env.Parse(&pgCfg)
	if err != nil {
		log.Fatal("Error parsing db config: ", err.Error())
	}
}

func NewPgConnect(ctx context.Context) (*pgx.Conn, error) {
	pgUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgCfg.Username, pgCfg.Password, pgCfg.Host, pgCfg.Port, pgCfg.Db)
	conn, err := pgx.Connect(ctx, pgUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

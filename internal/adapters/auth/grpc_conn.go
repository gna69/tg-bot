package auth

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var cfg struct {
	GrpcServerAddr string `env:"GRPC_SRV_ADDR" envDefault:"localhost"`
	GrpcPort       string `env:"GRPC_PORT" envDefault:"8080"`
}

func init() {
	if err := env.Parse(&cfg); err != nil {
		log.Error().Msgf("Error parsing grpc cli config: %s", err.Error())
		return
	}
}

func NewGrpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.GrpcServerAddr, cfg.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

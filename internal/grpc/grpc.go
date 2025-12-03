package grpc

import (
	"api-gateway/internal/config"

	"github.com/rs/zerolog/log"
	authpb "github.com/viktoralyoshin/playhub-proto/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthClient authpb.AuthServiceClient

func Init(cfg *config.Config) {
	authServiceAddr := cfg.AuthServiceAddr

	authServiceConn, err := grpc.NewClient(authServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("failed to initialize grpc connection: %v", err)
	}

	AuthClient = authpb.NewAuthServiceClient(authServiceConn)
}

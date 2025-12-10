package grpc

import (
	"api-gateway/internal/config"

	"github.com/rs/zerolog/log"
	authpb "github.com/viktoralyoshin/playhub-proto/gen/go/auth"
	gamepb "github.com/viktoralyoshin/playhub-proto/gen/go/games"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthClient authpb.AuthServiceClient
var GamesClient gamepb.GameServiceClient

func Init(cfg *config.Config) {
	authServiceConn := connect(cfg.AuthServiceAddr)
	gamesServiceConn := connect(cfg.GamesServiceAddr)

	AuthClient = authpb.NewAuthServiceClient(authServiceConn)
	GamesClient = gamepb.NewGameServiceClient(gamesServiceConn)
}

func connect(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Msgf("failed to initialize grpc connection to %s: %v", addr, err)
	}

	return conn
}

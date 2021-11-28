package cnm_film

import (
	"context"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/utils/logger"
	filmApi "github.com/ozonmp/omp-bot/pb/github.com/ozonmp/cnm-film-api/pkg/cnm-film-api"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, host string) (filmApi.CnmFilmApiServiceClient, error) {
	log := logger.LoggerFromContext(ctx)
	log.Info().Msg(fmt.Sprintf("Start establishing grpc connection with host '%s'", host))

	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if err != nil {
		log.Debug().Err(err).Msg("Can't establish grpc connection")
		return nil, err
	}

	client := filmApi.NewCnmFilmApiServiceClient(conn)

	return client, nil
}

package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/go/proto"
)

// NationalForestsGet retrieves all national forests or one by ID
func (s *API) NationalForestsGet(ctx context.Context, request *pb.NationalForestsGetRequest) (*pb.NationalForestsGetResponse, error) {
	forests, err := s.dbClient.NationalForestsGet(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forests")
	}

	response := &pb.NationalForestsGetResponse{
		Forests: forests,
	}

	return response, nil
}

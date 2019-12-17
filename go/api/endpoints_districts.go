package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/go/proto"
)

// GetAllDistricts retrieves all districts or one by ID
func (s *API) DistrictsGet(ctx context.Context, request *pb.DistrictsGetRequest) (*pb.DistrictsGetResponse, error) {
	districts, err := s.dbClient.DistrictsGet(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching districts")
	}

	response := &pb.DistrictsGetResponse{
		Districts: districts,
	}

	return response, nil
}

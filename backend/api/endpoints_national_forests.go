package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetNationalForest retrieves a forest given an id
func (s *API) GetNationalForest(ctx context.Context, request *pb.GetNationalForestRequest) (*pb.GetNationalForestResponse, error) {
	forest, err := s.dbClient.GetNationalForest(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forest")
	}

	response := &pb.GetNationalForestResponse{
		Forest: &pb.NationalForest{
			Id:      forest.ID,
			Name:    forest.Name,
			Website: forest.Website.String,
		},
	}

	return response, nil
}

// CreateNationalForest creates a new national forest in the database
func (s *API) CreateNationalForest(ctx context.Context, request *pb.CreateNationalForestRequest) (response *pb.CreateNationalForestResponse, err error) {
	// TODO write query to save forest

	return response, nil
}

// GetAllNationalForests retrieves all national forests
func (s *API) GetAllNationalForests(ctx context.Context, request *pb.GetAllNationalForestsRequest) (*pb.GetAllNationalForestsResponse, error) {
	forests, err := s.dbClient.GetAllNationalForests(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forest")
	}

	locs := make([]*pb.NationalForest, 0, len(forests))
	for _, f := range forests {
		forest := &pb.NationalForest{
			Id:      f.ID,
			Name:    f.Name,
			Website: f.Website.String,
		}

		locs = append(locs, forest)
	}

	response := &pb.GetAllNationalForestsResponse{
		Forests: locs,
	}

	return response, nil
}

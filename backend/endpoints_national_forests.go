package backend

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetNationalForest retrieves a forest given an id
func (s *Service) GetNationalForest(ctx context.Context, request *pb.GetNationalForestRequest) (response *pb.GetNationalForestResponse, err error) {
	forest, err := s.dbClient.GetNationalForest(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forest")
	}

	response = &pb.GetNationalForestResponse{
		Forest: &pb.NationalForest{
			Id:      forest.ID,
			Name:    forest.Name,
			Website: forest.Website.String,
		},
	}

	return response, nil
}

// CreateNationalForest creates a new national forest in the database
func (s *Service) CreateNationalForest(ctx context.Context, request *pb.CreateNationalForestRequest) (response *pb.CreateNationalForestResponse, err error) {
	// TODO write query to save forest

	return response, nil
}

// GetAllNationalForests retrieves all national forests
func (s *Service) GetAllNationalForests(ctx context.Context, request *pb.GetAllNationalForestsRequest) (response *pb.GetAllNationalForestsResponse, err error) {
	forests, err := s.dbClient.GetAllNationalForests(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forest")
	}

	for _, f := range forests {
		forest := &pb.NationalForest{
			Id:      f.ID,
			Name:    f.Name,
			Website: f.Website.String,
		}

		response.Forests = append(response.Forests, forest)
	}

	return response, nil
}

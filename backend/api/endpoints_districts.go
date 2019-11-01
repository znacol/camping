package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetDistrict retrieves a district given an id
func (s *API) GetDistrict(ctx context.Context, request *pb.GetDistrictRequest) (*pb.GetDistrictResponse, error) {
	district, err := s.dbClient.GetDistrict(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching district")
	}

	response := &pb.GetDistrictResponse{
		District: district,
	}

	return response, nil
}

// GetAllDistricts retrieves all districts
func (s *API) GetAllDistricts(ctx context.Context, request *pb.GetAllDistrictsRequest) (*pb.GetAllDistrictsResponse, error) {
	districts, err := s.dbClient.GetAllDistricts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetching districts")
	}

	response := &pb.GetAllDistrictsResponse{
		Districts: districts,
	}

	return response, nil
}

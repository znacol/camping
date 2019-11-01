package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetDistrict retrieves a district given an id
func (s *API) GetDistrict(ctx context.Context, request *pb.GetDistrictRequest) (*pb.GetDistrictResponse, error) {
	forest, err := s.dbClient.GetDistrict(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching district")
	}

	response := &pb.GetDistrictResponse{
		District: &pb.District{
			Id:               forest.ID,
			NationalForestId: forest.NationalForestID,
			Name:             forest.Name,
			MapLocation:      forest.MapLocation.String,
		},
	}

	return response, nil
}

// GetAllDistricts retrieves all districts
func (s *API) GetAllDistricts(ctx context.Context, request *pb.GetAllDistrictsRequest) (*pb.GetAllDistrictsResponse, error) {
	districts, err := s.dbClient.GetAllDistricts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetching districts")
	}

	locs := make([]*pb.District, 0, len(districts))
	for _, d := range districts {
		district := &pb.District{
			Id:               d.ID,
			NationalForestId: d.NationalForestID,
			Name:             d.Name,
			MapLocation:      d.MapLocation.String,
		}

		locs = append(locs, district)
	}

	response := &pb.GetAllDistrictsResponse{
		Districts: locs,
	}

	return response, nil
}

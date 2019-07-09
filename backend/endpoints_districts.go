package backend

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetDistrict retrieves a district given an id
func (s *Service) GetDistrict(ctx context.Context, request *pb.GetDistrictRequest) (response *pb.GetDistrictResponse, err error) {
	forest, err := s.dbClient.GetDistrict(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "fetching district")
	}

	response = &pb.GetDistrictResponse{
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
func (s *Service) GetAllDistricts(ctx context.Context, request *pb.GetAllDistrictsRequest) (response *pb.GetAllDistrictsResponse, err error) {
	districts, err := s.dbClient.GetAllDistricts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fetching national forest")
	}

	for _, d := range districts {
		district := &pb.District{
			Id:               d.ID,
			NationalForestId: d.NationalForestID,
			Name:             d.Name,
			MapLocation:      d.MapLocation.String,
		}

		response.Districts = append(response.Districts, district)
	}

	return response, nil
}

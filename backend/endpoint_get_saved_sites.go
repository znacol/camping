package backend

import (
	"log"

	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

func (s *Service) GetAllSites(ctx context.Context, request *pb.GetAllSitesRequest) (response *pb.GetAllSitesResponse, err error) {
	sites, err := s.dbClient.GetSites(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting sites")
	}

	locs := make([]*pb.Site, 0, len(sites))
	for _, s := range sites {
		site := &pb.Site{
			Id:               s.ID,
			Latitude:         float32(s.Latitude),
			Longitude:        float32(s.Longitude),
			NationalForestId: s.NationalForestID,
			DistrictId:       s.DistrictID,
			Altitude:         s.Altitude,
			Notes:            s.Notes,
		}
		locs = append(locs, site)
	}

	response = &pb.GetAllSitesResponse{
		Sites: locs,
	}

	return response, nil
}

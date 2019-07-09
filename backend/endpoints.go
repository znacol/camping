package backend

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetAllSites retrieves all sites and their info
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
			Notes:            s.Notes.String,
		}
		locs = append(locs, site)
	}

	response = &pb.GetAllSitesResponse{
		Sites: locs,
	}

	return response, nil
}

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

// CreateSite creates a new site in the database
func (s *Service) CreateSite(ctx context.Context, request *pb.CreateSiteRequest) (response *pb.CreateSiteResponse, err error) {
	err = s.dbClient.CreateSite(ctx, request.Site.Latitude, request.Site.Longitude, request.Site.NationalForestId, request.Site.DistrictId, request.Site.Altitude, request.Site.Notes)
	if err != nil {
		return nil, errors.Wrap(err, "creating site")
	}

	return &pb.CreateSiteResponse{}, nil
}

// CreateNationalForest creates a new national forest in the database
func (s *Service) CreateNationalForest(ctx context.Context, request *pb.CreateNationalForestRequest) (response *pb.CreateNationalForestResponse, err error) {
	// TODO write query to save forest

	return &pb.CreateNationalForestResponse{}, nil
}

package api

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
)

// GetAllSites retrieves all sites and their info
func (s *API) GetAllSites(ctx context.Context, request *pb.GetAllSitesRequest) (*pb.GetAllSitesResponse, error) {
	sites, err := s.dbClient.GetSites(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting sites")
	}

	response := &pb.GetAllSitesResponse{
		Sites: sites,
	}

	return response, nil
}

// CreateSite creates a new site in the database
func (s *API) CreateSite(ctx context.Context, request *pb.CreateSiteRequest) (*pb.CreateSiteResponse, error) {
	err := s.dbClient.CreateSite(ctx, request.Site.Latitude, request.Site.Longitude, request.Site.NationalForestId, request.Site.DistrictId, request.Site.Altitude, request.Site.Notes)
	if err != nil {
		return nil, errors.Wrap(err, "creating site")
	}

	return &pb.CreateSiteResponse{}, nil
}

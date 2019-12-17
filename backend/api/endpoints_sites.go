package api

import (
	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
	"golang.org/x/net/context"
)

// SitesGet retrieves all sites and or one by ID
func (s *API) SitesGet(ctx context.Context, request *pb.SitesGetRequest) (*pb.SitesGetResponse, error) {
	sites, err := s.dbClient.SitesGet(ctx, request.Id)
	if err != nil {
		return nil, errors.Wrap(err, "getting sites")
	}

	response := &pb.SitesGetResponse{
		Sites: sites,
	}

	return response, nil
}

// SiteUpsert creates or updates a site in the database
func (s *API) SiteUpsert(ctx context.Context, request *pb.SiteUpsertRequest) (*pb.SiteUpsertResponse, error) {
	site, err := s.dbClient.SiteUpsert(ctx, request.Latitude, request.Longitude, request.NationalForestId, request.DistrictId, request.Altitude, request.Notes)
	if err != nil {
		return nil, errors.Wrap(err, "upserting site")
	}


	return &pb.SiteUpsertResponse{Site: site}, nil
}

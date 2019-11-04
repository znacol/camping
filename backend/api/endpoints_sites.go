package api

import (
	"golang.org/x/net/context"
	"log"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/backend/proto"
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

	log.Printf("request: %+v", request)
	log.Printf("response: %+v", response)

	return response, nil
}

// SiteUpsert creates or updates a site in the database
func (s *API) SiteUpsert(ctx context.Context, request *pb.SiteUpsertRequest) (*pb.SiteUpsertResponse, error) {
	// TODO: update db call to handle upsert functionality
	err := s.dbClient.SiteUpsert(ctx, request.Site.Latitude, request.Site.Longitude, request.Site.NationalForestId, request.Site.DistrictId, request.Site.Altitude, request.Site.Notes)
	if err != nil {
		return nil, errors.Wrap(err, "upserting site")
	}

	// TODO: return created site
	return &pb.SiteUpsertResponse{}, nil
}

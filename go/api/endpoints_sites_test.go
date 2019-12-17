package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	pb "github.com/znacol/camping/go/proto"
)

func TestSitesGet(t *testing.T) {
	t.Parallel()

	// Get api and mocks
	api, mockCtrl, mockDB := newTestAPI(t)

	// Create parameters
	req := &pb.SitesGetRequest{}

	// Set mock expectations and return values
	expectation := &pb.SitesGetResponse{
	}
	mockDB.EXPECT().SitesGet(
		gomock.Any(), // ctx
		uint64(0),
	).Return(
		expectation.Sites,
		nil, // err
	)

	// Call method
	resp, err := api.SitesGet(context.Background(), req)
	if err != nil {
		t.Error("Expected no error: ", err)
	}

	// Confirm expectations
	if !proto.Equal(resp, expectation) {
		t.Error("Expected: ", expectation, "; Got: ", resp)
	}

	// Teardown and check mock controller
	teardown(t, mockCtrl)
}

func TestSiteUpsert(t *testing.T) {
	t.Parallel()

	// Get api and mocks
	api, mockCtrl, mockDB := newTestAPI(t)

	// Create parameters
	req := &pb.SiteUpsertRequest{
		Latitude: 0,
		Longitude: 0,
		NationalForestId: 1,
		DistrictId: 1,
		Altitude: 1000,
		Notes: "test notes",
	}

	// Set mock expectations and return values
	expectation := &pb.SiteUpsertResponse{
		Site: &pb.Site{
			Latitude: req.Latitude,
			Longitude: req.Longitude,
			NationalForestId: req.NationalForestId,
			DistrictId:req.DistrictId,
			Altitude: req.Altitude,
			Notes: req.Notes,
		},
	}
	mockDB.EXPECT().SiteUpsert(
		gomock.Any(), // ctx
		expectation.Site.Latitude,
		expectation.Site.Longitude,
		expectation.Site.NationalForestId,
		expectation.Site.DistrictId,
		expectation.Site.Altitude,
		expectation.Site.Notes,
	).Return(
		expectation.Site,
		nil, // err
	)

	// Call method
	resp, err := api.SiteUpsert(context.Background(), req)
	if err != nil {
		t.Error("Expected no error: ", err)
	}

	expectation.Site.Id = resp.Site.Id
	expectation.Site.CreatedAt = resp.Site.CreatedAt

	// Confirm expectations
	if !proto.Equal(resp, expectation) {
		t.Error("Expected: ", expectation, "; Got: ", resp)
	}

	// Teardown and check mock controller
	teardown(t, mockCtrl)
}

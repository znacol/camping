package api

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	pb "github.com/znacol/camping/backend/proto"
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

// TODO: site upsert
// TODO: singular site get

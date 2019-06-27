package main

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	pb "github.com/znacol/camping/api/proto"
)

func (s *server) GetSavedSites(c context.Context, request *pb.GetSavedSitesRequest) (response *pb.GetSavedSitesResponse, err error) {
	locations, err := svc.dbClient.GetSavedSites(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "getting s")
	}

	response = &pb.GetSavedSitesResponse{
		Message: "Hello world",
	}

	return response, nil
}

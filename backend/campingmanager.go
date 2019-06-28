package backend

import (
	"github.com/znacol/camping/backend/db"
	pb "github.com/znacol/camping/backend/proto"
)

var _ pb.CampingServiceServer = &Service{}

type Service struct {
	dbClient *db.DB
}

func New(db *db.DB) *Service {
	return &Service{
		dbClient: db,
	}
}

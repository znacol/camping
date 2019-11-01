package db

import (
	"context"

	pb "github.com/znacol/camping/backend/proto"
)

// GetSites returns all campsites
func (d *DB) GetSites(ctx context.Context) ([]*pb.Site, error) {
	list := []*pb.Site{}
	err := d.dbClient.SelectContext(ctx, &list, `
		SELECT *
		FROM site
	`)

	return list, err
}

// GetNationalForest retrieves a forest given an ID
func (d *DB) GetNationalForest(ctx context.Context, id int64) (*pb.NationalForest, error) {
	nf := &pb.NationalForest{}
	err := d.dbClient.GetContext(ctx, &nf, `
		SELECT *
		FROM national_forest
		WHERE id = ?`,
		id,
	)

	return nf, err
}

// GetAllNationalForests retrieves all national forests
func (d *DB) GetAllNationalForests(ctx context.Context) ([]*pb.NationalForest, error) {
	nf := []*pb.NationalForest{}
	err := d.dbClient.SelectContext(ctx, &nf, `
		SELECT *
		FROM national_forest`,
	)

	return nf, err
}

// GetDistrict retrieves a district given an ID
func (d *DB) GetDistrict(ctx context.Context, id int64) (*pb.District, error) {
	district := &pb.District{}
	err := d.dbClient.GetContext(ctx, &district, `
		SELECT *
		FROM district
		WHERE id = ?`,
		id,
	)

	return district, err
}

// GetAllDistricts retrieves all districts
func (d *DB) GetAllDistricts(ctx context.Context) ([]*pb.District, error) {
	districts := []*pb.District{}
	err := d.dbClient.SelectContext(ctx, &districts, `
		SELECT *
		FROM district`,
	)

	return districts, err
}

// CreateSite saves a new site
func (d *DB) CreateSite(ctx context.Context, latitude float32, longitude float32, nationalForestID int64, districtID int64, altitude int64, notes string) error {
	_, err := d.dbClient.ExecContext(ctx, `
		INSERT INTO site (latitude, longitude, national_forest_id, district_id, altitude, notes)
			VALUES ($1, $2, $3, $4, $5, $6)
	`,
		latitude,
		longitude,
		nationalForestID,
		districtID,
		altitude,
		notes)

	return err
}

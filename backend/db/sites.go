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
func (d *DB) GetNationalForest(ctx context.Context, id int64) (NationalForest, error) {
	nf := NationalForest{}
	err := d.dbClient.GetContext(ctx, &nf, `
		SELECT *
		FROM national_forest
		WHERE id = ?`,
		id,
	)

	return nf, err
}

// GetAllNationalForests retrieves all national forests
func (d *DB) GetAllNationalForests(ctx context.Context) ([]NationalForest, error) {
	nf := []NationalForest{}
	err := d.dbClient.SelectContext(ctx, &nf, `
		SELECT *
		FROM national_forest`,
	)

	return nf, err
}

// GetDistrict retrieves a district given an ID
func (d *DB) GetDistrict(ctx context.Context, id int64) (District, error) {
	district := District{}
	err := d.dbClient.GetContext(ctx, &district, `
		SELECT *
		FROM district
		WHERE id = ?`,
		id,
	)

	return district, err
}

// GetAllDistricts retrieves all districts
func (d *DB) GetAllDistricts(ctx context.Context) ([]District, error) {
	districts := []District{}
	err := d.dbClient.SelectContext(ctx, &districts, `
		SELECT *
		FROM district`,
	)

	return districts, err
}

// CreateSite saves a new site
func (d *DB) CreateSite(ctx context.Context, latitude float32, longitude float32, nationalForestID int64, districtID int64, altitude int64, notes string) error {
	_, err := d.dbClient.ExecContext(ctx, `
		INSERT INTO site
			SET
			latitude = ?,
			longitude = ?,
			national_forest_id = ?,
			district_id = ?,
			altitude = ?,
			notes = ?
	`,
		latitude,
		longitude,
		nationalForestID,
		districtID,
		altitude,
		notes)

	return err
}

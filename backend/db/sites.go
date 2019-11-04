package db

import (
	"context"
	"github.com/pkg/errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/znacol/camping/backend/proto"
)

// GetSites returns all campsites
func (d *DB) SitesGet(ctx context.Context, id uint64) ([]*pb.Site, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("site")

	if id > 0 {
		query = query.Where("id = ?", id)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.WithMessage(err, "Unable to convert query to sql")
	}

	list := []*pb.Site{}
	err = d.dbClient.SelectContext(ctx, &list, sql, args...)

	return list, err
}

// NationalForestsGet retrieves all national forests
func (d *DB) NationalForestsGet(ctx context.Context, id uint64) ([]*pb.NationalForest, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("national_forest")

	if id > 0 {
		query = query.Where("id = ?", id)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.WithMessage(err, "Unable to convert query to sql")
	}

	list := []*pb.NationalForest{}
	err = d.dbClient.SelectContext(ctx, &list, sql, args...)

	return list, err
}

// DistrictsGet retrieves all districts
func (d *DB) DistrictsGet(ctx context.Context, id uint64) ([]*pb.District, error) {
	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("district")

	log.Printf("district", id)

	if id > 0 {
		query = query.Where("id = ?", id)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, errors.WithMessage(err, "Unable to convert query to sql")
	}

	list := []*pb.District{}
	err = d.dbClient.SelectContext(ctx, &list, sql, args...)

	return list, err
}

// CreateSite saves a new site
// TODO: add upsert functionality
func (d *DB) SiteUpsert(ctx context.Context, latitude float32, longitude float32, nationalForestID uint64, districtID uint64, altitude uint64, notes string) error {
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

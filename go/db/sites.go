package db

import (
	"context"
	"github.com/pkg/errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	pb "github.com/znacol/camping/go/proto"
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

// SiteUpsert updates or creates a site
func (d *DB) SiteUpsert(ctx context.Context, latitude float32, longitude float32, nationalForestID uint64, districtID uint64, altitude uint64, notes string) (*pb.Site, error) {
	var lastInsertId uint64
	err := d.dbClient.QueryRow(`
		INSERT INTO site (latitude, longitude, national_forest_id, district_id, altitude, notes)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (latitude, longitude) DO UPDATE
			SET national_forest_id = $3,
			    district_id = $4,
			    altitude = $5,
			    notes = $6
		RETURNING id
	`,
		latitude,
		longitude,
		nationalForestID,
		districtID,
		altitude,
		notes).Scan(&lastInsertId)


	// Get site
	sites, err := d.SitesGet(ctx, lastInsertId)
	if err != nil {
		return nil, err
	}
	if len(sites) != 1 {
		log.Panic("Unable to retrieve site that was just inserted")
	}

	return sites[0], nil
}

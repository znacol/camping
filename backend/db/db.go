package db

import (
	"context"
	"database/sql"
	"net"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	dbClient *sqlx.DB
}

// Site represents a single campsite
type Site struct {
	ID               int64          `db:"id" json:"id"`
	Latitude         float64        `db:"latitude" json:"latitude"`
	Longitude        float64        `db:"longitude" json:"longitude"`
	NationalForestID int64          `db:"national_forest_id" json:"national_forest_id"`
	DistrictID       int64          `db:"district_id" json:"district_id"`
	Altitude         int64          `db:"altitude" json:"altitude"`
	Notes            sql.NullString `db:"notes" json:"notes"`
}

type NationalForest struct {
	ID      int64          `db:"id" json:"id"`
	Name    string         `db:"name" json:"name"`
	Website sql.NullString `db:"website" json:"website"`
}

type District struct {
	ID               int64          `db:"id" json:"id"`
	NationalForestID int64          `db:"national_forest_id" json:"national_forest_id"`
	Name             string         `db:"name" json:"name"`
	MapLocation      sql.NullString `db:"map_location" json:"map_location"`
}

// New attempts to instantiate a connection to the mysql database
func New(username, password, schemaName, host, port string) (*DB, error) {
	dbConfig := mysql.NewConfig()
	dbConfig.User = username
	dbConfig.Passwd = password
	dbConfig.Net = "tcp"
	dbConfig.Addr = net.JoinHostPort(host, port)
	dbConfig.DBName = schemaName
	dbConfig.ClientFoundRows = true
	dbConfig.ParseTime = true
	dbConfig.Params = map[string]string{"charset": "utf8mb4"}

	conn, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &DB{
		dbClient: sqlx.NewDb(conn, "mysql"),
	}, nil
}

// GetSites returns all campsites
func (d *DB) GetSites(ctx context.Context) ([]Site, error) {
	list := []Site{}
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

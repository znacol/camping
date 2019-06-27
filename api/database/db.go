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
	ID               int64   `db:"id" json:"id"`
	Latitude         float64 `db:"latitude" json:"latitude"`
	Longitude        float64 `db:"longitude" json:"longitude"`
	NationalForestID int32   `db:"national_forest_id" json:"national_forest_id"`
	DistrictID       int32   `db:"district_id" json:"district_id"`
	Altitude         int32   `db:"altitude" json:"altitude"`
	Notes            string  `db:"notes" json:"notes"`
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

// GetSavedSites returns all campsites
func (d *DB) GetSavedSites(ctx context.Context) ([]Site, error) {
	list := []Site{}
	err := d.dbClient.SelectContext(ctx, &list, `
		SELECT *
		FROM site
	`)

	return list, err
}

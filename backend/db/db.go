package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DavidHuie/gomigrate"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq" // Postgres driver

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


// New returns an instantiated DB or an error if the database
// connection parameters don't work or it can not connect.
func New(username, password, schemaName, host, port string) (*DB, error) {

	dbConfig := postgresConfig(username, password, schemaName, host, port)

	postgresDB, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	// `Open` does not directly open connection, so ensure connection established
	if err = postgresDB.Ping(); err != nil {
		return nil, err
	}

	// Golang SQL Package and the database driver do not close connections fast enough after queries end.
	// If max open connections is not set, it can keep spinning up new connections instead of
	// waiting for old connections to return back to the pool, causing "Too Many Connections" error.
	postgresDB.SetMaxOpenConns(500)
	postgresDB.SetMaxIdleConns(10)

	return newFromConn(postgresDB), nil
}

// newFromConn creates the DB struct
func newFromConn(postgresDBConn *sql.DB) *DB {

	postgresDB := sqlx.NewDb(postgresDBConn, "postgres")

	// Use the json tag instead of the db tag
	postgresDB.Mapper = reflectx.NewMapperFunc("json", func(str string) string {
		return str
	})

	return &DB{dbClient: postgresDB}
}

// postgresConfig sets up our postgres client configuration
func postgresConfig(user, password, schemaName, host, port string) string {
	return fmt.Sprintf(`
			host=%s
			port=%s
			user=%s
			password=%s
			dbname=%s
			sslmode=disable`,
		host,
		port,
		user,
		password,
		schemaName,
	)
}

// Migrate will perform any necessary migrations on the database.  The location
// of the migration directory will need to be passed to this, since it will
// probably be manually copied to the server
func (d *DB) Migrate(migrateDir string) error {
	migrator, err := gomigrate.NewMigrator(d.dbClient.DB, gomigrate.Postgres{}, migrateDir)
	if err != nil {
		return err
	}
	return migrator.Migrate()
}

// Close permanently closes the database connection
func (d *DB) Close() error {
	return d.dbClient.Close()
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

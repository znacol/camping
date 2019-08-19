package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/znacol/camping/backend/db"
)

// envConfig represents all environment/configuration needed for this application
type envConfig struct {
	DbUser           string `default:"root" split_words:"true"`
	DbPassword       string `default:"password" split_words:"true"`
	DbHost           string `default:"db" split_words:"true"`
	DbPort           string `default:"3306" split_words:"true"`
	DbName           string `default:"camping" split_words:"true"`
	DbMigrationsPath string `default:"/go/src/github.com/znacol/camping/backend/database/migrations" split_words:"true"`
}

func main() {
	env := &envConfig{}

	if err := envconfig.Process("", env); err != nil {
		log.WithError(err).Fatal("Unable to process environment config")
	}

	log := log.WithFields(log.Fields{
		"user": env.DbUser,
		"name": env.DbName,
		"host": env.DbHost,
		"port": env.DbPort,
		"path": env.DbMigrationsPath,
	})

	dbConn, err := db.New(env.DbUser, env.DbPassword, env.DbName, env.DbHost, env.DbPort)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}

	log.Info("Attempting database migrations")
	err = dbConn.Migrate(env.DbMigrationsPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to run migrations")
	}
	log.Info("Successfully applied migrations")
}


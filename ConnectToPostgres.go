package ms_util

import (
	"fmt"
	"github.com/felixcolaci/ms-util/logger"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"time"
)

func ConnectToPostgres(cfg *PostgresConfig) *sqlx.DB {

	dbName := cfg.Database
	dbUser := cfg.Username
	dbPass := cfg.Password
	dbPort := cfg.Port
	dbHost := cfg.Host
	dbSsl := cfg.UseSsl
	initSchema := cfg.ReinitSchema

	sslMode := "disable"
	if dbSsl {
		sslMode = "enable"
	}
	var dbString string
	if len(dbPass) > 0 {
		dbString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
			sslMode)
	} else {
		dbString = fmt.Sprintf("postgres://%v@%v:%v/%v?sslmode=%v",
			dbUser,
			dbHost,
			dbPort,
			dbName,
			sslMode)
	}

	logger.Boot(fmt.Sprintf("postgres connection attempt to: %v", dbString))

	//connect to db
	db := sqlx.MustConnect("postgres", dbString)
	err := db.Ping()
	if err != nil {
		logger.Err("error connecting to database: " + err.Error())
	}
	//Configure db
	db.SetMaxIdleConns(cfg.MaxIdleCon)
	db.SetMaxOpenConns(cfg.MaxCon)
	db.SetConnMaxLifetime(time.Minute * time.Duration(cfg.MaxConLifetime))

	migrations := &migrate.FileMigrationSource{
		Dir: "./../shared/data/resources/db/migrations",
	}

	if initSchema {
		_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Down)
		if err != nil {
			panic(err)
		}
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	logger.Boot(fmt.Sprintf("applied %d migrations", n))

	return db

}

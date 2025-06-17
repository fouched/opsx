package driver

import (
	"database/sql"
	"time"
)

// ConnectSQL creates database pool
func ConnectSQL(dsn string) (*sql.DB, error) {
	d, err := getConn(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(2)
	d.SetConnMaxLifetime(5 * time.Minute)

	//dbConn.SQL = d

	return d, nil
}

// NewDatabase creates a new database connection for the application
func getConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// test database connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

package startup

import (
	"database/sql"
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/config"
	"github.com/creative-snails/phisio-log-backend-go/internal/db"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitializeDB(dbc config.DatabaseConfig) (*db.Queries, error) {
	log.Infof("Attempting to connect to database at %s:%d", dbc.Host, dbc.Port)
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", 
		dbc.Host, 
		dbc.Port, 
		dbc.User, 
		dbc.Password,
		dbc.Dbname,
		dbc.Sslmode,
	)

	address := fmt.Sprintf("%s:%d", dbc.Host, dbc.Port)

    dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
	    return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := RunMigrations(dbConn); err != nil {
		return nil, fmt.Errorf("error running migrations: %w", err)
	}

	// Verify connextion
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Infof("DB starting on %s...", address)

	queries := db.New(dbConn)
	if queries == nil {
		return nil, fmt.Errorf("failed to create queries object")
	}

	return queries, nil
}


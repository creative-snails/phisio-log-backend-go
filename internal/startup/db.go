package startup

import (
	"database/sql"
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/config"
	"github.com/creative-snails/phisio-log-backend-go/internal/db"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitializeDB() (*db.Queries, error) {
	// Load configuration
	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}


    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", 
		config.Database.Host, 
		config.Database.Port, 
		config.Database.User, 
		config.Database.Password,
		config.Database.Dbname,
		config.Database.Sslmode,
	)

	address := fmt.Sprintf("%s:%d", config.Database.Host, config.Database.Port)

    dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
	    return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verify connextion
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Infof("DB starting on %s...", address)

	queries := db.New(dbConn)

	return queries, nil
}


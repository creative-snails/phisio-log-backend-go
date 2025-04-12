package startup

import (
	"database/sql"
	"fmt"

	"github.com/creative-snails/phisio-log-backend-go/config"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitializeDB() {
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

    db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Verify connextion
	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}

	log.Infof("DB starting on %s...", address)
}


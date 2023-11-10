package adapter

import (
	"os"
	"time"

	"github.com/electivetechnology/utility-library-go/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MAX_ATTEMPTS = 5               // Default number of attempts to connect to db
	INTERVAL     = 3 * time.Second // Time between attempts
)

var log logger.Logging
var models []func(db *gorm.DB)
var db *gorm.DB

type Adapter struct {
	IsConnected bool
	Logger      logger.Logging
	MaxAttempts int
	Interval    time.Duration
	db          *gorm.DB
}

func NewAdapter() *Adapter {
	// Add generic logger
	log = logger.NewLogger("adapter")

	// Setup Adapter
	a := Adapter{
		IsConnected: false,
		Logger:      log,
		MaxAttempts: MAX_ATTEMPTS,
		Interval:    INTERVAL,
	}

	return &a
}

func RegisterMigration(f func(db *gorm.DB)) {
	models = append(models, f)
}

func (a *Adapter) MigrationsMigrate() {
	a.Logger.Printf("Running db migrations")

	// Check if connected to db
	if a.IsConnected {
		a.Logger.Printf("Model connected to db and ready to run migrations")
		a.Logger.Printf("Got %d migrations registered", len(models))
		for _, model := range models {
			log.Printf("Migrating model %v", model)
			model(a.db)
		}
	}
}

func GetDb() *gorm.DB {
	return db
}

func (a *Adapter) Connect() (*gorm.DB, error) {
	isConnected := false
	var err error

	for i := 1; i <= a.MaxAttempts; i++ {
		db, err = isDbConnectionAvailable()

		if err == nil {
			isConnected = true
			log.Printf("Attempt %d: Connection to db established successfully", i)
			break
		}

		if err != nil && i < a.MaxAttempts {
			log.Fatalf("Attempt %d: Failed to established connection to db.", i)

			if i < (a.MaxAttempts) {
				log.Fatalf("Will retry connecting in %d seconds...", a.Interval/time.Second)
			}
		}

		time.Sleep(a.Interval)
	}

	// Add connection to adapter
	if isConnected {
		a.IsConnected = true
		a.db = db
	}

	return db, err
}

func isDbConnectionAvailable() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(getDbDsn()), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to db: %v", err)

		return db, err
	}

	return db, nil
}

func getDbDsn() string {
	dsn := ""

	// Get MySQL envs
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		user = "graph-ql-api-user"
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "graph-ql-api-password"
	}

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "8424"
	}

	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		database = "graph-ql-api"
	}

	return dsn + user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

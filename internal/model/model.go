package model

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

type Model struct {
	IsConnected bool
	Logger      logger.Logging
	MaxAttempts int
	Interval    time.Duration
	db          *gorm.DB
}

func NewModel() *Model {
	// Add generic logger
	log = logger.NewLogger("model")

	// Setup Model
	m := Model{
		IsConnected: false,
		Logger:      log,
		MaxAttempts: MAX_ATTEMPTS,
		Interval:    INTERVAL,
	}

	return &m
}

func RegisterMigration(f func(db *gorm.DB)) {
	models = append(models, f)
}

func (m *Model) MigrationsMigrate() {
	m.Logger.Printf("Running db migrations")

	// Check if connected to db
	if m.IsConnected {
		m.Logger.Printf("Model connected to db and ready to run migrations")
		m.Logger.Printf("Got %d migrations registered", len(models))
		for _, model := range models {
			log.Printf("Migrating model %v", model)
			model(m.db)
		}
	}
}

func GetDb() *gorm.DB {
	return db
}

func (m *Model) Connect() (*gorm.DB, error) {
	isConnected := false
	var err error

	for i := 1; i <= m.MaxAttempts; i++ {
		db, err = isDbConnectionAvailable()

		if err == nil {
			isConnected = true
			log.Printf("Attempt %d: Connection to db established successfully", i)
			break
		}

		if err != nil && i < m.MaxAttempts {
			log.Fatalf("Attempt %d: Failed to established connection to db.", i)

			if i < (m.MaxAttempts) {
				log.Fatalf("Will retry connecting in %d seconds...", m.Interval/time.Second)
			}
		}

		time.Sleep(m.Interval)
	}

	// Add connection to model
	if isConnected {
		m.IsConnected = true
		m.db = db
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
		user = "root"
	}

	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		password = "root"
	}

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("MYSQL_PORT")
	if port == "" {
		port = "3306"
	}

	database := os.Getenv("MYSQL_DATABASE")
	if database == "" {
		database = "app"
	}

	return dsn + user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

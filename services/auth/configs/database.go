package configs

import (
	"fmt"
	"os"

	"minisapi/services/auth/internal/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database holds the database connection
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database connection
func NewDatabase() (*Database, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &Database{db}, nil
}

// AutoMigrate runs database migrations
func (db *Database) AutoMigrate() error {
	return db.DB.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
	)
}

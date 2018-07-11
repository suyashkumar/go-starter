package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DefaultMaxIdleConns = 5

var ErrorNoConnectionString = errors.New("A connection string must be specified on the first call to Get")

// Handler abstracts away common DB operations needed for this project.
// Functions to fetch / insert db entities relevant to this project should
// be added to this interface (e.g. GetUser, InsertUser, etc)
type Handler interface {
	// GetDB returns the Handler's underlying *gorm.DB
	GetDB() *gorm.DB
}

type handler struct {
	db *gorm.DB
}

// NewHandler initializes and returns a new Handler
func NewHandler(dbConnection string) (Handler, error) {
	db, err := getDB(dbConnection)
	if err != nil {
		return nil, err
	}
	// AutoMigrate relevant schemas
	// db.AutoMigrate(&entities.User{})
	return &handler{
		db: db,
	}, nil
}

func (d *handler) GetDB() *gorm.DB {
	return d.db
}

func getDB(dbConnection string) (*gorm.DB, error) {
	if dbConnection == "" {
		return nil, ErrorNoConnectionString
	}

	d, err := gorm.Open("postgres", dbConnection)
	if err != nil {
		return nil, err
	}

	d.DB().SetMaxIdleConns(DefaultMaxIdleConns)

	return d, nil

}

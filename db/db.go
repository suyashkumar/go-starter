package db

import (
	"github.com/jinzhu/gorm"
	"github.com/suyashkumar/go-starter/config"
	"github.com/sirupsen/logrus"
)

const DefaultMaxIdleConns = 5

var db *gorm.DB

func Get() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}

	c := config.Get(config.DBConnString)
	d, err := gorm.Open("postgres", c)

	if err != nil {
		logrus.WithField("DBConnString", c).Error("Unable to connect to database")
	}

	d.DB().SetMaxIdleConns(DefaultMaxIdleConns)

	db = d

	return db, nil

}
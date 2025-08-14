package gormdb

import (
	"time"

	"github.com/umefy/go-web-app-template/pkg/validation"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB = gorm.DB

var ErrRecordNotFound = gorm.ErrRecordNotFound

type DBConfig struct {
	GormConfig      gorm.Config
	DBString        string
	EnableLog       bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func (c DBConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DBString, validation.Required),
	)
}

func NewDB(config DBConfig) (*DB, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var opts []gorm.Option
	opts = append(opts, &config.GormConfig)

	db, err := gorm.Open(postgres.Open(config.DBString), opts...)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}

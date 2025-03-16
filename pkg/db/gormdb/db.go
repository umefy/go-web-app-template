package gormdb

import (
	"log"
	"os"
	"time"

	"github.com/umefy/go-web-app-template/pkg/validation"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB = gorm.DB

var ErrRecordNotFound = gorm.ErrRecordNotFound

type DBConfig struct {
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
	if config.EnableLog {
		opts = append(opts, &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					ParameterizedQueries:      true,        // Don't include params in the SQL log
					Colorful:                  true,        // Disable color
				},
			),
		})
	}

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

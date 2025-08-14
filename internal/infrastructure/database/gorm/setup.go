package gorm

import (
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/umefy/go-web-app-template/internal/core/config"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"

	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	gormTracing "gorm.io/plugin/opentelemetry/tracing"
)

func NewDB(config config.Config) (*db.DB, error) {
	dbConfig := config.GetDBConfig()
	dbLoggerConfig := dbConfig.GetLoggerConfig()

	logger := gormLogger.New(
		log.New(getGormLoggerWriter(dbLoggerConfig.Writer), "\r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second * time.Duration(dbLoggerConfig.SlowThresholdInSeconds),
			LogLevel:                  getGormLoggerLevel(dbLoggerConfig.Level),
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      !dbLoggerConfig.ShowSqlParams, // show param as $ placeholder instead of actual params in the SQL log
			Colorful:                  true,
		},
	)

	gormConfig := gorm.Config{
		Logger: logger,
	}

	db, err := db.NewDB(db.DBConfig{
		DBString:        dbConfig.Url,
		MaxIdleConns:    dbConfig.MaxIdleConns,
		MaxOpenConns:    dbConfig.MaxOpenConns,
		ConnMaxLifetime: dbConfig.ConnMaxLifetime,
		GormConfig:      gormConfig,
	})

	if err != nil {
		return nil, err
	}

	// opentelemetry gorm plugin
	if err := db.Use(gormTracing.NewPlugin()); err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBQuery(db *db.DB) *query.Query {
	return query.Use(db)
}

func getGormLoggerLevel(level string) gormLogger.LogLevel {
	switch strings.ToLower(level) {
	case "info":
		return gormLogger.Info
	case "warn":
		return gormLogger.Warn
	case "error":
		return gormLogger.Error
	default:
		return gormLogger.Silent
	}
}

func getGormLoggerWriter(writerConfig string) io.Writer {
	switch strings.ToLower(writerConfig) {
	case "stdout":
		return os.Stdout
	default:
		return os.Stdout
	}
}

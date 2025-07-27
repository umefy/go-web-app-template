package gorm

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/core/config"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

func NewDB(config config.Config) (*db.DB, error) {
	dbConfig := config.GetDBConfig()

	db, err := db.NewDB(db.DBConfig{
		DBString:        dbConfig.Url,
		EnableLog:       dbConfig.EnableLog,
		MaxIdleConns:    dbConfig.MaxIdleConns,
		MaxOpenConns:    dbConfig.MaxOpenConns,
		ConnMaxLifetime: dbConfig.ConnMaxLifetime,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewDBQuery(db *db.DB) *query.Query {
	return query.Use(db)
}

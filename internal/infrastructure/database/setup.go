package database

import (
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	configSvc "github.com/umefy/go-web-app-template/internal/domain/config/service"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

func NewDB(configSvc configSvc.Service) (*db.DB, error) {
	dbConfig := configSvc.GetDBConfig()

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

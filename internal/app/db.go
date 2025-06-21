package app

import (
	configSvc "github.com/umefy/go-web-app-template/app/config/service"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

func newDB(configSvc configSvc.Service) (*db.DB, error) {
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

func newDBQuery(db *db.DB) *query.Query {
	return query.Use(db)
}

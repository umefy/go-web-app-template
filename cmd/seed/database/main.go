package main

import (
	"log"
	"os"

	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

func main() {
	query, err := getDbQuery()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	users, err := seedUsers(query)
	if err != nil {
		log.Fatalf("failed to seed user: %v", err)
	}

	_, err = seedOrders(query, users)
	if err != nil {
		log.Fatalf("failed to seed order: %v", err)
	}

	log.Println("seeded successfully")
}

func getDbQuery() (*query.Query, error) {
	db, err := db.NewDB(db.DBConfig{
		DBString: os.Getenv("DATABASE_URL"),
	})
	if err != nil {
		return nil, err
	}

	query := query.Use(db)
	return query, nil
}

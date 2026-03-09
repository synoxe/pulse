package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDatabase(cfg *Config) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sqlx.Connect("postrgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	log.Println("Veritabanına bağlanıldı")
	return db
}

package gorm_config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/radieske/manicure-scheduler-poc/pkg/config"
)

func ConnectDB() *gorm.DB {
	config.LoadEnv()

	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	user := config.GetEnv("DB_USER", "scheduler")
	password := config.GetEnv("DB_PASS", "scheduler123")
	dbname := config.GetEnv("DB_NAME", "schedulerdb")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	log.Println("[GORM] DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Falha ao conectar no banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Erro ao obter instância do DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conexão com banco de dados estabelecida com sucesso.")
	return db
}

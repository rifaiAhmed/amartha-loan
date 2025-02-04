package helpers

import (
	"amartha-loan/internal/models"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupPostgreSQL() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "5432"), GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_NAME", ""))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	logrus.Info("Successfully connect to database..")

	DB.AutoMigrate(&models.Loan{}, &models.Payoff{})
}

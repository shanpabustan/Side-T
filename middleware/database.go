// middleware/database.go
package middleware

import (
	"fmt"
	 "intern_template_v1/model"// <-- Import your model package

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	DBErr  error
)

func ConnectDB() bool {
	dns := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s TimeZone=%s",
		GetEnv("DB_HOST"), GetEnv("DB_PORT"), GetEnv("DB_NAME"),
		GetEnv("DB_UNME"), GetEnv("DB_PWRD"), GetEnv("DB_SSLM"),
		GetEnv("DB_TMEZ"))

	DBConn, DBErr = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if DBErr != nil {
		return true
	}

	// 🟢 Auto Migration of models here
	err := DBConn.AutoMigrate(
		   &model.Items{},
		   &model.Purchases{},
		// &model.Account{},
		// &model.ApprovalStatus{},
		// &model.Document{},
		// &model.Establishment{},
		// &model.Admin{},
		// &model.Announcement{},
	)
	if err != nil {
		fmt.Println("AutoMigrate error:", err)
		return true
	}

	return false
}

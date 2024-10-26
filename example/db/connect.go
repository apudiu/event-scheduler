package db

import (
	"fmt"
	"github.com/apudiu/event-scheduler/db/driver/gormdriver"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB database variable
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error

	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local // Postgres Database
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4 // MySQL/ MariaDB Database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		"root",
		"123456",
		"localhost",
		3306,
		"testdb",
	)
	// Connect to the DB and initialize the DB variable
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true, // required for gor.Err* compare
		Logger:         logger.Default.LogMode(logger.Silent),
		//Logger:         logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	// auto migrate
	mgErr := DB.AutoMigrate(
		&gormdriver.Model{},
	)
	if mgErr != nil {
		fmt.Println("Auto migration err", mgErr.Error())
		return
	} else {
		fmt.Println("Database Migrated")
	}
}
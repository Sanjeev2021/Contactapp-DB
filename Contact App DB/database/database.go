package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var database *gorm.DB

// InitDB initializes the database connection and performs necessary migrations
func InitDB() {
	// Database connection string , dsn = data source name
	dsn := "root:root@tcp(127.0.0.1:3306)/contactappangular?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("connection established")

	sqlDB, err := database.DB()
	if err != nil {
		panic("failed to access underlying DB object: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database: " + err.Error())
	}

	fmt.Println("MySQL connection established!")
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	if database == nil {
		InitDB()
	}

	return database
}

func Migrate(db *gorm.DB, data interface{}) error {
	err := db.AutoMigrate(data)
	if err != nil {
		return err
	}
	fmt.Println("Migration successful")
	return nil
}

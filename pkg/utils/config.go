package utils

import (
	"go-blog/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	// user     string = os.Getenv("DBUSER")
	// pass     string = os.Getenv("DBPASS")
	// host     string = os.Getenv("DBHOST")
	// port     string = os.Getenv("DBPORT")
	// database string = os.Getenv("DBNAME")
)

func Connect() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dsn := "root:191491@tcp(127.0.0.1:3306)/gonews?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
	// CreateDatabase()
}

func GetDB() *gorm.DB {
	if db == nil {
		Connect()
		Migrate()
	}
	return db
}

func CreateDatabase() {
	db.Migrator().DropTable(&models.User{}, &models.Blog{})
	db.Migrator().CreateTable(&models.User{})
	db.Migrator().CreateTable(&models.Blog{})
}

func Migrate() {
	db.Migrator().AutoMigrate(&models.User{})
	db.Migrator().AutoMigrate(&models.Blog{})
}

package connection

import (
	"go-blog/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:191491@tcp(127.0.0.1:3306)/gonews?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	// CreateDatabase()
	Migrate()
	if db == nil {
		Connect()
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

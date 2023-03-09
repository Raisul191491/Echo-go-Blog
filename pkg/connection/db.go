package connection

import (
	"fmt"
	"go-blog/pkg/models"
	"go-blog/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func DBConnect() {
	utils.SetConfig()
	config := utils.LocalConfig

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBIP, config.DbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	// CreateDatabase()
	if db == nil {
		DBConnect()
	}
	Migrate()
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

package connection

import (
	"fmt"
	"go-blog/pkg/models"
	"go-blog/pkg/utils"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	testDB *gorm.DB
)

func testDBConnect() {
	utils.SetConfig()
	config := utils.LocalConfig

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.TestDBUser, config.TestDBPass, config.TestDBIP, config.TestDbName)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	testDB = d
}

func RefreshDatabase() error {

	testDBConnect()

	if err := db.Migrator().DropTable(&models.User{}, &models.Blog{}); err != nil {
		return err
	}

	if err := db.Migrator().CreateTable(&models.User{}, &models.Blog{}); err != nil {
		return err
	}

	log.Printf("Successfully refreshed table")
	models.SetTestDBInstance(testDB)

	return nil
}

func GetTestDB() *gorm.DB {
	if err := RefreshDatabase(); err != nil {
		panic(err)
	}
	return testDB
}

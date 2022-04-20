package postgres

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	profile "github.com/lvornholt/go-profiles"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hello-fresh-menu-planning-service/internal/infra/repository/entity"
	"log"
	"sync"
)

var once sync.Once

var DbConnection *gorm.DB

func InitPostgresDB() {
	once.Do(func() {
		user := profile.GetStringValue("database.postgres.username")
		password := profile.GetStringValue("database.postgres.password")
		host := profile.GetStringValue("database.postgres.hostname")
		port := profile.GetStringValue("database.postgres.port")
		database := profile.GetStringValue("database.postgres.dbname")
		dbInfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			user,
			password,
			host,
			port,
			database,
		)
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dbInfo,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database")
			panic(err)
		}
		errMigrate := db.AutoMigrate(&entity.Recipe{}, &entity.Menu{}, &entity.Ingredient{}, &entity.User{}, &entity.Comment{})
		if errMigrate != nil {
			log.Println("Failed to migrate database")
			panic(errMigrate)
		}
		DbConnection = db
	})
}

func GetDB() *gorm.DB {
	return DbConnection
}

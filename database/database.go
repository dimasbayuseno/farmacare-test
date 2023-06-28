package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dimasbayuseno/farmacare-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.FightResult{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.BattleInfo{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.FoughtPokemon{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	var count int64
	db.Model(&models.User{}).Where("username = ?", "Operational of Team Rocket").Count(&count)
	if count == 0 {
		createOperational(db)
	}
	db.Model(&models.User{}).Where("username = ?", "Boss Giovanni").Count(&count)
	if count == 0 {
		createBoss(db)
	}
	db.Model(&models.User{}).Where("username = ?", "Black Market Merchant").Count(&count)
	if count == 0 {
		createBlackMarketMerchant(db)
	}

	return db, nil
}
func createBoss(db *gorm.DB) {
	user := &models.User{
		Username: "Boss Giovanni",
		Password: "password",
		Role:     "boss",
	}

	err := db.Create(user).Error
	if err != nil {
		log.Fatal(err)
	}
}

func createOperational(db *gorm.DB) {
	user := &models.User{
		Username: "Operational of Team Rocket",
		Password: "password",
		Role:     "operational",
	}

	err := db.Create(user).Error
	if err != nil {
		log.Fatal(err)
	}
}

func createBlackMarketMerchant(db *gorm.DB) {
	user := &models.User{
		Username: "Black Market Merchant",
		Password: "password",
		Role:     "merchant",
	}

	err := db.Create(user).Error
	if err != nil {
		log.Fatal(err)
	}
}

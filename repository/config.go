package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		//"password=%s dbname=%s sslmode=require",
		"password=%s dbname=%s",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection is succesful")

	/*err = DB.AutoMigrate(
		&model.Color{},
		&model.Apps{},
		&model.RelTeamApp{},
		&model.RelUserApp{},
		&model.Corporation{},
		&model.Office{},
		&model.Floor{},
		&model.Role{},
		&model.User{},
		&model.Visitor{},
		&model.UserAlertsHistory{},
		&model.Card{},
		&model.Company{},
		&model.VisitHistory{},
		&model.RecognitionAttemp{},
		&model.ColorHistory{},
		&model.Otp{},
		&model.Logs{})
	if err != nil {
		panic(err)
	}*/
	fmt.Println("Migration successful")

}

package db

import (
	"fmt"
	mod "gym-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	fmt.Println("Connecting to database...")
	dsn := "host=localhost port=5432 user=postgres password=Onkar17082001 dbname=gym_1 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error in connecting to database:", err)
	}
	db.AutoMigrate(&mod.Subscription{}, &mod.Payment{}, &mod.SubsType{}, &mod.User{}, &mod.GymEmp{}, &mod.Equipment{} , &mod.EmpTypes{}) 

	DB = db
	fmt.Println("Successfully connected to database")
}

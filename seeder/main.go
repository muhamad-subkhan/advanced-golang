package main

import (
	"fmt"
	"simple-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func main() {
	var err error

	// DB, err := gorm.Open("mysql", ""user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	DB, err = gorm.Open("postgres", "postgres://postgres:root@localhost/siswa?sslmode=disable")

	if err != nil {
		panic("failed to connect database")
	}

	defer DB.Close()

	fmt.Println("Successfully Connected to database")
	migrate()
}

func migrate(){
	DB.AutoMigrate(&models.Student{})

	data := models.Student{}
	if DB.Find(&data).RecordNotFound(){
		fmt.Println("data not found, Seeder will be running")
		seeder()
	}

}

func seeder(){
	data := models.Student{
		Student_id: 1,
		Student_name: "Jhon Doe",
		Student_age: 17,
		Student_address: "Jakarta",
		Student_phone: "0892656423",
	} 

	DB.Create(&data)
}
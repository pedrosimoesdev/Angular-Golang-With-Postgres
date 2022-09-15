package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Cars struct {
	Id        string
	Car       string
	Model     string
	Year      string
	DeleteAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func dbConn() (db *gorm.DB) {

	dbURL := "postgres://root:root@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Cars{})

	return db
}

func main() {

	server := gin.Default()

	server.Use(cors.Default())

	server.GET("/", getRecords)
	server.POST("/insert", insertRecords)
	server.PUT("/update", updateRecords)
	server.DELETE("/delete", deleteRecords)

	server.Run(":3333")

}

func getRecords(c *gin.Context) {

	db := dbConn()

	var cars []Cars
	db.Find(&cars)

	c.JSON(200, &cars)

}

func insertRecords(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	fmt.Println(string(data))
	//check why i need doing it
	//I guess that it use to convert to string this endpoint returns Json format
	records := Cars{}
	json.Unmarshal([]byte(data), &records)

	fmt.Println(records.Car)
	fmt.Println(records.Model)
	fmt.Println(records.Year)

	var Car = records.Car
	var Model = records.Model
	var Year = records.Year

	db := dbConn()

	car := Cars{Car: Car, Model: Model, Year: Year, CreatedAt: time.Now()}
	db.Select("car", "model", "year", "created_at").Create(&car)

	c.JSON(200, "Created records")

}

func updateRecords(c *gin.Context) {

	log.Println("test")

	data, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	//check why i need doing it
	//I guess that it use to convert to string this endpoint returns Json format
	records := Cars{}
	json.Unmarshal([]byte(data), &records)

	var id = records.Id
	var car = records.Car
	var model = records.Model
	var year = records.Year

	log.Println(id)
	log.Println(car)
	log.Println(year)
	log.Println(model)

	db := dbConn()

	db.Model(&Cars{}).Where("id = ?", id).Updates(Cars{Car: car, Model: model, Year: year})

	c.JSON(200, "Updated records")

}

func deleteRecords(c *gin.Context) {

	data, err := ioutil.ReadAll(c.Request.Body)

	var id = string(data)
	log.Println("ID:", id)

	if err != nil {
		log.Panicf("error: %s", err)
		c.JSON(500, "not ok records")
	}

	db := dbConn()

	db.Delete(&Cars{}, id)

	c.JSON(200, "Deleted records")

}

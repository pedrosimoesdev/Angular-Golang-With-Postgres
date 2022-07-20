package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Cars struct {
	id    int
	car   string
	model string
	year  int
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "crud"
)

func dbConn() (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	server := gin.Default()

	server.GET("/", getRecords)
	server.POST("/insert", insertRecords)
	server.PUT("/update", updateRecords)
	server.DELETE("/delete", deleteRecords)

	server.Run(":3333")

}

func getRecords(c *gin.Context) {

	db := dbConn()

	// We assign the result to 'rows'
	records, err := db.Query("SELECT * FROM cars")
	defer db.Close()

	// creates placeholder of the sandbox
	arrayCars := []Cars{}

	// we loop through the values of rows
	for records.Next() {
		car := Cars{}
		err := records.Scan(&car.id, &car.car, &car.model, &car.year)
		if err != nil {
			panic(err)
		}
		arrayCars = append(arrayCars, car)

	}
	fmt.Println(arrayCars)
	arrayCarsJson, err := json.Marshal(arrayCars)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	c.String(200, string(arrayCarsJson))
	fmt.Println(arrayCars)

}

func insertRecords(c *gin.Context) {

	db := dbConn()

	Statement := `
		INSERT INTO cars (car, model, year)
		VALUES ($1, $2, $3)`
	_, err := db.Exec(Statement, "BWM", "530", 2022)

	if err != nil {
		fmt.Println("error")
	} else {
		c.String(200, "records saved")
		fmt.Println("records saved")
	}
	defer db.Close()

}

func updateRecords(c *gin.Context) {

	db := dbConn()

	Statement := `
		update cars set car = $1 , model = $2 , year = $3 
		where id = $4`
	_, err := db.Exec(Statement, "BWMUpdated", "530Updates", 2045, 1)

	if err != nil {
		fmt.Println("error")
	} else {
		c.String(200, "records updated")
		fmt.Println("records updated")
	}
	defer db.Close()

}

func deleteRecords(c *gin.Context) {

	db := dbConn()

	Statement := `
		delete from cars where id = $1 `
	_, err := db.Exec(Statement, 1)

	if err != nil {
		fmt.Println("error")
	} else {
		c.String(200, "records deleted")
		fmt.Println("records deleted")
	}
	defer db.Close()

}

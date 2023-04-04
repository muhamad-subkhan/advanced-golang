package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	// "reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

type newStudent struct {
	Student_id      uint64 `json:"student_id" `
	Student_name    string `json:"student_name" `
	Student_age     uint64 `json:"student_age" `
	Student_address string `json:"student_address"`
	Student_phone   string `json:"student_phone"`
}

// func rowToStruct(rows *sql.Rows, dest interface{}) error{
// 	destv := reflect.ValueOf(dest).Elem()

// 	args := make([]interface{}, destv.Type().Elem().NumField())

// 	for rows.Next() {
// 		rowp := reflect.New(destv.Type().Elem())
// 		rowv := rowp.Elem()

// 		for i := 0; i < rowv.NumField(); i++ {
// 			args[i] = rowv.Field(i).Addr().Interface()
// 		}

// 		if err := rows.Scan(args...); err != nil {
// 			return err
// 		}

// 		destv.Set(reflect.Append(destv, rowv))
// 	}

// 	return nil
// }

func postHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent
	
	c.Bind(&newStudent)
	db.Create(&newStudent)
	c.JSON(http.StatusOK, gin.H{"message": "success create", "data": newStudent})

}

func getAll(c *gin.Context, db *gorm.DB) {
	var newStudent []newStudent

	db.Find(&newStudent)
	c.JSON(http.StatusOK, gin.H{"message": "success find all", "data": newStudent})
}

func getId(c *gin.Context, db *gorm.DB) {

	var newStudent []newStudent

	studentId := c.Param("student_id")


	if db.Find(&newStudent, "student_id=?", studentId).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success find by id", "data": newStudent})
}

func putHandler(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent

	studentId := c.Param("student_id")

	if db.Find(&newStudent, "student_id=?", studentId).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	var reqStudent = newStudent

	c.Bind(&reqStudent)

	db.Model(&newStudent).Where("student_id=?", studentId).Update(reqStudent)

	c.JSON(http.StatusOK, gin.H{
		"message": "success update",
		"data":    reqStudent,
	})

}


func delete(c *gin.Context, db *gorm.DB) {
	var newStudent newStudent


	studentId := c.Param("student_id")

	db.Delete(&newStudent, "student_id=?", studentId)

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete",
	})
}

func seederUser(db *gorm.DB) {
	data := newStudent{
		Student_id:       2,
		Student_name:     "Dono",
		Student_age:      20,
		Student_address:  "Jakarta",
		Student_phone: "0123456789",
	}

	db.Create(&data)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&newStudent{})

	data := newStudent{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("=================== run seeder user ======================")
		seederUser(db)
	}
}

func setRouter() *gin.Engine {
	connection := "postgres://postgres:root@localhost/coba?sslmode=disable"
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	
	Migrate(db)

	r := gin.Default()

	r.POST("/student", func(ctx *gin.Context) {
		postHandler(ctx, db)
	})


	r.GET("/student", func(ctx *gin.Context) {
		getAll(ctx, db)
	})

	r.GET("/student/:student_id", func(ctx *gin.Context) {
		getId(ctx, db)
	})

	r.PUT("/student/:student_id", func(ctx *gin.Context) {
		putHandler(ctx, db)
	})

	r.DELETE("/student/:student_id", func(ctx *gin.Context) {
		delete(ctx, db)
	})


	return r
}

func main() {

	r := setRouter()

	r.Run(":8080")

}
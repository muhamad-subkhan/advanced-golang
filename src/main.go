package main

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"

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

func rowToStruct(rows *sql.Rows, dest interface{}) error{
	destv := reflect.ValueOf(dest).Elem()

	args := make([]interface{}, destv.Type().Elem().NumField())

	for rows.Next() {
		rowp := reflect.New(destv.Type().Elem())
		rowv := rowp.Elem()

		for i := 0; i < rowv.NumField(); i++ {
			args[i] = rowv.Field(i).Addr().Interface()
		}

		if err := rows.Scan(args...); err != nil {
			return err
		}

		destv.Set(reflect.Append(destv, rowv))
	}

	return nil
}

func postHandler(c *gin.Context, db *sql.DB) {
	var newStudent newStudent

	if data := c.Bind(&newStudent); data == nil{
		_, err := db.Exec("insert into students values($1, $2, $3, $4, $5)", newStudent.Student_id, newStudent.Student_name, newStudent.Student_age, newStudent.Student_address, newStudent.Student_phone)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Seuccesfully Created"})

	}

	// c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
}

func getAll(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	row, err := db.Query("select * from students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	rowToStruct(row, &newStudent)
	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newStudent})

}

func getId(c *gin.Context, db *sql.DB) {
	var newStudent []newStudent

	studentId := c.Param("student_id")

	row, err := db.Query("select * from students where student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowToStruct(row, &newStudent)

	if newStudent == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newStudent})
}

func putHandler(c *gin.Context, db *sql.DB) {
	var newStudent newStudent

	studentId := c.Param("student_id")

	if c.Bind(&newStudent) == nil {
		_, err := db.Exec("update students set student_name = $1 where student_id = $2", newStudent.Student_name, studentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Success Update"})
	}
}


func delete(c *gin.Context, db *sql.DB) {
	studentId := c.Param("student_id")

	_, err := db.Exec("delete from students where student_id = $1", studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Success Delete"})
}


func setRouter() *gin.Engine {
	connection := "postgres://postgres:root@localhost/coba?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	
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
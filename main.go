package main

import (
	"awesomeProject/database"
	"awesomeProject/handler"
	"awesomeProject/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("The is Server Running on localhost port 3000")
	// initialize the database
	dbUser, dbPassword, dbName := "mohammad", "", "mohammad"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	//unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}
	//migration
	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		return
	}
	// route goes here

	http.HandleFunc("/students", handler.GetAllStudents)
	http.HandleFunc("/student", handler.GetStudentById)
	http.HandleFunc("/student/add", handler.AddStudent)
	http.HandleFunc("/student/delete", handler.DeleteStudent)
	http.HandleFunc("/student/update-student", handler.UpdateStudent)
	http.HandleFunc("/student/update-student-field", handler.UpdateStudentField)
	// listen port
	err = http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

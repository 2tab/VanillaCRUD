package main

import (
	"awesomeProject/controller"
	"awesomeProject/database"
	"awesomeProject/models"
	"awesomeProject/repository"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("The is Server Running on localhost port 3000")
	// initialize the database
	dbUser, dbPassword, dbName := "postgres", "password", "students"
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
	studentController := controller.StudentController{
		StudentRepository: repository.StudentRepository{
			DB: db,
		},
	}
	http.HandleFunc("/students", studentController.GetAll)
	http.HandleFunc("/student", studentController.GetById)
	http.HandleFunc("/student/add", studentController.Create)
	http.HandleFunc("/student/delete", studentController.Delete)
	http.HandleFunc("/student/update", studentController.Put)
	http.HandleFunc("/student/update-field", studentController.Patch)
	// listen port
	err = http.ListenAndServe(":3000", nil)
	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

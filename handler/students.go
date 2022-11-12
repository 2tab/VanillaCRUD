package handler

import (
	"awesomeProject/db"
	"awesomeProject/models"
	"awesomeProject/utils"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func GetAllStudents(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No students are registered"
`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	var students []models.Student
	// range returns index and element
	for _, student := range db.StudentDB {
		students = append(students, student)
	}

	//parse data into json format
	studentJSON, err := json.Marshal(&students)
	if err != nil {

		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No students are registered"
`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	utils.ReturnJsonResponse(res, http.StatusNotFound, studentJSON)
}

func GetStudentById(res http.ResponseWriter, req *http.Request) {
	// check if id exists in request and if the method is GET
	if _, ok := req.URL.Query()["id"]; !ok || req.Method != "GET" {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No students are registered"
`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]
	student, ok := db.StudentDB[id]

	if !ok {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No students are registered"
`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	//parse data into json format
	studentJSON, err := json.Marshal(&student)
	if err != nil {

		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No students are registered"
`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	utils.ReturnJsonResponse(res, http.StatusNotFound, studentJSON)
}

func AddStudent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	var student models.Student
	data := req.Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	//parse the student data into json format
	err := json.NewDecoder(data).Decode(&student)

	if err != nil {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	db.StudentDB[student.ID] = student

	HandlerMessage := []byte(`{
		"success" : true,
		"message" : "Student was added"
}`)
	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}

func DeleteStudent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		HandlerMessage := []byte(`{
		"success": false,
		"message" : "No students around"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	if _, ok := req.URL.Query()["id"]; !ok {

		HandlerMessage := []byte(`{
		"success": false,
		"message" : "No students around"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	id := req.URL.Query()["id"][0]
	if _, ok := db.StudentDB[id]; !ok {

		HandlerMessage := []byte(`{
		"success": false,
		"message" : "No students around"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	delete(db.StudentDB, id)

	HandlerMessage := []byte(`{
		"success": true,
		"message" : "Student deleted""
}`)
	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}

func UpdateStudent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}

	var student models.Student
	data := req.Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	//parse the student data into json format
	err := json.NewDecoder(data).Decode(&student)

	if err != nil {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	db.StudentDB[student.ID] = student

	HandlerMessage := []byte(`{
		"success" : true,
		"message" : "Student was updated"
}`)
	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}
func UpdateStudentField(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PATCH" {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	type PatchRequest struct {
		ID    string `json:"id"`
		FIELD string `json:"field"`
		VALUE string `json:"value"`
	}
	var patchRequest PatchRequest
	data := req.Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	//parse the student data into json format
	err := json.NewDecoder(data).Decode(&patchRequest)

	if err != nil {
		HandlerMessage := []byte(`{
		"success" : false,
		"message" : "No db available"
}`)
		utils.ReturnJsonResponse(res, http.StatusNotFound, HandlerMessage)
		return
	}
	student := db.StudentDB[patchRequest.ID]
	switch strings.ToLower(patchRequest.FIELD) {
	case "firstname":
		student.FirstName = patchRequest.VALUE
	case "lastname":
		student.LastName = patchRequest.VALUE

	}
	db.StudentDB[student.ID] = student

	HandlerMessage := []byte(`{
		"success" : true,
		"message" : "Student was updated"
}`)
	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}

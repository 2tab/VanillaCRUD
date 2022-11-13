package controller

import (
	"awesomeProject/DTO"
	"awesomeProject/db"
	"awesomeProject/models"
	"awesomeProject/repository"
	"awesomeProject/utils"
	"encoding/json"
	"io"
	"net/http"
)

type StudentController struct {
	StudentRepository repository.StudentRepository
}

func (s StudentController) GetAll(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	students := s.StudentRepository.GetAll()
	//parse data into json format
	studentJSON, err := json.Marshal(&students)
	if err != nil {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	utils.ReturnJsonResponse(res, http.StatusNotFound, studentJSON)
}

func (s StudentController) GetById(res http.ResponseWriter, req *http.Request) {
	// check if id exists in request and if the method is GET
	if _, ok := req.URL.Query()["id"]; !ok || req.Method != "GET" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	id := req.URL.Query()["id"][0]
	student := s.StudentRepository.Get(id)
	if student == nil {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	//parse data into json format
	studentJSON, err := json.Marshal(&student)
	if err != nil {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	utils.ReturnJsonResponse(res, http.StatusNotFound, studentJSON)
}

func (s StudentController) Create(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	var newStudent models.Student
	data := req.Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	//parse the newStudent data from json format into Student object
	err := json.NewDecoder(data).Decode(&newStudent)

	if err != nil {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	s.StudentRepository.AddStudent(&newStudent)

	handlerMessage := utils.HandleMessage(true, "new student added")
	utils.ReturnJsonResponse(res, http.StatusCreated, handlerMessage)
}

func (s StudentController) Put(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
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
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	if student := s.StudentRepository.UpdateStudent(&student); student == nil {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
	}
	handlerMessage := utils.HandleMessage(true, "Student updated")
	utils.ReturnJsonResponse(res, http.StatusCreated, handlerMessage)
}

func (s StudentController) Patch(res http.ResponseWriter, req *http.Request) {
	if req.Method != "Patch" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	var patchRequest DTO.PatchRequest
	data := req.Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)

	err := json.NewDecoder(data).Decode(&patchRequest)

	if err != nil {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	if student := s.StudentRepository.UpdateStudentField(&patchRequest); student == nil {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
	}
	handlerMessage := utils.HandleMessage(true, "Student updated")
	utils.ReturnJsonResponse(res, http.StatusCreated, handlerMessage)
}
func (s StudentController) Delete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	if _, ok := req.URL.Query()["id"]; !ok {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	id := req.URL.Query()["id"][0]
	if _, ok := db.StudentDB[id]; !ok {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}
	if flag := s.StudentRepository.Delete(id); !flag {

		handlerMessage := utils.HandleMessage(false, "No students are registered")
		utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
		return
	}

	handlerMessage := utils.HandleMessage(true, "Student deleted")
	utils.ReturnJsonResponse(res, http.StatusNotFound, handlerMessage)
	return
}

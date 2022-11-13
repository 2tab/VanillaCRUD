package repository

import (
	"awesomeProject/DTO"
	"awesomeProject/models"
	"fmt"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func (s StudentRepository) GetAll() []models.Student {
	var students []models.Student
	_ = s.DB.Find(&students)
	return students
}

func (s StudentRepository) Get(id string) *models.Student {
	student := &models.Student{}
	if err := s.DB.First(&student, id); err.Error != nil {
		return nil
	}
	return student
}

func (s StudentRepository) GetStudentByEmail(email string) *models.Student {
	student := &models.Student{}
	if err := s.DB.Where("email = ?", email).First(&student); err.Error != nil {
		return nil
	}
	return student
}

func (s StudentRepository) AddStudent(student *models.Student) *models.Student {
	if dbc := s.DB.Save(&student); dbc.Error != nil {
		return nil
	}
	return student
}

func (s StudentRepository) UpdateStudent(student *models.Student) *models.Student {
	s.DB.Model(&student).Updates(models.Student{FirstName: student.FirstName, LastName: student.LastName})
	return student
}

func (s StudentRepository) UpdateStudentField(patchRequest *DTO.PatchRequest) *models.Student {
	err := s.DB.Model(&models.Student{}).Where("id = ?", patchRequest.ID).Update(patchRequest.FIELD, patchRequest.VALUE)
	if err.Error != nil {
		fmt.Println(err.Error)
	}
	student := &models.Student{}
	_ = s.DB.First(&student, patchRequest.ID)
	return student
}
func (s StudentRepository) Delete(id string) bool {
	if err := s.DB.Delete(&models.Student{}, id); err.Error != nil {
		return false
	}
	return true
}

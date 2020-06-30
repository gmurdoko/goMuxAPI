package usecases

import (
	"goMuxAPI/main/master/models"
	"goMuxAPI/main/master/repositories"
)

//StudentUsecaseImpl app
type StudentUsecaseImpl struct {
	studentRepo repositories.StudentRepository
}

//GetStudent app
func (s StudentUsecaseImpl) GetStudent() ([]*models.Students, error) {
	students, err := s.studentRepo.GetAllStudent()
	if err != nil {
		return nil, err
	}
	return students, nil
}

//InitStudentUsecaseImpl app
func InitStudentUsecaseImpl(studentRepo repositories.StudentRepository) StudentUsecase {
	return &StudentUsecaseImpl{studentRepo}
}

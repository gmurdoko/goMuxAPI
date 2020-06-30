package usecases

import "goMuxAPI/main/master/models"

//StudentUsecase app
type StudentUsecase interface {
	GetStudent() ([]*models.Students, error)
}

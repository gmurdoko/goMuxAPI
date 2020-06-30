package repositories

import "goMuxAPI/main/master/models"

//StudentRepository app
type StudentRepository interface {
	GetAllStudent() ([]*models.Students, error)
}

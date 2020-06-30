package repositories

import (
	"database/sql"
	"goMuxAPI/main/master/models"
)

//StudentRepoImpl app
type StudentRepoImpl struct {
	db *sql.DB
}

//GetAllStudent app
func (s StudentRepoImpl) GetAllStudent() ([]*models.Students, error) {
	query := "SELECT * FROM students"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Students{}
	for data.Next() {
		var student = models.Students{}
		var err = data.Scan(&student.ID, &student.FistName, &student.LastName, &student.Email)
		if err != nil {
			return nil, err
		}
		result = append(result, &student)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
	// panic("implement me")
}

//InitStudentRepoImpl app
func InitStudentRepoImpl(db *sql.DB) StudentRepository {
	return &StudentRepoImpl{db}
}

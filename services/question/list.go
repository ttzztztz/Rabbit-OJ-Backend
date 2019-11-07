package question

import (
	"Rabbit-OJ-Backend/db"
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/utils"
)

func List(page uint32) ([]models.Question, error) {
	var list []models.Question

	err := db.DB.Table("question").
		Order("tid asc").
		Limit(utils.PageSize).
		Offset((page - 1) * utils.PageSize).
		Scan(&list).Error

	if err != nil {
		return nil, err
	}
	return list, nil
}
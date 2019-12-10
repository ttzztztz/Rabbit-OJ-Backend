package user

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/forms"
	"Rabbit-OJ-Backend/services/db"
	"Rabbit-OJ-Backend/utils"
	"errors"
	"time"
)

func Register(form *forms.RegisterForm) (string, error) {
	if UsernameExist(form.Username) {
		return InvalidUid, errors.New("username already exists")
	}

	if EmailExist(form.Email) {
		return InvalidUid, errors.New("email already exists")
	}

	user := models.User{
		Username: form.Username,
		Password: utils.SaltPasswordWithSecret(form.Password),
		Email:    form.Email,
		LoginAt:  time.Now(),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return InvalidUid, err
	}

	return user.Uid, nil
}

package repositories

import (
	"usersApi/initializers"
	"usersApi/models"
)

func AddUser(User models.User) (uint, error) {
	result := initializers.DB.Create(&User)
	if result.Error != nil {
		return 0, result.Error
	}
	return User.ID, nil
}

func GetById(id int) (models.User, error) {
	var user models.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	user.Password = ""
	return user, nil
}

func GetByEmail(email string) (error, models.User) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error, user
	}
	return nil, user
}

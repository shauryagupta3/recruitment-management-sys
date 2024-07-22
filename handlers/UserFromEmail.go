package handlers

import "github.com/shauryagupta3/recruitment-management-sys/models"

func (h handler) GetUserFromEmail(email string) (models.User, error) {
	var user models.User
	if result := h.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

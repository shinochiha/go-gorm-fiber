package common

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NotFoundResponse(name string, id string) fiber.Map {
	return fiber.Map{
		"message": name + " with id " + id + " not found",
	}
}

func InternalServerError() fiber.Map {
	return fiber.Map{
		"message": "Internal Server Error",
	}
}

func IsExists(userId string) (model.User, error) {
	var user model.User
	err := config.DB.First(&user, "id = ?", userId).Error
	return user, err
}

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}

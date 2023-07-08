package handler

import (
	"go-fiber-gorm/common"
	"go-fiber-gorm/config"
	"go-fiber-gorm/model"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {
	p := new(model.UserRequest)

	if err := ctx.BodyParser(p); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(p)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	var user model.User
	errEmailUser := config.DB.First(&user, "email = ? ", p.Email).Error
	if errEmailUser != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "email belum terdaftar",
		})
	}
	errHashingPassword := common.ComparePasswords(p.Password, user.Password)
	if errHashingPassword != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := common.GenerateToken(&claims)
	if errGenerateToken != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errGenerateToken.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}

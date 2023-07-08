package handler

import (
	"go-fiber-gorm/common"
	"go-fiber-gorm/config"
	"go-fiber-gorm/model"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const endPoint = "users"

func GetUserHandler(ctx *fiber.Ctx) error {
	var users []model.User

	err := config.DB.Order("id").Find(&users).Error

	if err != nil {
		log.Println(err.Error())
	}

	return ctx.JSON(users)
}

func CreateUserHandler(ctx *fiber.Ctx) error {
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

	newUser := model.User{
		Name:    p.Name,
		Email:   p.Email,
		Phone:   p.Phone,
		Address: p.Address,
	}

	hashedPassword, errHashingPassword := common.HashingPassword(p.Password)
	if errHashingPassword != nil {
		return ctx.Status(500).JSON(common.InternalServerError())
	}

	newUser.Password = hashedPassword

	errCreateUser := config.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(common.InternalServerError())
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newUser,
	})
}

func GetByIdUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := common.IsExists(userId)
	if err != nil {
		return ctx.Status(404).JSON(common.NotFoundResponse(endPoint, userId))
	}
	return ctx.Status(404).JSON(fiber.Map{
		"message": "succes",
		"data":    user,
	})
}

func UpdateByIdUser(ctx *fiber.Ctx) error {
	p := new(model.User)

	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	userId := ctx.Params("id")
	// Check Available User
	user, err := common.IsExists(userId)
	if err != nil {
		return ctx.Status(404).JSON(common.NotFoundResponse(endPoint, userId))
	}

	// Update User
	if p.Name != "" {
		user.Name = p.Name
	}
	if p.Email != "" {
		user.Email = p.Email
	}
	if p.Phone != "" {
		user.Phone = p.Phone
	}
	if p.Address != "" {
		user.Address = p.Address
	}

	errUpdate := config.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(common.InternalServerError())
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

}

func DeleteByIdUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := common.IsExists(userId)
	if err != nil {
		return ctx.Status(404).JSON(common.NotFoundResponse(endPoint, userId))
	}

	errDelete := config.DB.Delete(&user).Error
	if errDelete != nil {

		return ctx.Status(500).JSON(common.InternalServerError())
	}

	return ctx.JSON(fiber.Map{
		"message": "User has deleted",
	})
}

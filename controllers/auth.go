package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itzcodex24/edu-swipe-api/database"
	"github.com/itzcodex24/edu-swipe-api/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

func GetHello(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello, World!",
		"status":  fiber.StatusOK,
	})
}

func Register(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		if err := ctx.SendStatus(fiber.StatusBadRequest); err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			if err := ctx.JSON(fiber.Map{
				"message": "Invalid request",
				"ok":      false,
			}); err != nil {
				panic("Couldn't send response" + err.Error())
			}
		}
		return ctx.JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "Couldn't hash password",
			"ok":      false,
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	if err := database.DB.Create(&user); err.Error != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "Couldn't create user",
			"ok":      false,
		})
	}
	return ctx.JSON(fiber.Map{
		"message":    "User created successfully",
		"ok":         true,
		"statusCode": fiber.StatusCreated,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid request",
			"ok":      false,
		})
	}

	var user models.User

	if err := database.DB.Where("email = ?", data["email"]).First(&user); err.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
			"ok":      false,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password or email provided are not correct",
			"ok":      false,
		})
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Couldn't log you in",
			"ok":      false,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Logged in",
		"ok":      true,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
			"ok":      false,
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["iss"] == nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
			"ok":      false,
		})
	}

	var user models.User
	if err := database.DB.Where("id = ?", claims["iss"]).First(&user); err.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
			"ok":      false,
		})
	}
	return c.SendString("Hello, " + user.Name + "!")
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

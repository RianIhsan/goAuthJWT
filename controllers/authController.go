package controllers

import (
	"strconv"
	"time"

	"github.com/RianIhsan/ApiGoJwt/database"
	"github.com/RianIhsan/ApiGoJwt/models"
	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.DB.Create(&user)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Berhasil Membuat akun",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Name:  data["name"],
		Email: data["email"],
	}
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Email Salah!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Password salah!",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Tidak bisa login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "Lax",
	}
	c.Cookie(&cookie)
  
  //Header Authorization dengan bearer token
  c.Set("Authorization", "Bearer "+token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    user,
		"message": "Berhasil Login!",
		"token":   cookie,
	})
}

func User(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Tidak Terotentikasi",
		})
	}

	// Menghapus "Bearer " dari token
	token = strings.Replace(token, "Bearer ", "", 1)

	claims := &jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Tidak Terotentikasi",
		})
	}

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "Lax",
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Berhasil Logout",
	})
}

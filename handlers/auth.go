package handlers

import (
	"echo-jwt/database"
	"echo-jwt/models"
	"echo-jwt/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func Register(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	if isUserRegistered(&models.User{}, user.Email) {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	authUser := models.User{}
	if !isUserRegistered(&authUser, user.Email) {
		return c.JSON(http.StatusNotFound, "user does not exist")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.JSON(http.StatusBadRequest, "passwords do not match")
		}

		return nil
	}

	token := utils.GenerateToken(authUser.ID)

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//cookie.Secure = true
	cookie.HttpOnly = true
	//cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	//return c.JSON(http.StatusOK, echo.Map{
	//	"token": token,
	//})

	return c.NoContent(http.StatusOK)
}

func isUserRegistered(user *models.User, email string) bool {
	result := database.DB.Where("email", email).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

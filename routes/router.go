package routes

import (
	"echo-jwt/database"
	"echo-jwt/handlers"
	"echo-jwt/models"
	"echo-jwt/utils"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("api/register", handlers.Register)
	e.POST("api/login", handlers.Login)

	r := e.Group("api/restricted")
	r.Use(echojwt.WithConfig(utils.JwtConfig))

	r.GET("/users", getUser)

	e.Logger.Fatal(e.Start(":1323"))

	return e
}

func getUser(c echo.Context) error {
	claims := utils.GetClaims(c)

	id := claims.ID

	user := models.User{}
	database.DB.First(&user, id)

	return c.JSON(http.StatusOK, user)
}

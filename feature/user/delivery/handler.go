package delivery

import (
	"log"
	"net/http"
	"socialmedia-app/config"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.POST("/users", handler.InsertUser())
	e.POST("/login", handler.LogUser())
	e.GET("/profile", handler.GetProfile(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
			"token":   common.GenerateToken(data.ID),
		})
	}
}

func (uh *userHandler) LogUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, "email or password incorrect")
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, "email or password incorrect")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"token":   common.GenerateToken(data.ID),
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}

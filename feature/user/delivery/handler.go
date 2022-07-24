package delivery

import (
	"log"
	"net/http"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.POST("/users", handler.InsertUser())
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
		var userLogin InsertFormat
		err := c.Bind(&userLogin)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "email or password incorrect")
		}
		data, e := uh.userUsecase.LoginUser(userLogin.ToModel())
		if e != nil {
			return c.JSON(http.StatusBadRequest, "email or password incorrect")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"data":    data,
			"token":   common.GenerateToken(data.ID),
		})
	}
}

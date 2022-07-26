package delivery

import (
	"log"
	"net/http"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type newsHandler struct {
	newsUsecase domain.NewsUseCase
}

func New(nu domain.NewsUseCase) domain.NewsHandler {
	return &newsHandler{
		newsUsecase: nu,
	}
}

func (nh *newsHandler) InsertNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp NewsInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := nh.newsUsecase.AddNews(common.ExtractData(c), tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})

	}
}

func (nh *newsHandler) GetAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.newsUsecase.GetAllN()

		if err != nil {
			log.Println("Cannot get data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all news",
			"users":   data,
		})
	}
}

func (bh *newsHandler) GetMyNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := bh.newsUsecase.GetMyN(common.ExtractData(c))

		if err != nil {
			log.Println("Cannot get data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my news",
			"users":   data,
		})
	}
}
func (nh *newsHandler) DeleteNews() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := nh.newsUsecase.DelNews(cnv)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete news",
		})
	}
}

package delivery

import (
	"log"
	"net/http"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

type commentsHandler struct {
	commentsUsecase domain.CommentsUseCase
}

func New(cu domain.CommentsUseCase) domain.CommentsHandler {
	return &commentsHandler{
		commentsUsecase: cu,
	}
}

func (cu *commentsHandler) InsertComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp commentsInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := cu.commentsUsecase.AddComments(common.ExtractData(c), tmp.ToDomain())

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

func (cu *commentsHandler) GetAllComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := cu.commentsUsecase.GetAllC()

		if err != nil {
			log.Println("Cannot get data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "Problem from database")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Comments",
			"users":   data,
		})
	}
}

func (cu *commentsHandler) GetMyComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := cu.commentsUsecase.GetMyC(common.ExtractData(c))

		if err != nil {
			log.Println("Cannot get data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my comments",
			"users":   data,
		})
	}
}
func (cu *commentsHandler) DeleteComments() echo.HandlerFunc {
	return func(c echo.Context) error {

		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data, err := cu.commentsUsecase.DelComments(cnv)

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
			"message": "success delete Comments",
		})
	}
}

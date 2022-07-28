package delivery

import (
	"fmt"
	"log"
	"net/http"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"
	"strconv"

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

		fmt.Println(tmp)

		var userid = common.ExtractData(c)
		data, err := nh.newsUsecase.AddNews(common.ExtractData(c), tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Println(userid)

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    FromDomain(data),
		})

	}
}

func (nh *newsHandler) UpdateNews() echo.HandlerFunc {
	return func(c echo.Context) error {

		qry := map[string]interface{}{}
		cnv, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp NewsInsertRequest
		res := c.Bind(&tmp)

		if res != nil {
			log.Println(res, "Cannot parse data")
			return c.JSON(http.StatusInternalServerError, "error read update")
		}

		if tmp.Content != "" {
			qry["content"] = tmp.Content
		}
		if tmp.Images != "" {
			qry["images"] = tmp.Images
		}
		if tmp.FileAttachment != "" {
			qry["file_attachment"] = tmp.FileAttachment
		}

		data, err := nh.newsUsecase.UpNews(cnv, tmp.ToDomain())

		if err != nil {
			log.Println("Cannot update data", err)
			c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update data",
			"id":      data.ID,
			"content": data.Content,
			"images":  data.Images,
			"file":    data.FileAttachment,
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
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}

		if !data {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete news",
		})
	}
}

func (nh *newsHandler) GetAllNews() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := nh.newsUsecase.GetAllN()

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "error read input")

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

func (nh *newsHandler) GetNewsID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idNews := c.Param("id")
		id, _ := strconv.Atoi(idNews)
		data, err := nh.newsUsecase.GetSpecificNews(id)

		if err != nil {
			log.Println("Cannot get data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get my news",
			"users":   data,
		})
	}
}

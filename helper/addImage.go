package helper

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

func AddImageEvent(c echo.Context) (link string, message map[string]interface{}, err error) {
	file, err := c.FormFile("file")
	bucket := os.Getenv("DB_BUCKET")
	if err != nil {
		url := "https://storage.googleapis.com/" + bucket + "/default_event.jpg"
		return url, map[string]interface{}{
			"message": "Success create url",
			"code":    200,
		}, nil
	} else {
		var storageClient *storage.Client
		ctx := appengine.NewContext(c.Request())
		storageClient, errStorage := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
		if errStorage != nil {
			return "", map[string]interface{}{
				"message": "misssing credentials file",
				"code":    500,
			}, errStorage
		}
		if file.Size > 1024*1024 {
			return "", map[string]interface{}{
				"message": "The uploaded image is too big. Please use an image less than 1MB in size",
				"code":    404,
			}, fmt.Errorf("size to big")
		}
		src, err := file.Open()
		if err != nil {
			return "", map[string]interface{}{
				"message": "Failed to get Open File",
				"code":    500,
			}, err
		}
		defer src.Close()
		if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
			if file.Filename[len(file.Filename)-4:] != "jpeg" {
				return "", map[string]interface{}{
					"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
					"code":    404,
				}, fmt.Errorf("file type not accepted")
			}
		}

		name := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
		sw := storageClient.Bucket(bucket).Object(name).NewWriter(ctx)

		if _, err := io.Copy(sw, src); err != nil {
			return "", map[string]interface{}{
				"message": err,
				"code":    500,
			}, err
		}
		if err := sw.Close(); err != nil {
			return "", map[string]interface{}{
				"message": err,
				"code":    500,
			}, err
		}
		u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
		if err != nil {
			return "", map[string]interface{}{
				"message": "Failed create url",
				"code":    500,
			}, err
		}
		return u.String(), map[string]interface{}{
			"message": "Success create url",
			"code":    200,
		}, nil
	}
}

func AddImageUser(c echo.Context) (link string, message map[string]interface{}, err error) {
	file, err := c.FormFile("file")
	bucket := os.Getenv("DB_BUCKET")
	if err != nil {
		url := "https://storage.googleapis.com/" + bucket + "/default_profile.png"
		return url, map[string]interface{}{
			"message": "Success create url",
			"code":    200,
		}, nil
	} else {
		var storageClient *storage.Client
		ctx := appengine.NewContext(c.Request())
		storageClient, errStorage := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
		if errStorage != nil {
			return "", map[string]interface{}{
				"message": "misssing credentials file",
				"code":    500,
			}, errStorage
		}
		if file.Size > 1024*1024 {
			return "", map[string]interface{}{
				"message": "The uploaded image is too big. Please use an image less than 1MB in size",
				"code":    404,
			}, fmt.Errorf("size to big")
		}
		src, err := file.Open()
		if err != nil {
			return "", map[string]interface{}{
				"message": "Failed to get Open File",
				"code":    500,
			}, err
		}
		defer src.Close()
		if file.Filename[len(file.Filename)-3:] != "jpg" && file.Filename[len(file.Filename)-3:] != "png" {
			if file.Filename[len(file.Filename)-4:] != "jpeg" {
				return "", map[string]interface{}{
					"message": "The provided file format is not allowed. Please upload a JPG or JPEG or PNG image",
					"code":    404,
				}, fmt.Errorf("file type not accepted")
			}
		}

		name := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
		sw := storageClient.Bucket(bucket).Object(name).NewWriter(ctx)

		if _, err := io.Copy(sw, src); err != nil {
			return "", map[string]interface{}{
				"message": fmt.Sprintf("Copy err: %s", err),
				"code":    500,
			}, err
		}
		if err := sw.Close(); err != nil {
			return "", map[string]interface{}{
				"message": fmt.Sprintf("Close err: %s", err),
				"code":    500,
			}, err
		}
		u, err := url.Parse("https://storage.googleapis.com/" + bucket + "/" + sw.Attrs().Name)
		if err != nil {
			return "", map[string]interface{}{
				"message": "Failed create url",
				"code":    500,
			}, err
		}
		return u.String(), map[string]interface{}{
			"message": "Success create url",
			"code":    200,
		}, nil
	}
}

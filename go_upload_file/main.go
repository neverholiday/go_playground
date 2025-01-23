package main

import (
	"fmt"
	"net/http"

	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

type CsvModel struct {
	Name string `csv:"name"`
	Age  string `csv:"age"`
}

type UploadResponse struct {
	DomainID   string `json:"domain_id"`
	CampaignID string `json:"campaign_id"`
	FileName   string `json:"filename"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func upload(c echo.Context) error {

	domainID := c.FormValue("domain")
	campaignID := c.FormValue("campaign")

	f, err := c.FormFile("csvfile")
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			Response{
				Message: err.Error(),
			},
		)
	}

	src, err := f.Open()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			Response{
				Message: err.Error(),
			},
		)
	}

	peopleList := []CsvModel{}
	err = gocsv.UnmarshalMultipartFile(&src, &peopleList)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			Response{
				Message: err.Error(),
			},
		)
	}

	for _, p := range peopleList {
		fmt.Println(p.Name)
		fmt.Println(p.Age)
	}

	return c.JSON(
		http.StatusOK,
		Response{
			Message: "upload success",
			Data: UploadResponse{
				DomainID:   domainID,
				FileName:   f.Filename,
				CampaignID: campaignID,
			},
		},
	)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.POST("/upload", upload)
	e.Logger.Fatal(e.Start(":1323"))
}

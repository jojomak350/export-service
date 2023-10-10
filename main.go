package main

import (
	"export-service/core"
	"export-service/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(os.Getenv("TEMP_DIRECTORY")); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("TEMP_DIRECTORY"), os.ModePerm)
	}

	core.LoadLogger()

	core.NewUploader()
}

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"msg": "welcome",
		})
	})

	app.POST("/export", handlers.ExportHandler)

	core.Logger.Fatal(app.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}

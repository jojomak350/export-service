package main

import (
	"export-service/core"
	"export-service/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	core.NewUploader()
}

func main() {
	e := echo.New()

	e.POST("/export", handlers.ExportHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("PORT"))))
}

package handlers

import (
	"export-service/core"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PostRequest struct {
	Body     [][]interface{} `json:"body"`
	Header   []string        `json:"header"`
	Filepath string          `json:"filepath"`
}

func ExportHandler(c echo.Context) error {
	var request PostRequest
	if err := c.Bind(&request); err != nil {
		core.Logger.Println(err)

		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	file := core.NewFile(request.Filepath)

	file.Generate(request.Header, request.Body)

	return c.JSON(http.StatusOK, echo.Map{
		"url": file.Save(core.UploaderClient),
	})
}

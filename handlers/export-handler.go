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
		panic(err)
	}

	//client := &http.Client{}
	//req, _ := http.NewRequest("GET", request.Endpoint, nil)
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", request.Token))
	//response, err := client.Do(req)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, echo.Map{
	//		"error": err,
	//	})
	//}

	//defer response.Body.Close()

	//if response.StatusCode != http.StatusOK {
	//	return c.JSON(http.StatusBadRequest, echo.Map{
	//		"status": response.StatusCode,
	//	})
	//}

	//data, err := io.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}

	//var commentResponse []map[string]interface{}

	//err = json.Unmarshal(data, &commentResponse)
	//if err != nil {
	//	panic(err)
	//}

	file := core.NewFile(request.Filepath)

	file.Generate(request.Header, request.Body)

	return c.JSON(http.StatusOK, echo.Map{
		"url": file.Save(core.UploaderClient),
	})
}

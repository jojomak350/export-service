package core

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

type Excel struct {
	file       *excelize.File
	path       string
	tempPath   string
	sheetName  string
	sheetIndex int
	header     []string
	//reflection reflect.Type
}

func NewFile(path string) Excel {
	e := Excel{
		file:      excelize.NewFile(),
		sheetName: "Sheet1",
		path:      path,
		tempPath:  filepath.Join(os.Getenv("TEMP_DIRECTORY"), filepath.Base(path)),
	}

	defer func() {
		if err := e.file.Close(); err != nil {
			Logger.Println(err)
		}
	}()

	return e
}

func (e *Excel) SetSheetName(name string) {
	e.sheetName = name
}

func (e *Excel) setHeader(data []string) {
	e.header = data
}

//func (e *Excel) setStructType(ref any) {
//	e.reflection = reflect.TypeOf(ref)
//}

func (e *Excel) setActiveSheet() {
	e.file.SetActiveSheet(e.sheetIndex)
}

func (e *Excel) createSheet() {
	index, err := e.file.NewSheet(e.sheetName)
	if err != nil {
		Logger.Println(err)
	}

	e.sheetIndex = index
}

func (e *Excel) appendCell(place string, value interface{}) {
	err := e.file.SetCellValue(e.sheetName, place, value)
	if err != nil {
		Logger.Println(err)
	}
}

func (e *Excel) createFileHeader() {
	firstChar := 65

	for i, el := range e.header {
		e.appendCell(fmt.Sprintf("%v%v", string(rune(firstChar+i)), 1), el)
	}
}

//func (e *Excel) createFileHeader() {
//	firstChar := 65
//
//	for i := 0; i < e.reflection.NumField(); i++ {
//		e.appendCell(fmt.Sprintf("%v%v", string(rune(firstChar+i)), 1), e.reflection.Field(i).Name)
//	}
//}

func (e *Excel) createFileBody(data [][]interface{}) {
	firstChar := 65

	for index, record := range data {
		for i, _ := range e.header {
			e.appendCell(fmt.Sprintf("%v%v", string(rune(firstChar+i)), index+2), record[i])
		}
	}
}

//func (e *Excel) createFileBody(data []map[string]interface{}) {
//	firstChar := 65
//
//	for index, record := range data {
//		for i := 0; i < e.reflection.NumField(); i++ {
//			cell := e.reflection.Field(i).Tag.Get("json")
//			e.appendCell(fmt.Sprintf("%v%v", string(rune(firstChar+i)), index+2), record[cell])
//		}
//	}
//}

func (e *Excel) loadFile() *os.File {
	f, err := os.Open(e.tempPath)
	if err != nil {
		Logger.Println(err)
	}

	return f
}

func (e *Excel) Save(client Uploader) string {
	err := e.file.SaveAs(e.tempPath)
	if err != nil {
		Logger.Println(err)
	}

	return client.Upload(e.loadFile(), e.path)
}

func (e *Excel) Generate(header []string, body [][]interface{}) {
	e.setHeader(header)

	e.createSheet()

	e.createFileHeader()

	e.createFileBody(body)

	e.setActiveSheet()
}

//func (e *Excel) Generate(data []map[string]interface{}, ref any) {
//	e.setStructType(ref)
//
//	e.createSheet()
//
//	e.createFileHeader()
//
//	e.createFileBody(data)
//
//	e.setActiveSheet()
//}

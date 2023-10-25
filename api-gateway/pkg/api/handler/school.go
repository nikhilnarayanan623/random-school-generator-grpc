package handler

import (
	"api-gateway/pkg/api/handler/interfaces"
	clientInterface "api-gateway/pkg/client/interfaces"
	"api-gateway/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type schoolHandler struct {
	client clientInterface.SchoolClient
}

func NewSchoolHandler(client clientInterface.SchoolClient) interfaces.SchoolHandler {
	return &schoolHandler{
		client: client,
	}
}

func (s *schoolHandler) GetOne(ctx *gin.Context) {

	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "name in query not found",
		})
		return
	}

	data, err := s.client.GetOneInJSON(name)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to get school data",
			"error":   err,
		})
		return
	}

	// if the request content type is in json then write the json data to response
	if ctx.ContentType() == "application/json" {
		ctx.Header("Content-Type", "application/json")
		ctx.Writer.Write(data)
		return
	}

	// if it's need in excel unmarshal to school
	var school models.School

	if err := json.Unmarshal(data, &school); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to unmarshal data to school",
			"error":   err,
		})
		return
	}

	// convert the school to excel
	file, err := convertToExcel(school)

	ctx.Header("Content-Type", "application/octet-stream")
	contentDisp := fmt.Sprintf("attachment; filename=%s.xlsx", school.Name)
	ctx.Header("Content-Disposition", contentDisp)

	if err := file.Write(ctx.Writer); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to write excel file to response",
			"error":   err,
		})
	}

}

func convertToExcel(school models.School) (*excelize.File, error) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheetName := "school"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to create sheet: %w", err)
	}

	headers := []string{
		"School", "Class",
		"Name", "Age", "RollNumber", "Gender", "HaveDisability",

		"House No", "City", "State",

		"Subject1", "Score", "Grade", "Class Category", "Passed",
		"Subject2", "Score", "Grade", "Class Category", "Passed",
		"Subject3", "Score", "Grade", "Class Category", "Passed",
		"Subject4", "Score", "Grade", "Class Category", "Passed",
		"Subject5", "Score", "Grade", "Class Category", "Passed",
		"Subject6", "Score", "Grade", "Class Category", "Passed",
	}

	cellNames := make([]string, len(headers)) // to store cell names
	for i, header := range headers {

		// save cell names
		cellName, err := excelize.ColumnNumberToName(i + 1)
		if err != nil {
			return nil, fmt.Errorf("failed to get column name")
		}

		cellNames[i] = cellName
		// set header values
		f.SetCellValue(sheetName, cellName+"1", header)
	}

	row := 2

	for _, class := range school.Classes {
		for _, student := range class.Students {

			// school and class
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[0], row), school.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[1], row), class.Name)

			// student detail
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[2], row), student.Name)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[3], row), student.Age)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[4], row), student.RollNumber)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[5], row), student.Gender)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[6], row), student.HaveDisability)

			// set addresses
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[7], row), student.Address.HouseNumber)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[8], row), student.Address.City)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[9], row), student.Address.State)
			// set score details
			for i, subject := range student.Scores {

				// each subject's details diff is 5
				increment := (i * 5)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[10+increment], row), subject.Name)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[11+increment], row), subject.Score)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[12+increment], row), subject.Grade)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[13+increment], row), subject.ClassCategory)
				f.SetCellValue(sheetName, fmt.Sprintf("%s%d", cellNames[14+increment], row), subject.Passed)
			}

			row++
		}
	}

	f.SetActiveSheet(index)

	return f, nil
}

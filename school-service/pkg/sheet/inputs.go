package sheet

import (
	"fmt"

	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/domain"

	"github.com/xuri/excelize/v2"
)

func GetAllNames(fileName, sheetName string) ([]string, error) {

	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err)
	}

	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s sheet: %v", sheetName, err)
	}

	names := make([]string, len(rows)-1)

	for i := 1; i < len(rows); i++ {
		names[i-1] = rows[i][0]
	}

	return names, nil
}

func GetAllStates(fileName, sheetName string) ([]domain.State, error) {

	f, err := excelize.OpenFile(fileName)

	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %v", fileName, err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s sheet: %v", sheetName, err)
	}

	states := make([]domain.State, 3)

	var city, state string

	states[0].Name = "Kerala"
	states[1].Name = "Tamil Nadu"
	states[2].Name = "Karnataka"

	for i := 1; i < len(rows); i++ {

		city = rows[i][0]  //get city
		state = rows[i][1] // get state

		switch state {
		case "Kerala":
			states[0].Cities = append(states[0].Cities, city)
		case "Tamil Nadu":
			states[1].Cities = append(states[1].Cities, city)
		case "Karnataka":
			states[2].Cities = append(states[2].Cities, city)
		}

		// check map of districts have districts exist or not
		// if _, ok := districts[district]; !ok {
		// 	districts[district] = []string{city}
		// } else {
		// 	districts[district] = append(districts[district], city)
		// }
	}

	return states, nil
}

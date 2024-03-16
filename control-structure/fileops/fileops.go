package fileops 

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// somente funções ou variaveis que iniciam com a letra maiuscula podem ser importadas a partir desse package
func WriteFloatToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	valueString := string(data)

	value, err := strconv.ParseFloat(valueString, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}
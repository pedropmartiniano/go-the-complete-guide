package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for i, value := range strings {
		floatValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, errors.New("invalid Float64 value")
		}

		floats[i] = floatValue
	}

	return floats, nil
}

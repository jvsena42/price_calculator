package conversion

import (
	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	floatList := make([]float64, len(strings))

	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, err
		}

		floatList[stringIndex] = floatPrice
	}

	return floatList, nil
}

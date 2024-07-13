package conversions

import (
	"strconv"
)

func StringToFloat(input []string) ([]float64, error) {
	pricesInFloat64 := make([]float64, len(input))

	for _, value := range input {
		price, err := strconv.ParseFloat(value, 64)

		if err != nil {
			return nil, err
		}

		pricesInFloat64 = append(pricesInFloat64, price)
	}

	return pricesInFloat64, nil
}

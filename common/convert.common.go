package common

import "strconv"

func convertInt32(str string) int32 {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(num)
}
func convertFloat64(str string) float64 {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return floatValue
}
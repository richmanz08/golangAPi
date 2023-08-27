package common

import (
	"strconv"

	"gorm.io/gorm"
)

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
const (
    DefaultPage = 1
    DefaultPerPage = 10
)

func Paginate(db *gorm.DB, page, perPage int) (*gorm.DB, error) {
    if page < 1 {
        page = DefaultPage
    }

    if perPage < 1 {
        perPage = DefaultPerPage
    }

    offset := (page - 1) * perPage
    return db.Offset(offset).Limit(perPage), nil
}
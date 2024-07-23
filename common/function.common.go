package common

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func ConvertToIDSlice(idString string) ([]int, error) {
	idStrs := strings.Split(idString, ",")
	ids := make([]int, len(idStrs))
	for i, idStr := range idStrs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
}

func Paginate(db *gorm.DB, page, perPage int) (*gorm.DB, error) {

	 DefaultPage := 1
     DefaultPerPage := 10

    // Ensure the page number is at least the default page
    if page < 1 {
        page = DefaultPage
    }

    // Ensure the items per page is at least the default per page
    if perPage < 1 {
        perPage = DefaultPerPage
    }

    // Calculate the offset for the query
    offset := (page - 1) * perPage

  
    // Apply the offset and limit to the query
    paginatedQuery := db.Offset(offset).Limit(perPage)

    // Return the modified query and nil error
    return paginatedQuery, nil
}
package helper

import (
	"strings"
)

// ConvertSortingForGetProducts is func for converting sorting from query params into sort sql query.
// 'termurah' => 'price ASC'
func ConvertSortingForGetProducts(data []string) []string {
	result := []string{}

	for _, sort := range data {
		switch strings.ToLower(strings.ReplaceAll(sort, " ", "")) {
		case "termurah":
			result = append(result, "price ASC")
		case "termahal":
			result = append(result, "price DESC")
		case "name(a-z)":
			result = append(result, "name ASC")
		case "name(z-a)":
			result = append(result, "name DESC")
		case "terbaru":
			result = append(result, "created_at DESC")
		}
	}

	return result
}

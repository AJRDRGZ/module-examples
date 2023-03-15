package slices

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// SliceFilter filtra un slice de acuerdo con la función callback especificada
func SliceFilter[T constraints.Ordered](nums []T, callback func(T) bool) []T {
	result := make([]T, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	log(fmt.Sprintf("%s: the result is: %v", result))
	return result
}

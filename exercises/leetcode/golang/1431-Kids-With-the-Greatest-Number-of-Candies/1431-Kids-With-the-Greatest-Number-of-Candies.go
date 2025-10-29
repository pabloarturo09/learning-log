package main

import (
	"fmt"
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	maxNumberOfCandies := slices.Max(candies)
	for i := range len(candies) {
		result[i] = extraCandies+candies[i] >= maxNumberOfCandies
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

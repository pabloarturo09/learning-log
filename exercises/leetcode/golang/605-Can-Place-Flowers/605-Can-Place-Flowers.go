package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	var prev int
	var next int
	var counter int

	if len(flowerbed) > 1 {
		next = flowerbed[1]
	}

	for i := 0; i < len(flowerbed) && counter != n; i++ {
		if prev == 0 && flowerbed[i] == 0 && next == 0 {
			counter++
			prev = 1
		} else {
			prev = flowerbed[i]
		}

		if i <= len(flowerbed)-3 {
			next = flowerbed[i+2]
		} else {
			next = 0
		}
	}

	return counter == n
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
}

package main

import (
	"fmt"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var results [3]int32

	for _, value := range arr {

		if int(value) > 0 {
			results[0]++
		}

		if int(value) < 0 {
			results[1]++
		}

		if int(value) == 0 {
			results[2]++
		}
	}

	arrLenght := len(arr)

	for _, value := range results {
		fmt.Printf("%f\n", (float32(value) / float32(arrLenght)))
	}
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(arr)
}

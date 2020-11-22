package main

import (
	"fmt"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) [3]float32 {
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

	arrLenght := float32(len(arr))

	var output [3]float32
	for index, value := range results {
		output[index] = float32(value) / arrLenght
		fmt.Printf("%f\n", output[index])
	}

	return output
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	result := plusMinus(arr)

	fmt.Printf("Output: %f\n", result)
}

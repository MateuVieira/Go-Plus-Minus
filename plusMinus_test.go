package main

import (
	"reflect"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {

	t.Run("Base case", func(t *testing.T) {
		arr := []int32{-4, 3, -9, 0, 4, 1}

		got := plusMinus(arr)
		want := [3]float32{0.500000, 0.333333343267, 0.166666671634}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %f want %f given", got, want)
		}
	})
}

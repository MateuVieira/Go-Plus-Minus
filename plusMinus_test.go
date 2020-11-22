package main

import (
	"reflect"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {

	t.Run("Base case", func(t *testing.T) {
		arr := []int32{-4, 3, -9, 0, 4, 1}

		got := plusMinus(arr)
		want := [3]float32{0.500000, 0.333333, 0.166667}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Teste - type got: %s", reflect.TypeOf(got).String())
			t.Errorf("Teste - type want: %s", reflect.TypeOf(want).String())
			for index, value := range got {
				t.Errorf("Teste got: %f - type: %s /  want: %f - type: %s", value, reflect.TypeOf(value).String(), want[index], reflect.TypeOf(want[index]).String())
			}
			t.Errorf("got %f want %f given", got, want)
		}
	})
}

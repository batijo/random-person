package age

import (
	"reflect"
	"testing"
)

func Test_fillWeights(t *testing.T) {
	tests := []struct {
		weights []weightData
		exp     []weightData
	}{
		{
			[]weightData{
				{Age: 1, Weight: 1},
				{Age: 3, Weight: 2},
				{Age: 4, Weight: 10},
				{Age: 5, Weight: 12},
				{Age: 8, Weight: 10},
			},
			[]weightData{
				{Age: 1, Weight: 1},
				{Age: 2, Weight: 1},
				{Age: 3, Weight: 2},
				{Age: 4, Weight: 10},
				{Age: 5, Weight: 12},
				{Age: 6, Weight: 12},
				{Age: 7, Weight: 12},
				{Age: 8, Weight: 10},
			},
		},
	}
	for _, d := range tests {
		var newWeights = d.weights
		fillWeights(&newWeights)
		if !reflect.DeepEqual(newWeights, d.exp) {
			t.Errorf(
				"DATA: %v EXPECTED: %v, GOT: %v",
				d.weights,
				d.exp,
				newWeights,
			)
		}
	}
}

package colors

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	{
		t.Run("", func(t *testing.T) {
			tests := []struct {
				Input  string
				Output string
			}{
				{
					Input:  "red",
					Output: "green",
				},
			}
			for _, test := range tests {
				opInd, _ := OpositeColorNameIndexFromColorName(test.Input)
				if actual, _ := ColorNameFromIndex(opInd); actual != test.Output {
					fmt.Println(actual, test.Output)
					t.Errorf("expected>>%s<<, but got>>%s<<", test.Output, actual)
				}
			}
		})
	}
}

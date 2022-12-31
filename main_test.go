package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	unittest = true

	fmt.Println("ON DEVELOPMENT STAGE")
	// for _, test := range []struct {
	// 	Args   []string
	// 	Output string
	// }{
	// 	{
	// 		Args:   []string{"./reloaded", "test/inHex.txt", "test/temp.txt"},
	// 		Output: "Hex test 30",
	// 	},
	// } {
	// 	t.Run("", func(t *testing.T) {
	// 		os.Args = test.Args
	// 		out = bytes.NewBuffer(nil)
	// 		main()
	// 		if actual := out.(*bytes.Buffer).String(); actual != test.Output {
	// 			fmt.Println(actual, test.Output)
	// 			t.Errorf("expected>>%s<<, but got>>%s<<", test.Output, actual)
	// 		}
	// 	})
	// }
}

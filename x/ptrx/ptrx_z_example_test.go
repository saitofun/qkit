package ptrx_test

import (
	"fmt"

	"github.com/saitofun/qkit/x/ptrx"
)

func Example() {
	fmt.Println(*(ptrx.Int(10)))
	fmt.Println(*(ptrx.Float64(10)))
	fmt.Println(*(ptrx.String("abc")))

	// Output:
	// 10
	// 10
	// abc
}

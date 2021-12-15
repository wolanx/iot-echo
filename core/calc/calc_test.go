package calc

import (
	"fmt"

	_ "github.com/zx5435/iot-echo/log"
)

var store = make(map[string]interface{})

func Example_a1() {
	fmt.Println(Calc("1 + 1", store))
	// Output: 2
}

func Example_b2() {
	fmt.Println(Calc("2 * 3", store))
	// Output: 6
}

func Example_a3() {
	store["cpu"] = 222.
	store["a"] = 3.
	fmt.Println(Calc("2 * a", store))
	// Output: 6
}

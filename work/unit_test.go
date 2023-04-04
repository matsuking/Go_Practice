

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Divide(x, y int) float32 {
	if y == 0 {
		return 0
	}
	return float32(x) / float32(y)
}

func unit_test() {
	x, y := 3, 2
	fmt.Printf("%v %v\n", Add(x, y), Divide(x, y))
}

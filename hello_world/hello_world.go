package main

import "fmt"

func main() {
	str := "FooBar"
	intNum := 123
	floatNum := 1.23
	complexNum := 42 + 1i
	fmt.Println("Hello, World!", str, intNum, floatNum, complexNum)
	// Hello, World! FooBar 123 1.23 (42+1i)

	array := []int{1, 2, 3}
	mapObj := map[string]int16{
		"a": 1,
		"b": 2,
	}
	fmt.Println(array, mapObj)
	// [1 2 3] map[a:1 b:2]

	array = append(array, 4)
	mapObj["c"] = 3
	fmt.Println(array, len(array), mapObj, len(mapObj))
	// [1 2 3 4] 4 map[a:1 b:2 c:3] 3

	a, b := demo(42)
	fmt.Println("demo(42):", a, b)
	// demo(42): 84 doubled
}

func demo(param int) (int, string) {
	return 2 * param, "doubled"
}

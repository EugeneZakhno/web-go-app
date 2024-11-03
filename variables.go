package main

import "fmt"

func basicTypes() {
	var b bool = true
	var s string = `ghrh`
	var i int = 1
	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359
	var r rune = '陆'
	fmt.Println(b, s, i, f32, f64, r)
}

func shortVars() {
	b := false
	s := `hello`
	i := 1
	f64 := 3.14159265359
	r := '码'
	fmt.Println(b, s, i, f64, r)
}

func compositeTypes() {
	nums := [3]int{1, 2, 3}
	fmt.Println(nums[1])

	words := []string{"incurring", "charge", "Decent", "despite", "e", "f", "g"}
	words = append(words, `go`)
	fmt.Println(words)

	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(dict["one"])

}

func main() {
	basicTypes()
	shortVars()
	compositeTypes()
}

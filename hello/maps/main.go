package main

import "fmt"

func main() {
	var n map[string]int
	fmt.Println("n:", n, "n[\"a\"]", n["a"])
	fmt.Println("len(n):", len(n)) // 0
	// n["b"] = 1 // will panic
	

	m := make(map[string]int)
	fmt.Println("m:", m, "m[\"a\"]", m["a"])

	m["a"] = 1
	v, ok := m["a"]
	fmt.Println("v:", v, "ok:", ok)

	v, ok = m["b"]
	fmt.Println("v:", v, "ok:", ok)

	fmt.Println("len(m):", len(m))

	delete(m, "a")
	delete(m, "b")
	fmt.Println("len(m):", len(m))

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

}
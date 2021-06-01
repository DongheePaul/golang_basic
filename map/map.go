package main

import (
"fmt"
)

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1024] = "1kbytes"
	m[4096] = "4kbytes"
	m[9162] = "9kbytes"

	str := m[1024]
	fmt.Println(str)

	noData := m[999] //값이 없으면 nil 혹은 zero 리턴
	fmt.Println(noData)

	var n map[string]int
	n = make(map[string]int)
	n["paul"] = 1
	value := n["paul"]
	
	fmt.Println(value)
	fmt.Println(n["jason"]) //0 리턴

	value, exists := n["paul"]
	fmt.Println(value, exists)
	value, exists = n["json"]
	fmt.Println(value, exists)
	
	delete(n, "paul")
	value, exists = n["paul"]

	fmt.Println(value, exists)
}
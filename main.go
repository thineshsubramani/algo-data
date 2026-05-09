package main

import (
	"fmt"

	"algo-data/generator"
	"algo-data/linkedlist"
)

func main() {
	nums := generator.RandomRange(10, 1, 100)

	fmt.Println(nums)

	head := linkedlist.FromSlice(nums)

	linkedlist.Print(head)
}
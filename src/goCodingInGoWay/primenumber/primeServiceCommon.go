package main

import "fmt"

const (
	N = 20
)

func generate1() []int {
	var array []int = make([]int, N-1)
	for i := 2; i <= N; i++ {
		array[i-2] = i
	}
	return array
}

func filter1(src []int, prime int) []int {
	var array []int = make([]int, N-1)
	var j = 0
	for i := 1; i < len(src); i++ {
		if src[i]%prime != 0 {
			array[j] = src[i]
			j = j + 1
		}
	}
	return array[0:j]
}

func service1(src []int) {
	if (len(src) == 0) {
		return
	}

	prime := src[0]
	fmt.Print(prime, "\n")
	src = filter1(src, prime)
	service1(src)
}


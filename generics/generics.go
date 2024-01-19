package main

import (
	"fmt"
	"strings"
)

func mapInt(arr []int, f func(int) int) []int {
	result := make([]int, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	rs := mapInt(arr, func(i int) int {
		return i * 2
	})

	fmt.Println(rs)

	arrAny := []string{"get", "post"}

	rsAny := mapAny(arrAny, func(i string) string {
		return strings.ToUpper(i)
	})

	fmt.Println(rsAny)
}

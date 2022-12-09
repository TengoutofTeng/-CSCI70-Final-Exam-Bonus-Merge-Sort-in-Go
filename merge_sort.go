package main

import "fmt"

//Implementation taken directly from: https://blog.boot.dev/golang/merge-sort-golang/

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func main() {

	//Edit sort "input" here! (inputted as whole integers separated by a comma and a space)
	unsorted := []int{-2, 4, 52, 48, 127, 44}
	sorted := mergeSort(unsorted)
	fmt.Print(sorted)

}

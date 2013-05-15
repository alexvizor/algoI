package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 9, 7, 0, 8, 20, 11}

	fmt.Println(merge_sort(arr))
}

func merge_sort(array []int) []int {
	al := len(array)

	if al > 1 {
		med := al / 2
		
		left := merge_sort(array[:med])
		right := merge_sort(array[med:])

		return merge(left, right)
	}

	return array[0:1]
}

func merge(left, right []int) []int {
	i, j, ll, rl := 0, 0, len(left), len(right)
	res := make([]int, 0);

	for i < ll && j < rl {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else if right[j] < left[i] {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}


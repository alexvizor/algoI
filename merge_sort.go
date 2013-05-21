package main

import "fmt"
import "os"
import "strconv"

// Test data: 1 4 5 9 7 0 8 20 11
func main() {
	arr := make([]int, 0)
	for _, val := range os.Args[1:] {
		v, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		arr = append(arr, v)
	}

	fmt.Println(merge_sort(arr))
}

func merge_sort(array []int) []int {
	al := len(array)

	if al == 0 {
		return array
	}

	if al > 1 {
		med := al / 2

		left, right := make(chan []int), make(chan []int)
		go func() {
			left <- merge_sort(array[:med])
		}()
		go func() {
			right <- merge_sort(array[med:])
		}()

		return merge(<-left, <-right)
	}

	return array[0:1]
}

func merge(left, right []int) []int {
	i, j, ll, rl := 0, 0, len(left), len(right)
	res := make([]int, 0)

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

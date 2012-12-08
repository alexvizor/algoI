package main

import "fmt"

func main() {
	arr := [...]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

	fmt.Println(arr)
	sub, max := find_max_subarray(arr[:])
	fmt.Println("Max sum:", max, "sub-array", sub)
}

func find_max_subarray(array []int) ([]int, int) {
	l := len(array)
	if l == 1 {
		return array, array[0]
	}

	mid := l / 2
	left, lmax := find_max_subarray(array[:mid])
	right, rmax := find_max_subarray(array[mid:])
	cros, cmax := find_max_cross(array)

	if lmax > cmax && lmax > rmax {
		return left, lmax
	} else if rmax > cmax && rmax > lmax {
		return right, rmax
	}
	return cros, cmax
}

func find_max_cross(array []int) ([]int, int) {
	l := len(array)
	mid := l / 2

	max_left, left := -1000, mid-1
	for i, sum := mid-1, 0; i >= 0; i-- {
		sum += array[i]
		if max_left < sum {
			max_left = sum
			left = i
		}
	}

	max_right, right := -1000, mid
	for i, sum := mid, 0; i < l; i++ {
		sum += array[i]
		if max_right < sum {
			max_right = sum
			right = i
		}
	}

	return array[left:right], max_left + max_right
}

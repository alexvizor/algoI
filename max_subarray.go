package main

import "fmt"
import "math"
import "os"
import "strconv"

type result struct {
	sub []int
	max int
}

// Text data: 13 -3 -25 20 -3 -16 -23 18 20 -7 12 -5 -22 15 -4 7
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

	fmt.Println("Original array:", arr)
	sub := find_max_subarray(arr[:])
	fmt.Println("Max sum:", sub.max, "sub-array", sub.sub)
}

func find_max_subarray(array []int) result {
	ln := len(array)
	if ln == 0 {
		return result{array, 0}
	}
	if ln == 1 {
		return result{array, array[0]}
	}

	mid := ln / 2
	leftc, rightc, crossc := make(chan result), make(chan result), make(chan result)
	go func() {
		leftc <- find_max_subarray(array[:mid])
	}()
	go func() {
		rightc <- find_max_subarray(array[mid:])
	}()
	go func() {
		crossc <- find_max_cross(array)
	}()
	l := <-leftc
	r := <-rightc
	c := <-crossc

	if l.max > c.max && l.max > r.max {
		return l
	} else if r.max > c.max && r.max > l.max {
		return r
	}
	return c
}

func find_max_cross(array []int) result {
	l := len(array)
	mid := l / 2

	max_left, left := math.MinInt32, mid-1
	for i, sum := mid-1, 0; i >= 0; i-- {
		sum += array[i]
		if max_left < sum {
			max_left = sum
			left = i
		}
	}

	max_right, right := math.MinInt32, mid
	for i, sum := mid, 0; i < l; i++ {
		sum += array[i]
		if max_right < sum {
			max_right = sum
			right = i
		}
	}

	return result{array[left:right], max_left + max_right}
}

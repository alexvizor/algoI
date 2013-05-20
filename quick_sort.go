package main

import "fmt"
import "os"
import "io"
import "flag"
import "bufio"
import "strconv"

type pivot_selector_func func([]int) int

var path = flag.String("path", "QuickSort.txt", "Path to data file")

func main() {
	fmt.Println("Excercise one number of comparisons:", excercise1())
	fmt.Println("Excercise two number of comparisons:", excercise2())
	fmt.Println("Excercise three number of comparisons:", excercise3())
}

func excercise1() uint {
	array := populate_array()
	return quick_sort(array, func(array []int) int {
		return 0
	})
}

func excercise2() uint {
	array := populate_array()
	return quick_sort(array, func(array []int) int {
		return len(array) - 1
	})
}

func excercise3() uint {
	array := populate_array()
	return quick_sort(array, func(array []int) int {
		ln := len(array)
		last := ln - 1
		med := ln / 2
		if ln%2 == 0 {
			med -= 1
		}

		switch {
		case between(array[med], array[0], array[last]):
			return med
		case between(array[0], array[med], array[last]):
			return 0
		}
		return last
	})
}

func between(seek, left, right int) bool {
	return (seek < left && seek > right) || (seek > left && seek < right)
}

func quick_sort(array []int, pivot_selector pivot_selector_func) uint {
	var ln = uint(len(array))

	if ln <= 1 {
		return 0
	}
	com_count := ln - 1

	var i = uint(1)
	var pivot_index = pivot_selector(array)
	array[0], array[pivot_index] = array[pivot_index], array[0]
	var pivot = array[0]

	for j := i; j < ln; j++ {
		if array[j] <= pivot {
			array[j], array[i] = array[i], array[j]
			i++
		}
	}

	array[0], array[i-1] = array[i-1], array[0]

	res_ch := make(chan uint, 2)
	go func() {
		res_ch <- quick_sort(array[:i-1], pivot_selector)
	}()
	go func() {
		res_ch <- quick_sort(array[i:], pivot_selector)
	}()

	com_count += <-res_ch
	com_count += <-res_ch

	return com_count
}

func populate_array() []int {
	array := make([]int, 0, 10000)
	file, err := os.Open(*path)

	if err != nil {
		panic("Unable to open")
	}

	reader := bufio.NewReader(file)
	func() error {
		for {
			switch line, _, err := reader.ReadLine(); err {
			case nil:
				item, e := strconv.Atoi(string(line))
				if e == nil {
					array = append(array, item)
				}
			case io.EOF:
				return nil
			default:
				return err
			}
		}
		return nil
	}()
	file.Close()

	return array
}

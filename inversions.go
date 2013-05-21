package main

import "fmt"
import "os"
import "io"
import "bufio"
import "strconv"

type result struct {
	inv uint
	sub []int
}

// Input data in IntegerArray.txt
func main() {
	arr := make([]int, 0)

	reader := bufio.NewReader(os.Stdin)
	func() error {
		for {
			switch line, _, err := reader.ReadLine(); err {
			case nil:
				item, e := strconv.Atoi(string(line))
				if e == nil {
					arr = append(arr, item)
				}
			case io.EOF:
				return nil
			default:
				return err
			}
		}
		return nil
	}()

	fmt.Println("Input array contains", len(arr), "items")
	inv := get_invertions(arr).inv
	fmt.Println("Inversion count:", inv)
}

func get_invertions(array []int) result {
	al := len(array)

	if al == 0 {
		return result{0, array}
	}

	if al > 1 {
		med := al / 2

		left, right := make(chan result), make(chan result)
		go func() {
			left <- get_invertions(array[:med])
		}()
		go func() {
			right <- get_invertions(array[med:])
		}()
		l, r := <-left, <-right

		return compute_inverts(l.sub, r.sub, l.inv+r.inv)
	}

	return result{0, []int{array[0]}}
}

func compute_inverts(left, right []int, inv uint) result {
	var i, j, ll, rl = uint(0), uint(0), uint(len(left)), uint(len(right))
	res := make([]int, 0)

	for i < ll && j < rl {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else if right[j] < left[i] {
			res = append(res, right[j])
			j++
			inv += ll - i
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return result{inv, res}
}

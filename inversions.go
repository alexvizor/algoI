package main

import "fmt"
import "os"
import "io"
import "bufio"
import "strconv"

// Input data is in IntegerArray.txt
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
	inv, _ := get_invertions(arr)
	fmt.Println("Inversion count:", inv)
}

func get_invertions(array []int) (uint, []int) {
	al := len(array)

	if al == 0 {
		return 0, array
	}

	if al > 1 {
		med := al / 2

		linv, left := get_invertions(array[:med])
		rinv, right := get_invertions(array[med:])

		return compute_inverts(left, right, linv+rinv)
	}

	return 0, []int{array[0]}
}

func compute_inverts(left, right []int, inv uint) (uint, []int) {
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

	return inv, res
}

package main

import "fmt"
import "os"
import "io"
import "log"
import "flag"
import "bufio"
import "strconv"

var path = flag.String("path", "HashInt.txt", "Path to data file")

func main() {
	sums := [...]int{231552, 234756, 596873, 648219, 726312, 981237, 988331, 1277361, 1283379}
	hash, result := load_data(), make([]int, len(sums))

	for key, _ := range hash {
		for index, item := range sums {
			other := item - key
			if _, exists := hash[other]; exists {
				result[index] = 1
			}
		}
	}

	fmt.Println(result)
}

func load_data() map[int]int {
	numbers := make(map[int]int)

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal("Unable to open file")
	}

	reader := bufio.NewReader(file)
	err = func() error {
		for {
			switch line, _, err := reader.ReadLine(); err {
			case nil:
				item, err := strconv.Atoi(string(line))
				if err != nil {
					return err
				}

				numbers[item] = 1
			case io.EOF:
				return nil
			default:
				return err
			}
		}
		return nil
	}()

	if err != nil {
		log.Fatal("Unable to parse file", err)
	}

	return numbers
}

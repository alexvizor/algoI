package main

import "io"
import "os"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"
import "sort"

func main() {
	fmt.Println("Max vertecis:", find_max_sccs())
}

func find_max_sccs() []int {
	sccs := find_sccs(parse())
	return sccs[:5]
}

func find_sccs(vs []*vertex) []int {
	fts := make([]int, 0)
	starting := len(vs) - 1

	dfs_loop(vs, func() (int, bool) {
		starting--
		return starting + 1, starting < 0
	}, func(index int) []int { return vs[index].backrefs }, func(index int) { fts = append(fts, index) })

	reverse(fts)

	for _, vx := range vs {
		vx.visited = false
	}

	dfs_loop(vs, func() (index int, end bool) {
		if len(fts) == 0 {
			return 0, true
		}
		index, fts = fts[0], fts[1:]
		return
	}, func(index int) []int { return vs[index].siblings }, func(int) {})

	scc_map := make(map[int]int, 5)
	for _, vx := range vs {
		switch _, ok := scc_map[vx.leader]; ok {
		case true:
			scc_map[vx.leader] += 1
		case false:
			scc_map[vx.leader] = 1
		}
	}

	sccs := make([]int, 5)
	for _, val := range scc_map {
		sccs = append(sccs, val)
	}
	sort.Ints(sccs)
	reverse(sccs)

	return sccs
}

func dfs_loop(vs []*vertex, selector func() (int, bool), dfs_selector func(int) []int, fts func(index int)) {
	index, end := selector()
	for !end {
		if !vs[index].visited {
			dfs(vs, index, index, dfs_selector, fts)
		}
		index, end = selector()
	}
}

func dfs(vs []*vertex, start, leader int, selector func(int) []int, fts func(index int)) {
	vs[start].visited = true
	vs[start].leader = leader
	for _, vindex := range selector(start) {
		if !vs[vindex].visited {
			dfs(vs, vindex, leader, selector, fts)
		}
	}
	fts(start)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type vertex struct {
	visited  bool
	leader   int
	siblings []int
	backrefs []int
}

func NewVertex() *vertex {
	return &vertex{false, -1, make([]int, 0), make([]int, 0)}
}

func print_graph(vs []*vertex) {
	for ind, vx := range vs {
		fmt.Println(ind, vx)
	}
}

func parse() []*vertex {
	file, err := os.Open("SCC.txt")
	if err != nil {
		log.Fatal("Error openning file:", err)
	}

	bio := bufio.NewReader(file)
	vs := make([]*vertex, 875714)
	err = func() error {
		for {
			switch line, err := bio.ReadString('\n'); err {
			case nil:
				tokens := func() (data []int) {
					fields := strings.Fields(line)
					for _, field := range fields {
						val, _ := strconv.Atoi(field)
						data = append(data, val)
					}
					return
				}()

				lindex, rindex := tokens[0]-1, tokens[1]-1

				if vs[lindex] == nil {
					vs[lindex] = NewVertex()
				}
				if vs[rindex] == nil {
					vs[rindex] = NewVertex()
				}

				vs[lindex].siblings = append(vs[lindex].siblings, rindex)
				vs[rindex].backrefs = append(vs[rindex].backrefs, lindex)
			case io.EOF:
				return nil
			default:
				return err
			}
		}
		return nil
	}()

	if err != nil {
		log.Fatal("Error occured while reading file:", err)
	}

	return vs
}

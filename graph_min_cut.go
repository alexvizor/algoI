package main

import "os"
import "io"
import "fmt"
import "log"
import "bufio"
import "strings"
import "strconv"
import "math/rand"

func main() {
    fmt.Println("Min cut:", find_min_cut())
}

func find_min_cut() int {
    min_cut := 1600
    for i := 0; i < 10; i ++ {
        vertecis, edges := parse_graph()
        cut := mincut(vertecis, edges)
        
        if cut < min_cut { min_cut = cut }
    }
    return min_cut
}

func mincut(vertecis, edges[][]int) int {
    ln, vln := len(edges), len(vertecis)
    for vln > 2 {
        edge_indecis := make([]int, 0)
        for index, edge := range edges {
            if edge != nil {
                edge_indecis = append(edge_indecis, index)
            }
        }
        rand.Seed(rand.Int63())
        rei := edge_indecis[rand.Intn(len(edge_indecis) - 1)];
        
        lver, rver := edges[rei][0], edges[rei][1]
        for index, edge := range edges {
            if edge != nil {
                if edges[index][0] == lver {
                    edges[index][0] = rver
                } else if edges[index][1] == lver {
                    edges[index][1] = rver
                }
                if edge[0] == edge[1] {
                    edges[index] = nil
                    ln--
                }
            }
        }
        vln--
    }
    
    return ln / 2
}

func print_edges(edges [][]int) {
    fmt.Println("============ Edges =====================")
    for i, v := range edges {
        fmt.Println(i, v)
    }
    fmt.Println("========================================")
}

func print_graph(vertecis, edges [][]int) {
    for index, ver := range vertecis {
        fmt.Print(index + 1, " ")
        for _, ed := range ver {
            fmt.Print(edges[ed][1], " ")
        }
        fmt.Println("")
    }
}

func parse_graph() (vertecis, edges [][]int) {
    file, err := os.Open("kargerAdj.txt")
    if err != nil {
        log.Fatal("Unable to open file")
    }
    
    reader := bufio.NewReader(file)
    edge_index, vertex_index := 0, 0
    func () error {
        for {
            switch line, _, err := reader.ReadLine(); err {
                case nil:
                    tokens := strings.Fields(string(line))
                    
                    vertecis = append(vertecis, make([]int, 0))
                    vertex, _ := strconv.Atoi(tokens[0])
                    for i := 1; i < len(tokens); i++ {
                        edge := make([]int, 2)
                        edge[0] = vertex - 1
                        edge[1], _ = strconv.Atoi(tokens[i])
                        edge[1] -= 1
                        
                        edges = append(edges, edge)
                        vertecis[vertex_index] = append(vertecis[vertex_index], edge_index)
                        edge_index++
                    }
                    vertex_index++
                case io.EOF:
                    return nil
                default:
                    return err
            }
        }
        return nil
    }()
    
    return vertecis, edges
}
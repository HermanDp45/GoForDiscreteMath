package main

import "fmt"

func DFS(v int, p int, timer int, used *[]bool, graph *[][]int, tin *[]int, fup *[]int) int {
	(*used)[v] = true
	count := 0
	timer++
	(*tin)[v] = timer
	(*fup)[v] = timer

	for i := 0; i < len((*graph)[v]); i++ {
		to := (*graph)[v][i]
		if to == p {
			continue
		}
		if (*used)[to] {
			if (*fup)[v] > (*fup)[to] {
				(*fup)[v] = (*fup)[to]
			}
		} else {
			count += DFS(to, v, timer, used, graph, tin, fup)
			if (*fup)[v] > (*fup)[to] {
				(*fup)[v] = (*fup)[to]
			}
			if (*fup)[to] > (*tin)[v] {
				count++
				fmt.Println("+1", v, to)
			}
		}
	}
	return count
}

func FindBridge(N, timer int, used *[]bool, graph *[][]int, tin, fup *[]int) int {
	count := 0
	for i := 0; i < N; i++ {
		(*used)[i] = false
	}
	for i := 0; i < N; i++ {
		if !(*used)[i] {
			count = DFS(i, -1, timer, used, graph, tin, fup)
		}
	}

	return count
}

func main() {
	var N, M int
	var inx, iny int
	fmt.Scan(&N, &M)
	var graph [][]int = make([][]int, N, N)
	for i := 0; i < M; i++ {
		fmt.Scan(&inx, &iny)
		graph[inx] = append(graph[inx], iny)
		graph[iny] = append(graph[iny], inx)
	}

	used := make([]bool, N, N)
	tin, fup := make([]int, N, N), make([]int, N, N)

	fmt.Print(FindBridge(N, 0, &used, &graph, &tin, &fup))

}

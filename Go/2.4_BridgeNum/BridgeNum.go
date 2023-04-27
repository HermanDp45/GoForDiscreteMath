package main

import "fmt"

func DFS(v int, p int, timer int, used *[]bool, graph *[][]int, timeIn *[]int, timeMin *[]int) int {
	(*used)[v] = true
	count := 0
	timer++
	(*timeIn)[v] = timer
	(*timeMin)[v] = timer

	for i := 0; i < len((*graph)[v]); i++ {
		to := (*graph)[v][i]
		if to == p {
			continue
		}
		if (*used)[to] {
			if (*timeMin)[v] > (*timeMin)[to] {
				(*timeMin)[v] = (*timeMin)[to]
			}
		} else {
			count += DFS(to, v, timer, used, graph, timeIn, timeMin)
			if (*timeMin)[v] > (*timeMin)[to] {
				(*timeMin)[v] = (*timeMin)[to]
			}
			if (*timeMin)[to] > (*timeIn)[v] {
				count++
				fmt.Println("+1", v, to)
			}
		}
	}
	return count
}

func FindBridge(N, timer int, used *[]bool, graph *[][]int, timeIn, timeMin *[]int) int {
	count := 0
	for i := 0; i < N; i++ {
		(*used)[i] = false
	}
	for i := 0; i < N; i++ {
		if !(*used)[i] {
			count = DFS(i, -1, timer, used, graph, timeIn, timeMin)
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
	timeIn, timeMin := make([]int, N, N), make([]int, N, N)

	fmt.Print(FindBridge(N, 0, &used, &graph, &timeIn, &timeMin))

}

//https://e-maxx.ru/algo/bridge_searching

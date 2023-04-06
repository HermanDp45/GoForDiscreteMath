package main

import "fmt"

type Graph struct {
	E [][]int;
	V []Node;
};


type Node struct {
	num int;
	mark int;
};

func DFS(G *Graph){
	for _, val := range G.V {
		val.mark = -1
	}
	for _, val := range G.V {
		if val.mark == -1 {
			VisitVertex(G, &val)
		}
	}
}

func VisitVertex(G *Graph, val *Node){
	fmt.Println(val.num);
	val.mark = 0;
	for _, line := range G.E{
		for ind2 := 0; ind2 < len(line); ind2++ {
			if G.V[ind2].mark == -1 {
				VisitVertex(G, &G.V[ind2]);
			}
		}
	}
	val.mark = 1;
}


func main(){
	var n int;
	var inp string;
	var m [][]int = make([][]int, 0, 0);
	var v []Node = make([]Node, 0, n);
	
	fmt.Scanln(&n);
	for i := 1; i <= n; i++ {
		v = append(v, Node{i, -1});
	}
	fmt.Println(v);
	for i := 0; i < n; i++{
		m = append(m, make([]int, 0, 1))
		for q := 0; q < n; q++{
			fmt.Scan(&inp)
			if inp == "-" {
				m[i] = append(m[i], 1);
			} else {
				m[i] = append(m[i], 0)
			}
		}
	}
	fmt.Println(m);

	var G Graph = Graph{m, v};
	DFS(&G);
	// for i := 0; i < n; i++ {
	// 	for q := 0
	// }
	
}
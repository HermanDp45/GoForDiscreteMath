package main

import "fmt"
// 0 - red, 1 - blue;
func DFS(graph *[][]int ,used *[]int, v int, color int, Sum[]) ([2]int) {
	(*used)[v] = color;
	fmt.Println(*used);
    for _, v1 := range (*graph)[v] {
        if (*used)[v1] == 0 {
			if color == 1 {
				color = 2;
			} else {
				color = 1;
			}
			DFS(graph, used, v1, color)
        }
    }
}


func main(){
	var n int;
	var inp string;

	var l [][]int = make([][]int, 0, 1)
	fmt.Scanln(&n);
	for i := 0; i < n; i++{
		l = append(l, make([]int, 0, 1))
		for q := 0; q < n; q++{
			fmt.Scan(&inp)
			if inp == "+" {
				l[i] = append(l[i], q);
			}
		}
	}

	fmt.Println(l)
	

	used := make([]int, n, n);
	fmt.Println(used);

	fmt.Println("START");
	for i := 0; i < n; i++ {
		if used[i] == 0 {
			DFS(&l, &used, i, 1);
		}
	}
	fmt.Println(used);
	fmt.Println("END");
	
}
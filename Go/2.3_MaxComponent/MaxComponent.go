package main

import "fmt"

func DFS(countVert int, listInc [][]int) []int {
	maxL := make([]int, 0, 2)
	filled := make([]int, countVert, countVert)
	for i := 0; i < countVert; i++ {
		if filled[i] == 0 {
			retList := VisitVertex(i, &filled, listInc, make([]int, 0, 2))
			if len(retList) > len(maxL) {
				maxL = retList
			}
		}
	}
	return maxL
}

func VisitVertex(indV int, filled *[]int, listInc [][]int, Ret []int) []int {
	(*filled)[indV] = 1
	Ret = append(Ret, indV)
	for i := 0; i < len(listInc[indV]); i++ {
		if (*filled)[listInc[indV][i]] == 0 {
			//fmt.Println("RET = ", Ret, "  i = ", i, "  vertexe : ", listInc[indV][i])
			Ret = VisitVertex(listInc[indV][i], filled, listInc, Ret)
		}
	}
	return Ret
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	list := make([][]int, N, N)
	connections := make([][]int, 0, 0)
	ina, inb := 0, 0
	for i := 0; i < M; i++ {
		fmt.Scan(&ina, &inb)
		connections = append(connections, []int{ina, inb})
		list[ina] = append(list[ina], inb)
		list[inb] = append(list[inb], ina)
	}
	ans := DFS(N, list)

	fmt.Println("graph {")
	for i := 0; i < len(list); i++ {
		boolForColor := false
		for _, val := range ans {
			if i == val {
				boolForColor = true
				break
			}
		}
		if boolForColor {
			fmt.Println("\t", i, "[color = red];")
		} else {
			fmt.Println("\t", i, ";")
		}
	}
	for _, val := range connections {
		boolForColor := false
		for _, vaq := range ans {
			if vaq == val[0] || vaq == val[1] {
				fmt.Println("\t", val[0], "--", val[1], "[color = red];")
				boolForColor = true
				break
			}
		}
		if !boolForColor {
			fmt.Println("\t", val[0], "--", val[1], ";")
		}
	}
	fmt.Println("}")

	// fmt.Println("grapgh {")
	// for i := 0; i < len(ans); i++ {
	// 	fmt.Println("\t", ans[i], ";")
	// }
	// connection := [][]int{}
	// for i := 0; i < len(ans); i++ {
	// 	for q := 0; q < len(list[ans[i]]); q++ {
	// 		con := []int{ans[i], list[ans[i]][q]}
	// 		reverseCon := []int{con[1], con[0]}
	// 		exists := false
	// 		for _, c := range connection {
	// 			if c[0] == reverseCon[0] && c[1] == reverseCon[1] {
	// 				exists = true
	// 				break
	// 			}
	// 		}
	// 		if !exists {
	// 			connection = append(connection, con)
	// 			fmt.Println("\t", con[0], "--", con[1], ";")
	// 		}
	// 	}
	// }
	// fmt.Println("}")

}

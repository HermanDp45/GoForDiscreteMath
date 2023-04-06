package main
import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	less = quickSort(less)
	greater = quickSort(greater)
	return append(append(less, pivot), greater...)
}

func contain(v int, team []int) bool {
	for i:= 0; i < len(team); i++ {
		if v == team[i] {
			return true;
		}
	}
	return false;
}
func check(v int, team []int, graph [][]int) bool {
	for i := 0; i < len(team); i++ {
		if graph[v][team[i]-1] == 1 {
			return false;
		}
	}
	return true
}

func checkNoConflicts(v int, graph [][]int) bool {
	for i := 0; i < len(graph[v]); i++ {
		if graph[v][i] == 1 {
			return false
		}
	}
	return true;
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
				l[i] = append(l[i], 1);
			} else {
				l[i] = append(l[i], 0);
			}
		}
	}
	ideal := make([]int, 0, 0);
	team1 := make([]int, 0, 0);
	team2 := make([]int, 0, 0);
	used := make([]int, n, n);
	for i := 1; i <= n; i++ {
		if checkNoConflicts(i-1,l){
			ideal = append(ideal, i)
		}
	}

	checkOpport := true;
	for i := 1; i <= n; i++ {
		for q := i+1; q <= n; q++ {
			if l[i-1][q-1] == 1 {
				if check(i - 1, team1, l) && check(q-1, team2, l) {
					if used[i-1] == 0 {
						team1 = append(team1, i)
						used[i-1] = 1;
					}
					if used[q-1] == 0 {
						team2 = append(team2, q)
						used[q-1] = 1;
					}
					continue;
				} else if check(i - 1, team2, l) && check (q-1, team1, l){
					if used[q-1] == 0 {
						team1 = append(team1, q)
						used[q-1] = 1;
					}
					if used[i-1] == 0 {
						team2 = append(team2, i)
						used[i-1] = 1;
					}
					continue;
				} else {
					fmt.Println("NoSolution");
					checkOpport = false;
					break;
				}
			} 
		}
		if !checkOpport {
			break;
		}
	}
	if checkOpport {
		for _, v := range ideal {
			if len(team1) < len(team2) {
				team1 = append(team1, v);
			} else if len(team1) == len(team2) {
				team2 = append(team2, v);
			} else {
				team2 = append(team2, v)
			}
		}
		fmt.Println(quickSort(team1));
	}
}
package main
import . "fmt"

func Partition(low, high int,  less func(i, j int) bool, swap func(i, j int)) int {
	i := low
	for j := low; j < high; j++ {
		if less(j,high) {
			swap(i,j)
			i++
		}
	}
	swap(i, high)
	return i
}

func QuickSortRec(low, high int,  less func(i, j int) bool, swap func(i, j int)) {
	if low < high {
		q := Partition(low, high, less, swap)
		QuickSortRec(low, q-1, less, swap)
		QuickSortRec(q+1, high, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	QuickSortRec(0, n-1, less, swap)
}

func main() {
	var n int
    Scan(&n)

    list := make([]int, 0, n)

    for i := 0; i < n; i++ {
        var x int
        Scan(&x)
        list = append(list, x) 
    }

	less := func(i, j int) bool {
        return list[i] < list[j]
    }

    swap := func(i, j int) {
        list[i], list[j] = list[j], list[i]
    }
	
	qsort(len(list), less, swap)
	for i := 0; i < n; i++ {
		Print(list[i], " ")
	}
	

}
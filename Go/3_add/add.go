package main

import . "fmt"


func add(a, b []int32, p int) []int32 {
	var maxlen int
	if len(a) <= len(b) {
		maxlen = len(b)
	} else {
		maxlen = len(a)
	}
	c := make([]int32, 0, maxlen + 1)
	var sum int32
	c = append(c, 0)
	for i := 0; i < maxlen; i++ {
		sum = 0
		if i < len(a) {sum += a[i]}
		if i < len(b) {sum += b[i]}
		sum += c[i]
		ost := sum % int32(p)
		whole := sum / int32(p)
		c[i] = ost
		c = append(c, whole)
	}
	if c[maxlen] == 0 { c = c[:maxlen]}
	Println(c)
	return c
}

func main(){	
	Println("TEST")
	A := []int32{7, 1, 2, 7, 4, 5, 6, 7, 4}
	B := []int32{9, 8, 7, 6, 5, 4, 3, 2, 7}

	Print(add(A ,B ,10))

}

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type AssocArray interface {
    Assign(s string, x int)
    Lookup(s string) (x int, exists bool)
}

type Node struct {
	key string;
	val int;
}

type Array struct {
	Elems []Node;
}

func (a *Array) Assign(key string, val int) {
	a.Elems = append(a.Elems, Node{key, val})
}

func (a *Array) Lookup(key string) (int, bool) {
	for _, v := range a.Elems {
		fmt.Println(string(v.key) == string(key));
		if (string(key) == string(v.key)) {
			fmt.Println()
			return v.val, true;
		}
	}
	return 0, false;
}

func lex(sentence string, array AssocArray) []int {
    massiv := strings.Fields(sentence)
	ans := make([]int, 0, 10);
	counter := 1;
	for i := 0; i < len(massiv); i++ {
		if num, ok := array.Lookup(massiv[i]); ok {
			ans = append(ans, num)
			
		} else {
			array.Assign(massiv[i], counter);
			ans = append(ans, counter);
			counter++
		}
	}
	return ans;
}

func main(){
	r := bufio.NewReader(os.Stdin);
	str, err := r.ReadString('\n');
	if err != nil {
		panic("Error with string reading.");
	}
	array := Array{make([]Node, 0, 10)}
	fmt.Println(str[:len(str) - 2] );
	ansArr := lex(str[:len(str) - 1] , &array);
	fmt.Println()
	for _, v := range ansArr {
		fmt.Print(v, " ");
	}
}
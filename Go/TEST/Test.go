package main

import (
	//"errors"
	"bufio"
	"fmt"
	"os"
	"strings"
	//"unicode"
	//"strconv"
)

func main() {
	//var users []string
	file, err := os.Open("C:/Users/edelk/Code/Go/TEST/text.txt");
	if err != nil {panic("Problem with opening the file");}
	defer file.Close();

	rd := bufio.NewReader(file);
	s, err := rd.ReadString('\n');
	if err != nil {panic("Problem with reading the file");}
	ars := strings.Split(s, ";");
	for i, v := range ars {
		fmt.Println(v)
		if string(v) == "0" {
			fmt.Println(i + 1);
			break
		} 
	}

}

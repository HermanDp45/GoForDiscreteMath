package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	//"strconv"
)

type Ans struct {
	Average float64
}

type SecStruct struct {
	LastName string		//:"ABA",
	FirstName string	//:"BABA",
	MiddleName string	//:"LADA",
	Birthday string 	//:"FEFSE",
	Address string		//:"ALALLAA",
	Phone string		//:"LALALALA",
	Rating []int		//:[1,2,3]
}

type Mystruct struct {
	ID int					//:134,
    Number string			//:"FFF1",
    Year int				//:2,
    Students []SecStruct	//:[
}

func main() {
	f, err := os.Open("C:/Users/edelk/Code/Go/JsonPractice/text.json");
	if err != nil {
		panic("Problem with opening file");
	}
	defer f.Close();

	scanner := bufio.NewScanner(f);
	var lines string;
	for scanner.Scan() {
		input := scanner.Text();
		if input == "" {
			break;
		}
		lines += input;
	}

	if err := scanner.Err(); err != nil {
        panic("Errors with input")
    }

	byteToJson := []byte(lines)
	var jsin Mystruct;

	if err := json.Unmarshal(byteToJson, &jsin); err != nil {
		panic("Errors with Unmarshal");
	} 

	countStud, countOcenk := len(jsin.Students), 0;
	for i := 0; i < len(jsin.Students); i++ {
		countOcenk += len(jsin.Students[i].Rating); 
	}

	sr := float64(countOcenk) / float64(countStud);

	ans := Ans{
		Average: sr,
	}

	data, err := json.MarshalIndent(ans, "", "    ")
	if err != nil {
		panic("problem with json.MarshalIndent");
	}

	fmt.Printf("%s",data)

}
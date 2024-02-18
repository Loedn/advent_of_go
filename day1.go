package main

import (
	"fmt"
	"os"
)

var counter int

func main() {
	data, err := os.ReadFile("day1.txt")
	if err != nil {
		return
	}
	var nu_string string = string(data)
	for i := 0; i < len(nu_string); i++ {
		if string(nu_string[i]) == "(" {
			counter += 1
		} else {
			counter -= 1
		}
		if counter == -1 {
			fmt.Println(i - 1) // counter starts at one index starts at 0
		}
	}
	fmt.Println(counter)
}

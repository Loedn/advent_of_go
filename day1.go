// Day 1

// package main

// import (
// 	"fmt"
// 	"os"
// )

// var counter int

// func main() {
// 	data, err := os.ReadFile("day1.txt")
// 	if err != nil {
// 		return
// 	}
// 	var nu_string string = string(data)
// 	for i := 0; i < len(nu_string); i++ {
// 		if string(nu_string[i]) == "(" {
// 			counter += 1
// 		} else {
// 			counter -= 1
// 		}
// 		if counter == -1 {
// 			fmt.Println(i - 1) // counter starts at one index starts at 0
// 		}
// 	}
// 	fmt.Println(counter)
// }

// Day 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOfThree(a, b, c int) int {
	return min(min(a, b), c)
}

func minTwoOfThree(a, b, c int) (int, int) {
	smallest := min(a, min(b, c))
	if smallest == a {
		if b < c {
			return a, b
		}
		return a, c
	} else if smallest == b {
		if a < c {
			return a, b
		}
		return b, c
	} else {
		if a < b {
			return a, c
		}
		return b, c
	}
}

func main() {

	readFile, err := os.Open("day2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	var area_result = 0
	var ribbon_result = 0
	for _, line := range fileLines {
		var recipe = strings.Split(line, "x")
		l, err := strconv.Atoi(recipe[0])
		w, err := strconv.Atoi(recipe[1])
		h, err := strconv.Atoi(recipe[2])
		if err != nil {
			fmt.Println(err)
		}
		var l_area int = l * w
		var w_area int = w * h
		var h_area int = h * l
		var smol_area int = minOfThree(l_area, w_area, h_area)
		smol1, smol2 := minTwoOfThree(l, w, h)
		area_result += 2*l_area + 2*w_area + 2*h_area + smol_area
		ribbon_result += l*w*h + smol1*2 + smol2*2
	}

	fmt.Println("total square feet of wrapping paper is ", area_result)
	fmt.Println("total feet of ribbon is ", ribbon_result)

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)


func main(){
	//part 1
	f1, err := os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(f1))
	f1.Close()


	//part 2
	f2, err := os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part2(f2))
	f2.Close()
}


func part1(f *os.File) (int) {
	depth := 0
	horizontal := 0

	scanner := bufio.NewScanner(f)
	
	for scanner.Scan(){
		s := scanner.Text()
		if s[0] == 'f' {
			//go forward
			horizontal = horizontal + extractNum(s)
		}
		if s[0] == 'u' {
			//decrease depth
			depth = depth - extractNum(s)
		}
		if s[0] == 'd' {
			//increase depth
			depth = depth + extractNum(s)
		}
	}

	return depth * horizontal
}


func part2(f *os.File) (int) {
	depth := 0
	horizontal := 0
	aim := 0

	scanner := bufio.NewScanner(f)
	
	for scanner.Scan(){
		s := scanner.Text()
		if s[0] == 'f' {
			//go forward
			increment := extractNum(s)
			horizontal += increment
			depth += aim * increment
		}
		if s[0] == 'u' {
			//decrease depth
			aim -= extractNum(s)
		}
		if s[0] == 'd' {
			//increase depth
			aim += extractNum(s)
		}
	}

	return depth * horizontal
}


func extractNum(s string) (int) {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result += string(s[i])
		}
	}

	returnInt, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return returnInt
}
  

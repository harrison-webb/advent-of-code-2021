package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main(){
	//part 1
	f1, err := os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1Better(f1))
	f1.Close()


	//part 2
	f2, err := os.Open("day2Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part2Better(f2))
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


//improved solution inspired by Liz Fong-Jones AOC stream
func part1Better(f *os.File) (int) {
	depth := 0
	horizontal := 0

	scanner := bufio.NewScanner(f)
	
	for scanner.Scan(){
		s := scanner.Text()
		split := strings.Split(s, " ")
		direction := split[0]
		num, _ := strconv.Atoi(split[1])
		switch direction {
			case "forward":
				horizontal += num
			case "up":
				depth -= num
			case "down":
				depth += num
		}
	}

	return depth * horizontal
}


//improved solution inspired by Liz Fong-Jones AOC stream
func part2Better(f *os.File) (int) {
	depth := 0
	horizontal := 0
	aim := 0

	scanner := bufio.NewScanner(f)
	
	for scanner.Scan(){
		s := scanner.Text()
		split := strings.Split(s, " ")
		direction := split[0]
		num, _ := strconv.Atoi(split[1])
		switch direction {
			case "forward":
				horizontal += num
				depth += aim * num
			case "up":
				aim -= num
			case "down":
				aim += num
		}
	}

	return depth * horizontal
}
  

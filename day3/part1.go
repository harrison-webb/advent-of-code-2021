package main

import (
   "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
	var gammaRate int = 0
	var epsilonRate int = 0
	var numberOfLines int = 0

	f, err := os.Open("day3Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var intValues [12]int //should make this dynammic but IDK HOW

	for scanner.Scan() {
		num := scanner.Text()
		for i, v := range num {
			fmt.Print(v - '0')
			intValues[i] += int(v - '0')
		}
		fmt.Println()
		numberOfLines++
	}

	for i, val := range(intValues) {
		if val > numberOfLines/2 {
			val = 1
		} else {
			val = 0
		}

		// 1<<n == 2^n
		// 0th index corresponds to 2^(len(intValues) - 1)
		gammaRate += val * 1<<(len(intValues) - 1 - i)
		epsilonRate += (1^val) * 1<<(len(intValues) - 1 - i)
	}

	fmt.Println(gammaRate)
	fmt.Println(epsilonRate)
	fmt.Println(gammaRate * epsilonRate)
	

}
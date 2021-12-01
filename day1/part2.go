package main

import (
   "bufio"
    "fmt"
    "log"
    "strconv"
    "os"
)

func main(){
  results, _ := ReadInts("aocDay1Input.txt")
  numGreater := 0
  var window1 int = results[0] + results[1] + results[2]
  var window2 int
  for i := 2; i < len(results) - 1; i++{
    window2 = results[i-1] + results[i] + results[i+1]
    if window2 > window1{
      numGreater++
    }
    window1 = window2
  }
  fmt.Println(numGreater)
}

func ReadInts(fname string) ([]int, error) {
  f, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
  defer f.Close()
  
  scanner := bufio.NewScanner(f)
  var result []int
  for scanner.Scan() {
    x, err := strconv.Atoi(scanner.Text())
    if err != nil {
        return result, err
    }
    result = append(result, x)
  }
  return result, scanner.Err()
}

//function to convert text file of numbers to an array
//elements inserted by lines of file
func textToIntArray(fname string) ([]int, error){
  f, err := os.Open(fname)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  var result []int
  for scanner.Scan() {
    x, err := strconv.Atoi(scanner.Text())
    if err != nil {
      return result, err
    }
    result = append(result, x)
  }
  return result, scanner.Err()
}

//function to convert text file to string array
//elements inserted by lines of file
func textToStringArray(fname string) ([]string, error){
  f, err := os.Open(fname)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  scanner := bufio.NewScanner(f)
  var result []string
  for scanner.Scan() {
    result = append(result, scanner.Text())
  }
  return result, scanner.Err()
}
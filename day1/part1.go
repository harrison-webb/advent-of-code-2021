package main

import (
   "bufio"
    "fmt"
    "log"
    "strconv"
    "os"
)

func main(){
  // open file
    f, err := os.Open("aocDay1Input.txt")
    if err != nil {
      log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    var prevMeasurement int = 100000;
    var numGreater = 0;
    for scanner.Scan() {
      // do something with a line scanner.Text()
      //check if current > prevMeasurement
      currentMeasurement, err := strconv.Atoi(scanner.Text())
      _ = err
      if currentMeasurement > prevMeasurement{
        numGreater++;
      }
      prevMeasurement = currentMeasurement
    }

    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }

    fmt.Println(numGreater)

}
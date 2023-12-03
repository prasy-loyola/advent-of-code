package main

import (
	"fmt"
    "os"
    "log"
)

func main() {
    filepath := "input-day1-1.txt"
    input, err := os.Open(filepath)
    if err != nil {
        log.Fatal("Couldn't open file", filepath, err)
    }

    sum := 0
    buf := make([]byte, 200)
    first := 0
    last := 0
    foundFirstDigit := false
    line := 0

    for size, err := input.Read(buf); err == nil && size != 0; size, err = input.Read(buf){
        for i := 0; i < size; i++ {
            var char byte = buf[i]
            if char >= '0' && char <= '9' {
                last = int(char) - '0'
                if !foundFirstDigit {
                    first = int(char) - '0'
                    foundFirstDigit = true
                }
            }
            if char == '\n' || char == 0 {
                coord := first * 10 + last
                sum += coord
                line++
//                fmt.Println(line, ": ", coord)
                foundFirstDigit = false
                first = 0
                last = 0
            }
        }
    }
    fmt.Println("Answer: ", sum)

}

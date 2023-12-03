package main

import (
	"fmt"
	"log"
	"os"
    "bufio"
)


func expect(line []byte, pos int, subtext string) (int, string) {

    if len(subtext) > len(line) - pos {
        return pos, ""
    }
    for i := 0; i < len(subtext); i++ {
        char1 := line[pos + i]
        char2 := subtext[i]
        if char1 != char2 {
            return pos, ""
        }
    }
    return pos + len(subtext) - 1, subtext
}


func parseLine(line []byte) int {

	nums := make([]int, 0)

	for i := 0; i < len(line); i++ {
		char := line[i]
		if char >= '0' && char <= '9' {
            nums = append(nums, int(char - '0'))
        } else if pos, word := expect(line, i, "zero"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 0)
        } else if pos, word := expect(line, i, "one"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 1)
        } else if pos, word := expect(line, i, "two"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 2)
        } else if pos, word := expect(line, i, "three"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 3)
        } else if pos, word := expect(line, i, "four"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 4)
        } else if pos, word := expect(line, i, "five"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 5)
        } else if pos, word := expect(line, i, "six"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 6)
        } else if pos, word := expect(line, i, "seven"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 7)
        } else if pos, word := expect(line, i, "eight"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 8)
        } else if pos, word := expect(line, i, "nine"); word != ""{
            i = pos - 1 // in case the last letter is shared by the next digit
            nums = append(nums, 9)
        }
        
	}
    length := len(nums)
    if length == 0 {
        return 0
    }
    coord := nums[0] * 10 + nums[length - 1]
	return coord

}

func main() {
	filepath := "input-day1-1.txt"
	input, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Couldn't open file", filepath, err)
	}

	sum := 0
    reader := bufio.NewReader(input)

    for {

        line, err := reader.ReadBytes('\n')
        sum += parseLine(line)
        if err != nil {
            fmt.Println(line, err)
            break;
        }
    }

	fmt.Println("Answer: ", sum)

}

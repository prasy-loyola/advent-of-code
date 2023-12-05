package main

import (
	"bufio"
	"log"
	"os"
	"parser"
)

func main() {
	filepath := "input.txt"
	var input *os.File
	var err error
	if input, err = os.Open(filepath); err != nil {
		panic("Couldn't read input file")
	}
	reader := bufio.NewReader(input)

	puzzle1Result := 0
	winsPerCard := make([]int, 0)
	copiesPerCard := make([]int, 0)
	for {
		var lineBytes []byte
		pos := 0
		if lineBytes, err = reader.ReadBytes('\n'); err != nil {
			break
		}
		line := string(lineBytes)
		cardNumber := 0
		//Card <card number>:
		if _, pos, err = parser.ExpectWord(line, pos, "Card"); err != nil {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}
		if pos, err = parser.Skip(line, pos, ' '); err != nil {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}
		if cardNumber, pos, err = parser.ExpectNumber(line, pos); err != nil {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}
		if _, pos, err = parser.ExpectWord(line, pos, ":"); err != nil {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}

		// <winNum1> <winNum2> <winNum3> <winNum4> <winNum5>
		winningNumbers := [10]int{}
		for i := 0; i < len(winningNumbers); i++ {
			if pos, err = parser.Skip(line, pos, ' '); err != nil {
				log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
			}
			var num int
			if num, pos, err = parser.ExpectNumber(line, pos); err != nil {
				log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
			}
			winningNumbers[i] = num
		}

		// | <num1> <num2> <num3> <num4> <num5> <num6> <num7> <num8>
		ourWinningNumbers := 0
		numbers := [25]int{}
		if _, pos, err = parser.ExpectWord(line, pos, " |"); err != nil {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}
		for i := 0; i < len(numbers); i++ {
			if pos, err = parser.Skip(line, pos, ' '); err != nil {
				log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
			}
			var num int
			if num, pos, err = parser.ExpectNumber(line, pos); err != nil {
				log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
			}
			numbers[i] = num

			for _, n := range winningNumbers {
				if n == num {
					ourWinningNumbers++
					break
				}
			}
		}

		if len(copiesPerCard) < cardNumber {
			copiesPerCard = append(copiesPerCard, 1)
		} else {
			copiesPerCard[cardNumber-1] += 1
		}

		winsPerCard = append(winsPerCard, ourWinningNumbers)
		//log.Printf("INFO: Card number: %d, winsPerCard: %v, copiesPerCard: %v", cardNumber, winsPerCard, copiesPerCard)
		if ourWinningNumbers > 0 {
			score := 1
			for x := 0; x < ourWinningNumbers-1; x++ {
				score *= 2
			}
			puzzle1Result += score
		}
		for x := 0; x < ourWinningNumbers; x++ {
			if len(copiesPerCard) < cardNumber+x+1 {
				copiesPerCard = append(copiesPerCard, 0)
			}
			copiesPerCard[cardNumber+x] += copiesPerCard[cardNumber-1]
		}
		//log.Printf("INFO: line: %s, ourWinning %d", line,ourWinningNumbers)
	}

	log.Printf("INFO: Puzzle1 result %d", puzzle1Result)
	puzzle2Result := 0
	for _, copies := range copiesPerCard {
		puzzle2Result += copies
	}
	log.Printf("INFO: Puzzle2 result %d", puzzle2Result)

}


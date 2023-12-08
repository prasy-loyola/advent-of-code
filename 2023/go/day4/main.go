package main

import (
	"bufio"
	"log"
	"os"
	parsermod "parser"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	filepath := "input.txt"
	var input *os.File
	var err error
	if input, err = os.Open(filepath); err != nil {
		panic("Couldn't read input file")
	}
	reader := bufio.NewReader(input)

	parser := parsermod.Parser{
		Reader: reader,
	}

	puzzle1Result := 0
	winsPerCard := make([]int, 0)
	copiesPerCard := make([]int, 0)
	var line string
	for {

		if err != nil && err != parsermod.EOL {
			log.Fatalf("ERROR: Invalid line '%s', \n\t%s", line, err.Error())
		}
		if line, err = parser.ReadNextLine(); err == parsermod.EOF {
			break
		}

		cardNumber := 0
		//Card <card number>:
		if _         , err = parser.ExpectWord("Card"); err != nil { continue }
		if _         , err = parser.Skip(' ');          err != nil { continue }
		if cardNumber, err = parser.ExpectNumber();     err != nil { continue }
		if _         , err = parser.ExpectWord(":");    err != nil { continue }

		// <winNum1> <winNum2> <winNum3> <winNum4> <winNum5>
		winningNumbers := [10]int{}
		for i := 0; i < len(winningNumbers); i++ {
			var num int
			if _  , err = parser.Skip(' ');      err != nil { continue }
			if num, err = parser.ExpectNumber(); err != nil { continue }
			winningNumbers[i] = num
		}

		// | <num1> <num2> <num3> <num4> <num5> <num6> <num7> <num8>
		ourWinningNumbers := 0
		numbers := [25]int{}
		if _, err = parser.ExpectWord(" |"); err != nil { continue }
		for i := 0; i < len(numbers); i++ {
			var num int
			if _  , err = parser.Skip(' ');      err != nil { continue }
			if num, err = parser.ExpectNumber(); err != nil { continue }
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

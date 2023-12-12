package main

import (
	"bufio"
	"log"
	"math"
	"os"
	parsermod "parser"
)

func solveQuadratic(totalTime, winningDistance int) (int, int) {

	quadPart := math.Sqrt(float64(totalTime*totalTime - 4*winningDistance))
	solution1 := int((float64(totalTime) - quadPart) / 2.0)
	solution2 := int((float64(totalTime) + quadPart) / 2.0)

	isWinning := func(holdTime int) bool {
		dist := (totalTime - holdTime) * holdTime
		return dist > winningDistance
	}

	if solution1 > 1 && isWinning(solution1-1) {
		solution1 -= 1
	} else if isWinning(solution1 + 1) {
		solution1 += 1
	}
	if !isWinning(solution2) {
		if solution2 > 1 && isWinning(solution2-1) {
			solution2 -= 1
		}
	} else if isWinning(solution2 + 1) {
		solution2 += 1
	}

	return int(solution1), int(solution2)
}

func main() {
	filepath := "input.txt"
	var err error
	var input *os.File
	if input, err = os.Open(filepath); err != nil {
		log.Fatalf("ERROR: Couldn't read file %s", filepath)
	}

	reader := bufio.NewReader(input)
	parser := parsermod.Parser{
		Reader: reader,
	}

	if _, err := parser.ReadNextLine(); err != nil {
		log.Fatalf("ERROR: invalid format %s", err.Error())
	}
	if _, err := parser.ExpectWord("Time:"); err != nil {
		log.Fatalf("ERROR: invalid format %s", err.Error())
	}
	var time []int
	var distance []int
	time, err = parser.ExpectListOfNumbers(' ')

	if _, err := parser.ReadNextLine(); err != nil {
		log.Fatalf("ERROR: invalid format %s", err.Error())
	}
	if _, err := parser.ExpectWord("Distance:"); err != nil {
		log.Fatalf("ERROR: invalid format %s", err.Error())
	}
	distance, err = parser.ExpectListOfNumbers(' ')

	log.Printf("INFO: time: %v, distance:%v", time, distance)

	puzzle1Result := 1
	for i := 0; i < len(time); i++ {
		sol1, sol2 := solveQuadratic(time[i], distance[i])
		log.Printf("INFO: solution for %d , time: %d, distance: %d, is [%d, %d]", i, time[i], distance[i], sol1, sol2)
		puzzle1Result *= (sol2 - sol1 + 1)
	}

    log.Printf("INFO: puzzle1Result %d" , puzzle1Result)

}

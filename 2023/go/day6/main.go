package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	parsermod "parser"
	"strconv"
)

func solveQuadratic(totalTime, winningDistance int64) (int64, int64) {

	quadPart := math.Sqrt(float64(totalTime*totalTime - 4*winningDistance))
	solution1 := int64((float64(totalTime) - quadPart) / 2.0)
	solution2 := int64((float64(totalTime) + quadPart) / 2.0)

	isWinning := func(holdTime int64) bool {
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

	return int64(solution1), int64(solution2)
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

	puzzle1Result := int64(1)
    puzzle2Time := ""
    puzzle2Distance := ""
	for i := 0; i < len(time); i++ {
        puzzle2Time = fmt.Sprint(puzzle2Time, time[i])
        puzzle2Distance = fmt.Sprint(puzzle2Distance, distance[i])
		sol1, sol2 := solveQuadratic(int64(time[i]), int64(distance[i]))
		log.Printf("INFO: solution for %d , time: %d, distance: %d, is [%d, %d]", i, time[i], distance[i], sol1, sol2)
		puzzle1Result *= (sol2 - sol1 + 1)
	}

    log.Printf("INFO: puzzle1Result: %d" , puzzle1Result)
    log.Printf("INFO: puzzle2Time: %s, puzzle2Distance: %s" , puzzle2Time, puzzle2Distance)

    puzzle2TimeI, _ := strconv.ParseInt(puzzle2Time, 10, 64)
    puzzle2DistanceI, _ := strconv.ParseInt(puzzle2Distance, 10, 64)
	sol1, sol2 := solveQuadratic(puzzle2TimeI, puzzle2DistanceI)
    puzzle2Result := sol2 - sol1 + 1
    log.Printf("INFO: puzzle2Result: %d" , puzzle2Result)


}

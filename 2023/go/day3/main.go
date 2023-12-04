package main

import (
	"bufio"
	"log"
	"os"
	"parser"
)

type enginePart = struct {
	number int
	row    int
	start  int
	end    int
	valid  bool
}

type marker = struct {
	markerType  byte
	row         int
	position    int
	nearbyParts []*enginePart
}

func main() {
	filepath := "input.txt"
	input, err := os.Open(filepath)
	if err != nil {
		panic("Couldn't read the file")
	}
	reader := bufio.NewReader(input)

	partsGrid := make([][]enginePart, 0)
	markerGrid := make([][]marker, 0)

	rowNum := 0
	for {
		lineBytes, err := reader.ReadBytes('\n')
		line := string(lineBytes)
		//	log.Printf("INFO: processing line %s", line)
		if err != nil {
			break
		}
		parts := make([]enginePart, 0)
		markers := make([]marker, 0)
		pos := 0
		var char byte = 0

		for {
			if pos, err = parser.Skip(line, pos, '.'); err != nil {
				break
			}
			if pos, err = parser.Skip(line, pos, '\n'); err != nil {
				break
			}
			if num, newPos, err := parser.ExpectNumber(line, pos); err == nil {
				parts = append(parts, enginePart{
					number: num,
					start:  pos,
					end:    newPos - 1,
					row:    rowNum,
				})
				pos = newPos
			}
			if char, pos, err = parser.Peek(line, pos); err == nil && char != '.' && char != '\n' {
				markers = append(markers, marker{
					row:        rowNum,
					position:   pos,
					markerType: char,
				})
				pos++
			}
		}
		rowNum++
		partsGrid = append(partsGrid, parts)
		markerGrid = append(markerGrid, markers)

	}

	dummyParts := make([]enginePart, 0)
	prevRowParts := &dummyParts
	puzzle1Result := 0

	var currRowParts *[]enginePart
	var nextRowParts *[]enginePart

	partsGrid = append(partsGrid, dummyParts)
	for row := 0; row < rowNum-1; row++ {
		currRowParts = &partsGrid[row]
		nextRowParts = &partsGrid[row+1]
		markers := &markerGrid[row]
		for i := range *markers {
			currMarker := &(*markers)[i]
			for _, partRow := range []*[]enginePart{prevRowParts, currRowParts, nextRowParts} {
				for i := 0; i < len(*partRow); i++ {
					part := (*partRow)[i]
					if currMarker.position >= part.start-1 && currMarker.position <= part.end+1 {
					    currMarker.nearbyParts = append(currMarker.nearbyParts, &part)
						if !part.valid {
							part.valid = true
							puzzle1Result += part.number
						}
					}
				}

			}

		}
		prevRowParts = currRowParts
	}
	log.Printf("INFO: Puzzle1 Result %d", puzzle1Result)

	puzzle2Result := int64(0)
	for r := range markerGrid {
		markers := &markerGrid[r]
		for i := range *markers {
			currMarker := &(*markers)[i]
			if currMarker.markerType == '*' && len(currMarker.nearbyParts) == 2 {
				puzzle2Result += int64((currMarker.nearbyParts[0].number) * (currMarker.nearbyParts[1].number))
			}
		}
	}

	log.Printf("INFO: Puzzle2 Result: %d", puzzle2Result)
}

/*

.....1234....... abs(mpos - start) <=1
....*...........
................
.....1234....... abs(mpos - end) <=1
.........*.....
................
.....123456..... (mpos - start) + (end - mpos) = end - start
.......*........
................

*/

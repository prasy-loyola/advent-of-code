package main

import (
	"bufio"
	"log"
	"os"
	parsermod "parser"
)

type Node struct {
	name  string
	left  string
	right string
}

func main() {
	filepath := "input.txt"
	var err error
	var input *os.File
	if input, err = os.Open(filepath); err != nil {
		log.Fatal("Couldn't read file %s", err.Error())
	}
	parser := parsermod.Parser{
		Reader: bufio.NewReader(input),
	}

	line, _ := parser.ReadNextLine()
	var instructions string
	if instructions, err = parser.ReadString('\n'); err != nil {
		log.Fatalf("ERROR: invalid format %s. %s", line, err.Error())
		return
	}

	network := make(map[string]Node)
	line, err = parser.ReadNextLine()
	for line, err = parser.ReadNextLine(); err != parsermod.EOF; line, err = parser.ReadNextLine() {
		//	log.Printf("INFO: parsing line '%s'", line)
		if err != nil && err != parsermod.EOL {
			log.Fatalf("ERROR: invalid text '%s'. %s", line, err.Error())
		}
		var name, left, right string
		if name, err = parser.ReadString(' '); err != nil {
			continue
		}
		if _, err = parser.ExpectWord(" = ("); err != nil {
			continue
		}
		if left, err = parser.ReadString(','); err != nil {
			continue
		}
		if _, err = parser.ExpectWord(", "); err != nil {
			continue
		}
		if right, err = parser.ReadString(')'); err != nil {
			continue
		}

		network[name] = Node{
			name:  name,
			left:  left,
			right: right,
		}
	}

	log.Printf("INFO: instructions %s", instructions)
	log.Printf("INFO: network %v", network)

	currNode := network["AAA"]
	i := 0
	for currNode != network["ZZZ"] {
		if instructions[i % len(instructions)] == 'L' {
			currNode = network[currNode.left]
		} else if instructions[i % len(instructions)] == 'R' {
			currNode = network[currNode.right]
		}
		i += 1
	}

    log.Printf("INFO: no. of steps taken %d", i)
}

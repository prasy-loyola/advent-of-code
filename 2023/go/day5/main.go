package main

import (
	"bufio"
	"log"
	"os"
	parsermod "parser"
	"sort"
)


type mappings [][3]int


func (m mappings) Len() int {
    return len(m)
}

func (m mappings) Less(i, j int) bool {
    return m[i][1] < m[j][1]
}

func (m mappings) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m mappings) convert(num int) int {
    sort.Sort(m)

    for _, curmapping := range m {
        destStart, sourceStart, _range := curmapping[0], curmapping[1], curmapping[2]
        if num >= sourceStart && num < sourceStart + _range {
            return destStart + (num - sourceStart)
        }
    }
    return num
}


func getNextMapping(p *parsermod.Parser) mappings {
    mappings := make(mappings, 0)
    var name, line string
    var err error
    name, err = p.ReadString(' ')
    _   , err = p.ExpectWord(" map:")
    for {
        if err != nil && err != parsermod.EOL {
            log.Fatalf("ERROR: invalid mapping %s. text '%s', \n%s", name, line, err.Error())
        }
        line, err = p.ReadNextLine()
        //log.Printf("INFO: parsing line %s", line)
        var numbers []int
        if numbers, err = p.ExpectListOfNumbers(' '); err != nil {
            continue
        } else if len(numbers) != 3 {
            break
        } else {
            mappings = append(mappings, [3]int{numbers[0], numbers[1], numbers[2]})
        }
    }
    return mappings
}

func main() {

	filepath := "input.txt"
	var input *os.File
	var err error
	if input, err = os.Open(filepath); err != nil {
		panic("Couldn't read input")
	}
	reader := bufio.NewReader(input)
	parser := parsermod.Parser{
		Reader: reader,
	}

	var seeds []int
	var line string
	line, err = parser.ReadNextLine()
	log.Printf("INFO: parsing line %s", line)
	if _, err = parser.ExpectWord("seeds:"); err != nil {
		log.Fatalf("ERROR: Invalid format %s", err.Error())
	}
	if seeds, err = parser.ExpectListOfNumbers(' '); err != nil && err != parsermod.EOL {
		log.Fatalf("ERROR: Invalid format %s", err.Error())
	}

    for {
        if err != nil && err != parsermod.EOL {
		    log.Fatalf("ERROR: line %s Invalid format %s", line, err.Error())
        }
        if line, err = parser.ReadNextLine(); err == parsermod.EOF {
            break
        }
        if line == "\n" {
            //log.Printf("INFO: starting new mapping. %s", line)
            continue
        }
        mappings := getNextMapping(&parser)

        for i := range seeds {
            seeds[i] = mappings.convert(seeds[i])
        }

        //log.Printf("INFO: mappings %v seeds %v", mappings, seeds)
    }

    puzzle1Result := seeds[0]

    for _, num := range seeds {
        if puzzle1Result > num {
            puzzle1Result = num
        }
    }

    log.Printf("INFO: puzzle1 result : %d", puzzle1Result)


}

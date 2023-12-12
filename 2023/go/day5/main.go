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
		if num >= sourceStart && num < sourceStart+_range {
			return destStart + (num - sourceStart)
		}
	}
	return num
}

func (m mappings) convertToRange(num, rng int) [][2]int {

	if rng < 1 {
		return [][2]int{}
	}

	for _, curmapping := range m {
		destStart, sourceStart, _range := curmapping[0], curmapping[1], curmapping[2]
		if num < sourceStart {
			if num+rng < sourceStart {
				break
			} else {
				result := make([][2]int, 0)
				result = append(result, [2]int{num, sourceStart - num})
				for _, newRange := range m.convertToRange(sourceStart, rng-(sourceStart-num)) {
					if newRange[1] > 0 {
						result = append(result, newRange)
					}
				}
				return result
			}

		} else if num >= sourceStart && num < sourceStart + _range {
			result := make([][2]int, 0)
			rangeEnd := min(sourceStart+_range, num+rng)
			result = append(result, [2]int{destStart + num - sourceStart, rangeEnd - num})

			if rangeEnd >= sourceStart+_range {
				for _, newRange := range m.convertToRange(rangeEnd, abs((sourceStart+_range)-(num+rng))) {
					result = append(result, newRange)
				}
			}
			return result
		}
	}
	return [][2]int{{num, rng}}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func getNextMapping(p *parsermod.Parser) mappings {
	mappings := make(mappings, 0)
	var name, line string
	var err error
	name, err = p.ReadString(' ')
    //log.Printf("INFO: mapping name %s", name)
	_, err = p.ExpectWord(" map:")
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
	//log.Printf("INFO: parsing line %s", line)
	if _, err = parser.ExpectWord("seeds:"); err != nil {
		log.Fatalf("ERROR: Invalid format %s", err.Error())
	}
	if seeds, err = parser.ExpectListOfNumbers(' '); err != nil && err != parsermod.EOL {
		log.Fatalf("ERROR: Invalid format %s", err.Error())
	}
	ranges := make([][2]int, 0)
	for i := 0; i < len(seeds)-1; i += 2 {
		ranges = append(ranges, [2]int{seeds[i], seeds[i+1]})
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

		newRanges := make([][2]int, 0)
		for _, newRange := range ranges {

			for _, r := range mappings.convertToRange(newRange[0], newRange[1]) {
				newRanges = append(newRanges, r)
			}
		}
		ranges = newRanges

	}

	puzzle1Result := seeds[0]

	for _, num := range seeds {
		if puzzle1Result > num {
			puzzle1Result = num
		}
	}

    puzzle2Result := ranges[0][0]
    for _, rng := range ranges {
        if rng[0] < puzzle2Result {
            puzzle2Result = rng[0]
        }
    }

	log.Printf("INFO: puzzle1 result : %d", puzzle1Result)
	log.Printf("INFO: puzzle2 result : %d", puzzle2Result)

}

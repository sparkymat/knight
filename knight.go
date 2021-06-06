package main

import "fmt"

func findPossibilities(boardSize uint16, position uint16, visited []uint16) []uint16 {
	visitedMap := map[uint16]interface{}{}

	for _, visitedCell := range visited {
		visitedMap[visitedCell] = struct{}{}
	}

	row := int16(position / boardSize)
	col := int16(position % boardSize)

	allPossibilities := []uint16{}

	// -2, -1
	if row-2 >= 0 && col-1 >= 0 {
		allPossibilities = append(allPossibilities, (uint16(row)-2)*boardSize+(uint16(col)-1))
	}

	// -2, +1
	if row-2 >= 0 && col+1 < int16(boardSize) {
		allPossibilities = append(allPossibilities, (uint16(row)-2)*boardSize+(uint16(col)+1))
	}

	// +2, -1
	if row+2 < int16(boardSize) && col-1 >= 0 {
		allPossibilities = append(allPossibilities, (uint16(row)+2)*boardSize+(uint16(col)-1))
	}

	// +2, +1
	if row+2 < int16(boardSize) && col+1 < int16(boardSize) {
		allPossibilities = append(allPossibilities, (uint16(row)+2)*boardSize+(uint16(col)+1))
	}

	// -1, -2
	if row-1 >= 0 && col-2 >= 0 {
		allPossibilities = append(allPossibilities, (uint16(row)-1)*boardSize+(uint16(col)-2))
	}

	// -1, +2
	if row-1 >= 0 && col+2 < int16(boardSize) {
		allPossibilities = append(allPossibilities, (uint16(row)-1)*boardSize+(uint16(col)+2))
	}

	// +1, -2
	if row+1 < int16(boardSize) && col-2 >= 0 {
		allPossibilities = append(allPossibilities, (uint16(row)+1)*boardSize+(uint16(col)-2))
	}

	// +1, +2
	if row+1 < int16(boardSize) && col+2 < int16(boardSize) {
		allPossibilities = append(allPossibilities, (uint16(row)+1)*boardSize+(uint16(col)+2))
	}

	filteredPossibilities := []uint16{}
	for _, possiblity := range allPossibilities {
		if _, visited := visitedMap[possiblity]; !visited {
			filteredPossibilities = append(filteredPossibilities, possiblity)
		}
	}

	return filteredPossibilities
}

func main() {
	const boardSize = uint16(8)
	startPos := uint16(0)

	positions := [boardSize * boardSize]uint16{0}
	positionsCursor := 0

	// Place start
	positions[positionsCursor] = startPos

	for positionsCursor < len(positions) {
		position := positions[positionsCursor]
		possibilities := findPossibilities(boardSize, position, positions[0:positionsCursor])
		fmt.Printf("possibilities at (%d,%d)=%+v\n", position/boardSize, position%boardSize, possibilities)
		if len(possibilities) == 0 {
			panic(fmt.Sprintf("no possibilies at (%d,%d) for boardSize=%d after %+v", position/boardSize, position%boardSize, boardSize, positions))
		}
		positionsCursor++
		if positionsCursor < len(positions) {
			positions[positionsCursor] = possibilities[0]
		}
	}

	fmt.Printf("%+v\n", positions)
}

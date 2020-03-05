package dominoes

import "fmt"

type Domino [2]int

func MakeChain(dominoes []Domino) ([]Domino, bool) {

	fmt.Println("beginning - ", dominoes)

	inputLength := len(dominoes)
	if inputLength == 0 {
		return dominoes, true
	}
	if inputLength == 1 {
		if dominoes[0][0] == dominoes[0][1] {
			return dominoes, true
		} else {
			return dominoes, false
		}
	}
	resultMap := make(map[int]int)
	for i := range dominoes {
		resultMap[i+1] = 0
	}

	for _, value := range dominoes {
		resultMap[value[0]] = resultMap[value[0]] + 1
		resultMap[value[1]] = resultMap[value[1]] + 1
	}

	checkForMoreOccurence := false

	for _, value := range resultMap {
		if value%2 != 0 {
			return nil, false
		}
		if value != 2 {
			checkForMoreOccurence = true
		}
	}

	resultDominoes := make([]Domino, 1)
	resultDominoes[0] = dominoes[0]
	dominoes = RemoveDomino(dominoes, 0)
	currentNumberToFind := resultDominoes[0][1]
	fmt.Println("current result - ", resultDominoes, "input - ", dominoes, "to find - ", currentNumberToFind, "check for more occurencies - ", checkForMoreOccurence)

	for i := 1; i < inputLength; i++ {
		var removed Domino
		dominoes, removed, currentNumberToFind = FindAndRemove(dominoes, currentNumberToFind, checkForMoreOccurence)
		if currentNumberToFind == -1 {
			return nil, false
		}
		resultDominoes = append(resultDominoes, removed)
		fmt.Println("iteration - ", i, "len - ", "current result - ", resultDominoes, "input - ", dominoes, "to find - ", currentNumberToFind)
	}
	fmt.Println("\n", "input - ", dominoes)
	return resultDominoes, true
}

func RemoveDomino(dominoes []Domino, index int) []Domino {
	if index == 0 && len(dominoes) == 1 || len(dominoes) == 0 {
		return []Domino{}
	}
	if index == 0 {
		return dominoes[1:]
	}
	if index == len(dominoes)-1 {
		return dominoes[0 : len(dominoes)-1]
	}
	return append(dominoes[0:index], dominoes[index+1:]...)
}

func FindAndRemove(dominoes []Domino, numberToFind int, checkForMoreOccurence bool) ([]Domino, Domino, int) {
	var newNumberToFind = -1
	var resultDominoes []Domino
	var removedDomino Domino

	if checkForMoreOccurence {
		for i, value := range dominoes {
			if value[0] == numberToFind && value[1] == numberToFind {
				resultDominoes = RemoveDomino(dominoes, i)
				return resultDominoes, value, numberToFind
			}
		}
	}

	for i, value := range dominoes {
		if value[0] == numberToFind || value[1] == numberToFind {
			resultDominoes = RemoveDomino(dominoes, i)
			if value[0] == numberToFind {
				newNumberToFind = value[1]
				removedDomino = value
			} else {
				newNumberToFind = value[0]
				removedDomino = Flip(value)
			}
			break
		}
	}
	return resultDominoes, removedDomino, newNumberToFind
}

func Flip(domino Domino) Domino {
	return Domino{domino[1], domino[0]}
}

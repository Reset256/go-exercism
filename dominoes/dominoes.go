package dominoes

import "fmt"

type Domino [2]int

func MakeChain(dominoes []Domino) ([]Domino, bool) {

	if len(dominoes) == 0 {
		return dominoes, true
	}
	if  len(dominoes) == 1 {
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

	fmt.Println(dominoes, resultMap)

	for _, value := range resultMap {
		if value != 2 {
			return nil, false
		}
	}


	resultDominoes := make([]Domino, len(dominoes))
	resultDominoes[0] = dominoes[0]
	dominoes = removeDomino(dominoes, 0)
	currentNumberToFind := resultDominoes[0][1]

	for i := 1; i < len(dominoes); i++ {
		var removed Domino
		dominoes, removed, currentNumberToFind = findAndRemove(dominoes, currentNumberToFind)
		if currentNumberToFind == -1 {
			return nil, false
		}
		resultDominoes = append(resultDominoes, removed)
	}
	return resultDominoes, true
}


func removeDomino(dominoes []Domino, index int) []Domino {
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

func findAndRemove(dominoes []Domino, numberToFind int) ([]Domino, Domino, int) {
	var newNumberToFind = -1
	var resultDominoes []Domino
	var removedDomino Domino

	for i, value := range dominoes {
		if value[0] == numberToFind || value[1] == numberToFind {
			if value[0] == numberToFind {
				newNumberToFind = value[1]
			} else {
				newNumberToFind = value[0]
			}
			resultDominoes = removeDomino(dominoes, i)
			removedDomino = value
			break
		}
	}
	return resultDominoes, removedDomino, newNumberToFind
}

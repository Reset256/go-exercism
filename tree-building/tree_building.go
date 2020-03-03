package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	fmt.Println("Input - ", records)
	if len(records) < 1 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for position, record := range records {
		if position > 0 && record.ID-records[position-1].ID > 1 {
			return nil, errors.New("Non-continious")
		}
	}
	fmt.Println("After first sorting - ", records)

	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent == records[j].Parent {
			return records[i].ID < records[j].ID
		}
		return records[i].Parent < records[j].Parent
	})
	fmt.Println("After second sorting - ", records)

	if records[0].Parent != records[0].ID {
		return nil, errors.New("Empty records")
	}
	baseNode := Node{
		ID: records[0].ID,
	}

	for i := 1; i < len(records); i++ {
		value := records[i]
		if getNodeById(&baseNode, value.ID) == nil {
			node := Node{
				ID: value.ID,
			}
			parentNode := getNodeById(&baseNode, value.Parent)
			if parentNode == nil {
				return nil, errors.New("Non-continious")
			} else if parentNode.ID > value.ID{
				return nil, errors.New("higher id parent of lower id")
			}
			parentNode.Children = append(parentNode.Children, &node)
		} else {
			return nil, errors.New("Duplicate node")
		}
	}

	return &baseNode, nil
}

func getNodeById(rootNode *Node, id int) *Node {
	if rootNode.ID == id {
		return rootNode
	} else if rootNode.Children == nil || len(rootNode.Children) == 0 {
		return nil
	} else {
		for _, node := range rootNode.Children {
			noteById := getNodeById(node, id)
			if noteById != nil {
				return noteById
			}
		}
	}
	return nil
}

/*
func getById(rootNode *Node, id int) *Node {
	currentNode := rootNode
	if currentNode.ID != id {
		for {
			if currentNode.ID != id {
				for _, child := range rootNode.Children {
					byId := getById(child, id)
					if byId != nil {
						return byId
					}
				}
			}
		}
	} else {
		return currentNode
	}
	return nil
}

func Build(records []Record) (*Node, error) {
	if len(records) < 1 {
		return nil, nil
	}
	if records[0].Parent != records[0].ID {
		return nil, errors.New("Empty records")
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID > records[j].ID
	})
	rootNode := Node{
		ID:       records[0].ID,
		Children: []*Node{},
	}
	for i := 1; i < len(records); i++ {
		n := Node{
			ID:       records[i].ID,
			Children: []*Node{},
		}
		node := getById(&rootNode, records[i].Parent)
		node.Children = append(node.Children, &n)
	}

	return &rootNode, nil
}
*/

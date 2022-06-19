package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	node := make(map[int]*Node, len(records))

	fmt.Println(records)

	for i, r := range records {
		// verify that record is valid
		if r.ID != i || r.Parent > r.ID || (r.ID != 0 && r.ID == r.Parent) {
			return nil, errors.New("not in sequence or has bad parent")
		}

		node[r.ID] = &Node{ID: r.ID}
		// if not root node, add children node to parent
		if r.ID != 0 {
			node[r.Parent].Children = append(node[r.Parent].Children, node[r.ID])
		}
	}

	return node[0], nil
}

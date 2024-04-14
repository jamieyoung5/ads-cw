package dlx

import "fmt"

type Matrix struct {
	Root *Node
}

/*
Node
The structure of the exact covert matrix consists of nodes (quadruply linked circular list), all of which also link to a column and represent a 1 in the matrix.
A column contains info about its respective id and vertical size (not including header node).
The header node, also referred to as a 'control node' in the context of dlx serves as a root point for a column in the matrix and facilitates linking/unlinking (covering and uncovering) a columns data nodes (the 1's)
*/
type Node struct {
	Left, Right, Up, Down *Node
	Column                *Column
}

// NewMatrix generates columns and each columns' respective header node
func NewMatrix(columnIDs []string) *Matrix {
	root := &Column{Node: &Node{}}
	root.Left = root.Node
	root.Right = root.Node

	lastHeader := root
	for _, id := range columnIDs {
		newHeader := &Column{ID: id}
		newHeader.Node = &Node{Left: lastHeader.Node, Right: root.Node, Column: newHeader}
		newHeader.Up = newHeader.Node
		newHeader.Down = newHeader.Node

		lastHeader.Right = newHeader.Node
		root.Left = newHeader.Node
		lastHeader = newHeader
	}

	return &Matrix{Root: root.Node}
}

func (m *Matrix) PrintMatrix() {
	currentHeader := m.Root.Right // Start from the first actual header, not the root itself
	for currentHeader != m.Root {
		fmt.Printf("Column ID: %s, Size: %d\n", currentHeader.Column.ID, currentHeader.Column.Size)
		currentNode := currentHeader.Down
		for currentNode != currentHeader { // Iterate until we loop back to the dummy node
			fmt.Println("  Node at", currentNode)
			currentNode = currentNode.Down
		}
		currentHeader = currentHeader.Right
	}
}

// AppendRow adds a new node to the bottom of every column included in columnIds and links them together to form a row
func (m *Matrix) AppendRow(columnIds []string) {
	/*
		The obvious solution here is to search through the headers and for each search through column ids to find
		the specific headers that we want or vice versa, however it is more efficient to put the id's into a map which
		allows O(1) average-time complexity lookups and then iterate over the headers once.

		The latter option has a worst-case complexity of O(n×m) whereas the former has a complexity of O(n+m) — O(m)O(m)
		for creating the map and then a complexity of O(n) for traversing through the headers.
	*/
	idSet := make(map[string]struct{})
	for _, id := range columnIds {
		idSet[id] = struct{}{}
	}

	var firstNode, lastNode *Node

	header := m.Root.Right
	for header != m.Root {
		if _, exists := idSet[header.Column.ID]; exists {
			newNode := &Node{
				Column: header.Column,
				Up:     header.Up,
				Down:   header,
			}

			// Link vertically
			newNode.Up.Down = newNode
			newNode.Down.Up = newNode
			header.Column.Size++

			// Link horizontally
			if firstNode == nil {
				firstNode = newNode
				lastNode = newNode
			} else {
				newNode.Left = lastNode
				lastNode.Right = newNode
				lastNode = newNode
			}
		}
		header = header.Right
	}

	// Close the horizontal loop
	if firstNode != nil {
		if lastNode == firstNode {
			// If there is only one node, it should point to itself
			firstNode.Left = firstNode
			firstNode.Right = firstNode
		} else {
			firstNode.Left = lastNode
			lastNode.Right = firstNode
		}
	}
}

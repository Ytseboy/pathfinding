package main

import (
	"fmt"
	"math"
)

type PriorityQueue []*Node

func (pq *PriorityQueue) Len() int {
	panic("not implemented")
}

func (pq *PriorityQueue) Less(i int, j int) bool {
	panic("not implemented")
}

func (pq *PriorityQueue) Swap(i int, j int) {
	panic("not implemented")
}

func (pq *PriorityQueue) Push(x interface{}) {
	panic("not implemented")
}

func (pq *PriorityQueue) Pop() interface{} {
	panic("not implemented")
}

type Node struct {
	X      int
	Y      int
	Parent *Node
	H      float64
	Index  int
}

var (
	dirs = map[string][2]int{
		"north": [2]int{1, 0},
		"south": [2]int{0, -1},
		"east":  [2]int{0, 1},
		"west":  [2]int{0, -1},
	}
)

func (n *Node) Neighbours() (neighbours [4]Node) {
	i := 0
	for _, dir := range dirs {
		neighbours[i] = Node{
			X: n.X + dir[0],
			Y: n.Y + dir[1],
		}
		i++
	}
	return
}

func (n *Node) Eq(comp *Node) (isEq bool) {
	if n.X == comp.X && n.Y == comp.Y {
		return true
	}

	return
}

func ManhatanDistance(start, destination Node) (H float64) {
	H = math.Abs(float64(start.X-destination.X)) + math.Abs(float64(start.Y-destination.Y))
	return
}

func main() {
	var (
		openSet   []Node
		closedSet []Node
		startNode Node
	)

	startNode = Node{
		X: 37,
		Y: 17,
	}

	destinationNode := Node{
		X: 48,
		Y: 17,
	}

	currentNode := startNode

	for _, node := range startNode.Neighbours() {
		node.Parent = &currentNode
		node.H = ManhatanDistance(startNode, destinationNode)
		openSet = append(openSet, node)
	}
	fmt.Println("hello world")
}

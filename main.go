package main

import (
	"fmt"
	"math"
)

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].H > pq[j].H
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// func (pq *PriorityQueue) update(item *Node, value string, priority int) {
// 	item.value = value
// 	item.H = priority
// 	heap.Fix(pq, item.index)
// }

type Node struct {
	X      int
	Y      int
	Parent *Node
	H      float64
	index  int
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
		openSet   PriorityQueue
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
		openSet.Push(node)
	}
	fmt.Println("hello world")
}

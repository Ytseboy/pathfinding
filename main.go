package main

import "fmt"

//import "os"
import "container/heap"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].H < pq[j].H
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

type Node struct {
	X      int
	Y      int
	Parent *Node
	H      float64
	index  int
	Dir    string
}

var (
	dirs = map[string][2]int{
		"E":  [2]int{1, 0},
		"W":  [2]int{-1, 0},
		"S":  [2]int{0, 1},
		"SE": [2]int{1, 1},
		"SW": [2]int{-1, 1},
		"N":  [2]int{0, -1},
		"NE": [2]int{1, -1},
		"NW": [2]int{-1, -1},
	}
)

func (n *Node) Neighbours() (neighbours [8]Node) {
	i := 0
	for key, dir := range dirs {
		neighbours[i] = Node{
			X:   n.X + dir[0],
			Y:   n.Y + dir[1],
			Dir: key,
		}
		i++
	}
	return
}

func (n *Node) Eq(comp *Node) (isEq bool) {
	if n.X == comp.X && n.Y == comp.Y {
		isEq = true
		return
	}

	return
}

func ManhatanDistance(start, destination Node) (H float64) {
	H = math.Abs(float64(start.X-destination.X)) + math.Abs(float64(start.Y-destination.Y))
	return
}

func PointDistance(start, destination Node) (H float64) {
	H = math.Sqrt(math.Pow(float64(start.X-destination.X), 2) + math.Pow(float64(start.Y-destination.Y), 2))
	return
}

func In(nodes []*Node, toFind Node) (found bool) {
	for _, node := range nodes {
		if node.Eq(&toFind) {
			found = true
			return
		}
	}
	return
}

func main() {
	// lightX: the X position of the light of power
	// lightY: the Y position of the light of power
	// initialTX: Thor's starting X position
	// initialTY: Thor's starting Y position
	var lightX, lightY, initialTX, initialTY int
	fmt.Scan(&lightX, &lightY, &initialTX, &initialTY)
	var (
		openSet   PriorityQueue
		closedSet []*Node
		startNode *Node
	)
	startNode = &Node{
		X: initialTX,
		Y: initialTY,
	}

	destinationNode := Node{
		X: lightX,
		Y: lightY,
	}

	currentNode := startNode
	heap.Init(&openSet)

	for {
		// remainingTurns: The remaining amount of turns Thor can move. Do not remove this line.
		var remainingTurns int
		fmt.Scan(&remainingTurns)
		for !currentNode.Eq(&destinationNode) {
			neighbours := currentNode.Neighbours()
			for i := range neighbours {
				if In(closedSet, neighbours[i]) {
					continue
				} else {
					neighbours[i].Parent = currentNode
					if !In(openSet, neighbours[i]) {
						neighbours[i].H = PointDistance(neighbours[i], destinationNode)
						fmt.Fprintf(os.Stderr, "Pushing {node: %v x -> %d; y -> %d; h -> %.2f} into heap\n", neighbours[i], neighbours[i].X, neighbours[i].Y, neighbours[i].H)
						heap.Push(&openSet, &neighbours[i])
					}
				}
			}

			if openSet.Len() <= 0 {
				break
			}

			currentNode = heap.Pop(&openSet).(*Node)
			closedSet = append(closedSet, currentNode)
			fmt.Fprintf(os.Stderr, "Moving to {node: %v x -> %d; y -> %d; h -> %.2f} tile\n", currentNode, currentNode.X, currentNode.Y, currentNode.H)
		}

		var path []*Node
		if currentNode.Eq(&destinationNode) {
			n := closedSet[len(closedSet)-1]
			for n != nil {
				path = append(path, n)
				n = n.Parent
			}
		} else {
			fmt.Fprintf(os.Stderr, "path unsolvable")
			os.Exit(1)
		}
		for i, node := range path {
			fmt.Fprintf(os.Stderr, "%d tile -> x: %d; y: %d; dir -> %s\n", i, node.X, node.Y, node.Dir)
			fmt.Println(node.Dir)
		}
		// fmt.Fprintln(os.Stderr, "Debug messages...")

		// A single line providing the move to be made: N NE E SE S SW W or NW
	}
}

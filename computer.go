package main

import (
	"container/heap"
	"math"
	"slices"
)

type Node struct {
	pos     *Coordinate
	g, h, f float64
	parent  *Node
}

func newNode(pos *Coordinate, g, h float64) *Node {
	return &Node{
		pos: pos,
		g:   g,
		h:   h,
		f:   g + h,
	}
}

func euclideanDistance(pos1, pos2 *Coordinate) float64 {
	x := float64(pos2.x - pos1.x)
	y := float64(pos2.y - pos1.y)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func getValidNeighbours(grid [][]int, pos *Coordinate) []*Coordinate {
	x := pos.x
	y := pos.y

	possibleMoves := []*Coordinate{
		&Coordinate{x + 1, y},
		&Coordinate{x - 1, y},
		&Coordinate{x, y + 1},
		&Coordinate{x, y - 1},
	}

	validMoves := *new([]*Coordinate)
	for _, move := range possibleMoves {
		nx := move.x
		ny := move.y
		if nx > 0 || ny > 0 || nx < GRID_SIZE || ny < GRID_SIZE || grid[ny][nx] == 0 {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func reconstructPath(goalNode *Node) []*Coordinate {
	path := *new([]*Coordinate)
	current := goalNode

	for current != nil {
		path = append(path, current.pos)
		current = current.parent
	}

	slices.Reverse(path)
	return path
}

func FindPath(grid [][]int, startPos, goalPos *Coordinate) []*Coordinate {
	start := newNode(startPos, 0, euclideanDistance(startPos, goalPos))

	openList := make(PriorityQueue, 0)
	openList.Push(&Priority{priority: start.f, pos: startPos})
	heap.Init(&openList)

	openMap := make(map[*Coordinate]*Node)
	openMap[startPos] = start
	closedSet := make(map[*Coordinate]bool)

	for len(openList) != 0 {
		current := heap.Pop(&openList).(*Priority)
		openList = openList[1:]
		currentNode := openMap[current.pos]

		if current.pos == goalPos {
			return reconstructPath(currentNode)
		}

		closedSet[current.pos] = true

		for _, neighbourPos := range getValidNeighbours(grid, current.pos) {
			if closedSet[neighbourPos] {
				continue
			}

			newG := currentNode.g + euclideanDistance(current.pos, neighbourPos)

			if openMap[neighbourPos] == nil {
				neighbour := newNode(neighbourPos, newG, euclideanDistance(neighbourPos, goalPos))
				neighbour.parent = currentNode
				heap.Push(&openList, Priority{priority: neighbour.f, pos: neighbourPos})
				openMap[neighbourPos] = neighbour
			} else if newG < openMap[neighbourPos].g {
				neighbour := openMap[neighbourPos]
				neighbour.g = newG
				neighbour.f = newG + neighbour.h
				neighbour.parent = currentNode
			}
		}
	}

	return *new([]*Coordinate)
}

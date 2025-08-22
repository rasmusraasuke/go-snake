package main

import "math"

type Priority struct {
	rate  float64
	value *Coordinate
}

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

func euclideanDistance(pos1, pos2 Coordinate) float64 {
	x := float64(pos2.x - pos1.x)
	y := float64(pos2.y - pos1.y)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func getValidNeighbours(grid [][]int, pos Coordinate) []Coordinate {
	x := pos.x
	y := pos.y

	possibleMoves := []Coordinate{
		Coordinate{x + 1, y},
		Coordinate{x - 1, y},
		Coordinate{x, y + 1},
		Coordinate{x, y - 1},
	}

	validMoves := *new([]Coordinate)
	for _, move := range possibleMoves {
		nx := move.x
		ny := move.y
		if nx > 0 || ny > 0 || nx < GRID_SIZE || ny < GRID_SIZE || grid[ny][nx] == 0 {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

//func FindPath(grid [][]int, startPos, goalPos *Coordinate) {
//	start := newNode(startPos, 0, euclideanDistance(*startPos, *goalPos))
//
//	openList := []*Priority{&Priority{start.f, start.pos}}
//	openMap := make(map[*Coordinate]*Node)
//	openMap[startPos] = start
//
//	closedList := make(map[*Node]bool)
//
//}

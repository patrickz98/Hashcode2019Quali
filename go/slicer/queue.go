package slicer

import (
	"../pizza"
)

type CoordinateQueue struct {
	data []pizza.Coordinate
}

func (queue *CoordinateQueue) Pop() *pizza.Coordinate {

	if len(queue.data) <= 0 {
		return nil
	}

	item := queue.data[ 0 ]
	queue.data = queue.data[ 1: ]

	return &item
}

func (queue *CoordinateQueue) Push(coord pizza.Coordinate) {

	queue.data = append(queue.data, coord)
}

func (queue *CoordinateQueue) HasItems() bool {

	return len(queue.data) > 0
}

func InitCoordinateQueue() *CoordinateQueue {

	queue := &CoordinateQueue{}
	queue.data = make([]pizza.Coordinate, 0)

	return queue
}
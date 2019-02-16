package slicer

import "fmt"

func (slicer *Slicer) ExpandThroughMove() {

	fmt.Println("Expand through move")

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		cell := slicer.Pizza.Cells[ xy ]

		if cell.Slice != nil {
			continue
		}

		queue.Push(xy)
	}

	for queue.HasItems() {
		fmt.Printf("Destruction queue --> %-7d\r", len(queue.data) - 1)
		slicer.tryExpand(queue)
	}

	fmt.Printf("Destruction queue --> done\n")
}

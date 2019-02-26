package slicer

import (
	"../pizza"
	"../set"
	"../simple"
	"fmt"
)

type Hole struct {
	*pizza.Slice
}

func (hole Hole) CountNilCells() (count int) {

	for _, xy := range hole.Traversal() {
		cell := hole.Pizza.Cells[xy]

		if cell.Slice == nil {
			count++
		}
	}

	return count
}

func (slicer *Slicer) findHoleAt(xy pizza.Coordinate) *Hole {

	if slicer.Pizza.HasSliceAt(xy) {
		return nil
	}

	piz := slicer.Pizza
	row := pizza.Vector{Start: xy.Row, End: xy.Row}
	col := pizza.Vector{Start: xy.Column, End: xy.Column}

	holeSlice := &pizza.Slice{Pizza: piz, Row: row, Column: col}
	hole := &Hole{holeSlice}

	nilCount := hole.CountNilCells()

	// Expand left
	for {

		left := simple.Max(piz.Column.Start, col.Start - 1)
		col = pizza.Vector{Start: left, End: col.End}

		holeSlice.Column = col
		hole = &Hole{holeSlice}

		count := hole.CountNilCells()

		if nilCount < count {
			nilCount = count
		} else {
			break
		}
	}

	// Expand top
	for {

		top := simple.Max(piz.Row.Start, row.Start - 1)
		row = pizza.Vector{Start: top, End: row.End}

		holeSlice.Row = row
		hole = &Hole{holeSlice}

		count := hole.CountNilCells()

		if nilCount < count {
			nilCount = count
		} else {
			break
		}
	}

	// Expand right
	for {

		right := simple.Min(piz.Column.End, col.End + 1)
		col = pizza.Vector{Start: col.Start, End: right}

		holeSlice.Column = col
		hole = &Hole{holeSlice}

		count := hole.CountNilCells()

		if nilCount < count {
			nilCount = count
		} else {
			break
		}
	}

	// Expand bottom
	for {

		bottom := simple.Min(piz.Row.End, row.End + 1)
		row = pizza.Vector{Start: row.Start, End: bottom}

		holeSlice.Row = row
		hole = &Hole{holeSlice}

		count := hole.CountNilCells()

		if nilCount < count {
			nilCount = count
		} else {
			break
		}
	}

	//fmt.Println(nilCount)
	//hole.Print()

	//simple.Exit()

	return hole
}

func (slicer *Slicer) findAllHoles() []*Hole {
	holeSet := set.New(&Hole{})

	for _, xy := range slicer.Pizza.TraversalNotSlicedCells() {

		hole := slicer.findHoleAt(xy)

		if hole == nil {
			continue
		}

		holeSet.Insert(hole)
	}

	holes := make([]*Hole, 0)

	holeSet.Do(func(val interface{}) {

		hole := val.(*Hole)

		if hole.Slice == nil {
			return
		}

		holes = append(holes, hole)
	})

	return holes
}

func (slicer *Slicer) DestructShrinkHoles() {

	holes := slicer.findAllHoles()
	holesLen := len(holes)

	for inx, hole := range holes {

		overlaps := slicer.overlapSlices(hole.Slice)
		queuePrep := set.New()

		for _, xy := range hole.Traversal() {
			queuePrep.Insert(xy)
		}

		for _, overlaps := range overlaps {

			for _, xy := range overlaps.Traversal() {
				queuePrep.Insert(xy)
			}
		}

		queue := InitCoordinateQueue()

		queuePrep.Do(func(val interface{}) {

			xy := val.(pizza.Coordinate)
			queue.Push(xy)
		})

		slicer.tryDestuctShinkWithQueue(queue)

		fmt.Printf("--> holes %d/%d\r", holesLen, inx)
	}

	fmt.Println()
}

func (slicer *Slicer) ShakeHoles() {

	holes := slicer.findAllHoles()
	holesLen := len(holes)

	for inx, hole := range holes {

		overlaps := slicer.overlapSlices(hole.Slice)
		queuePrep := set.New()

		for _, xy := range hole.Traversal() {
			queuePrep.Insert(xy)
		}

		for _, overlaps := range overlaps {

			for _, xy := range overlaps.Traversal() {
				queuePrep.Insert(xy)
			}
		}

		queue := InitCoordinateQueue()

		queuePrep.Do(func(val interface{}) {

			xy := val.(pizza.Coordinate)
			queue.Push(xy)
		})

		slicer.ShakeSlicesWithQueue(queue)

		fmt.Printf("--> shake %d/%d\r", holesLen, inx)
	}
}
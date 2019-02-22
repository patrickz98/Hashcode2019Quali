package slicer

import (
	"../pizza"
	"fmt"
	// "os"
)

func (slicer *Slicer) setCorners() {
	fmt.Printf("Setting Corners with depth: %d\n", slicer.Params.CornerDepth)

	fmt.Println("Top, Left")
	xy := pizza.Coordinate{0, 0}
	slicer.findBestSliceForCoordinatesRecursive(xy, 1, 0, slicer.Params.CornerDepth)
	slicer.findBestSliceForCoordinatesRecursive(xy, 0, 1, slicer.Params.CornerDepth)

	fmt.Println("Top, Right")
	xy = pizza.Coordinate{0, slicer.Pizza.Column.End}
	slicer.findBestSliceForCoordinatesRecursive(xy, 1, 0, slicer.Params.CornerDepth)
	slicer.findBestSliceForCoordinatesRecursive(xy, 0, -1, slicer.Params.CornerDepth)

	fmt.Println("Bottom, Left")
	xy = pizza.Coordinate{slicer.Pizza.Row.End, 0}
	slicer.findBestSliceForCoordinatesRecursive(xy, -1, 0, slicer.Params.CornerDepth)
	slicer.findBestSliceForCoordinatesRecursive(xy, 0, 1, slicer.Params.CornerDepth)

	fmt.Println("Bottom, Right")
	xy = pizza.Coordinate{slicer.Pizza.Row.End, slicer.Pizza.Column.End}
	slicer.findBestSliceForCoordinatesRecursive(xy, -1, 0, slicer.Params.CornerDepth)
	slicer.findBestSliceForCoordinatesRecursive(xy, 0, -1, slicer.Params.CornerDepth)
}

func (slicer *Slicer) findBestSliceForCoordinatesRecursive(xy pizza.Coordinate, rowDir int, columnDir int, depth int) {

	if (xy.Row > depth && slicer.Pizza.Row.End-xy.Row > depth) ||
		(xy.Column > depth && slicer.Pizza.Column.End-xy.Column > depth) ||
		xy.Row < 0 || xy.Row > slicer.Pizza.Row.End ||
		xy.Column < 0 || xy.Column > slicer.Pizza.Column.End {
		fmt.Printf("Stopping: %d, rowDir: %d, columnDir: %d, Row %d, Collumn %d\n", depth, rowDir, columnDir, xy.Row, xy.Column)
		return
	}

	possibleSlices := make([]*pizza.Slice, 0)

	for _, sl := range slicer.SliceCache[xy] {
		if sl.Row.Start == xy.Row && sl.Column.Start == xy.Column {
			possibleSlices = append(possibleSlices, sl)
		}
	}

	if slicer.Pizza.Cells[xy].Slice != nil {
		if rowDir == 1 {
			xy.Row = slicer.Pizza.Cells[xy].Slice.Row.End + 1
		} else if rowDir == -1 {
			xy.Row = slicer.Pizza.Cells[xy].Slice.Row.Start - 1
		}

		if columnDir == 1 {
			xy.Column = slicer.Pizza.Cells[xy].Slice.Column.End + 1
		} else if columnDir == -1 {
			xy.Column = slicer.Pizza.Cells[xy].Slice.Column.Start - 1
		}

		slicer.findBestSliceForCoordinatesRecursive(xy, rowDir, columnDir, depth)
		return
	}

	lastSlice := slicer.getLastSlice(xy, rowDir, columnDir, slicer.Params.CornerLastSliceDepth)
	addedSlice := (*pizza.Slice)(nil)

	if len(possibleSlices) != 0 {
		if lastSlice != nil {
			for _, sl := range possibleSlices {
				if xy.Column != 0 {
					if sl.Row == lastSlice.Row {
						if !slicer.Pizza.SafeAddSlice(sl) {
							fmt.Printf("Stopping: %d, rowDir: %d, columnDir: %d, Row %d, Collumn %d\n", depth, rowDir, columnDir, xy.Row, xy.Column)
							return
						}
						addedSlice = sl
						break
					}
				} else if rowDir != 0 {
					if sl.Column == lastSlice.Column {
						if !slicer.Pizza.SafeAddSlice(sl) {
							fmt.Printf("Stopping: %d, rowDir: %d, columnDir: %d, Row %d, Collumn %d\n", depth, rowDir, columnDir, xy.Row, xy.Column)
							return
						}
						addedSlice = sl
						break
					}
				} else {
					panic("Run Time Error when setting corners!")
				}
			}
		}

		if addedSlice == nil {
			addedSlice = possibleSlices[0]
			for _, sl := range possibleSlices {
				if addedSlice.Size() < sl.Size() {
					addedSlice = sl
				}
			}

			if !slicer.Pizza.SafeAddSlice(addedSlice) {
				fmt.Printf("Stopping: %d, rowDir: %d, columnDir: %d, Row %d, Collumn %d\n", depth, rowDir, columnDir, xy.Row, xy.Column)
				return
			}
		}
	}

	xy.Row += rowDir
	xy.Column += columnDir

	slicer.findBestSliceForCoordinatesRecursive(xy, rowDir, columnDir, depth)

}

func (slicer *Slicer) getLastSlice(xy pizza.Coordinate, rowDir int, columnDir int, goBackDepth int) *pizza.Slice {

	for i := 0; i < goBackDepth; i++ {
		if xy.Row < 0 || xy.Row >= slicer.Pizza.Row.End ||
			xy.Column < 0 || xy.Column >= slicer.Pizza.Column.End {
			return nil
		}

		xy.Row += rowDir
		xy.Column += columnDir

		if slicer.Pizza.Cells[xy].Slice != nil {
			return slicer.Pizza.Cells[xy].Slice
		}
	}

	return nil
}

func (slicer *Slicer) ExpandThroughCorners() {

	paramSlices := make([]SlicerParams, 0)

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 0/3", 0, 3,
	})

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 50/3", 50, 3,
	})

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 100/3", 100, 3,
	})

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 150/3", 150, 3,
	})

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 500/3", 500, 3,
	})

	paramSlices = append(paramSlices, SlicerParams{
		"Corner 1000/3", 500, 3,
	})

	scorePercent := make([]float32, 0)

	for _, param := range paramSlices {
		fmt.Println("\nNow trying: " + param.Name)

		slicer.Pizza.RemoveAllSlice()
		slicer.Params = param
		slicer.setCorners()

		//slicer.FindBiggestParts()
		//slicer.FindSingles()

		//slicer.ExpandThroughEdge()
		slicer.FindSmallestParts()
		slicer.ExpandThroughDestruction()
		slicer.ExpandThroughShrink()

		_, addToScore := slicer.Pizza.Score()
		scorePercent = append(scorePercent, addToScore*100)
		slicer.Pizza.PrintScore()
	}

	fmt.Println()

	for i, param := range paramSlices {
		fmt.Printf(param.Name + ": %f\n", scorePercent[i])
	}
}
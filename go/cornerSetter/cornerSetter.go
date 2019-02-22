package cornerSetter

import (
	"../pizza"
	"../slicer"
	"fmt"
	// "os"
)

type CornerSetter struct {
	Slicer *slicer.Slicer
	Pizza  *pizza.Pizza
	Params CornerParams
}

type CornerParams struct {
	Name                   string
	CornerDepth            int
	CornerAcceptableOffset int
	CornerToMiddle         int
	CornerLastSliceDepth   int
	CornerMinScore         float32
}

func (cornerSetter *CornerSetter) setCorners() {
	fmt.Printf("Setting Corners with depth: %d\n", cornerSetter.Params.CornerDepth)

	xy := pizza.Coordinate{0, 0}
	for i := 0; i <= cornerSetter.Params.CornerToMiddle; i++ {
		cornerSetter.setCornersWithOffset(xy)
		xy.Row = i
		xy.Column = i
	}
}

func (cornerSetter *CornerSetter) setCornersWithOffset(offset pizza.Coordinate) {

	xy := pizza.Coordinate{0, 0}

	xy.Row += offset.Row
	xy.Column += offset.Column

	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 1, 0, true, true, cornerSetter.Params.CornerDepth)
	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 0, 1, true, true, cornerSetter.Params.CornerDepth)

	xy = pizza.Coordinate{0, cornerSetter.Pizza.Column.End}

	xy.Row += offset.Row
	xy.Column -= offset.Column

	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 1, 0, true, false, cornerSetter.Params.CornerDepth)
	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 0, -1, true, false, cornerSetter.Params.CornerDepth)

	xy = pizza.Coordinate{cornerSetter.Pizza.Row.End, 0}

	xy.Row -= offset.Row
	xy.Column += offset.Column

	cornerSetter.findBestSliceForCoordinatesRecursive(xy, -1, 0, false, true, cornerSetter.Params.CornerDepth)
	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 0, 1, false, true, cornerSetter.Params.CornerDepth)

	xy = pizza.Coordinate{cornerSetter.Pizza.Row.End, cornerSetter.Pizza.Column.End}

	xy.Row -= offset.Row
	xy.Column -= offset.Column

	cornerSetter.findBestSliceForCoordinatesRecursive(xy, -1, 0, false, false, cornerSetter.Params.CornerDepth)
	cornerSetter.findBestSliceForCoordinatesRecursive(xy, 0, -1, false, false, cornerSetter.Params.CornerDepth)
}

func (cornerSetter *CornerSetter) findBestSliceForCoordinatesRecursive(xy pizza.Coordinate, rowDir int, columnDir int, top bool, left bool, depth int) {

	if (xy.Row > depth && cornerSetter.Pizza.Row.End-xy.Row > depth) ||
		(xy.Column > depth && cornerSetter.Pizza.Column.End-xy.Column > depth) ||
		xy.Row < 0 || xy.Row > cornerSetter.Pizza.Row.End ||
		xy.Column < 0 || xy.Column > cornerSetter.Pizza.Column.End {
		return
	}

	possibleSlices := make([]*pizza.Slice, 0)

	for _, sl := range cornerSetter.Slicer.SliceCache[xy] {
		if ((sl.Row.Start == xy.Row && top) || (sl.Row.End == xy.Row && !top)) &&
			((sl.Column.Start == xy.Column && left) || (sl.Column.End == xy.Column && !left)) {
			possibleSlices = append(possibleSlices, sl)
		}
	}

	if cornerSetter.Pizza.Cells[xy].Slice != nil {
		if rowDir == 1 {
			xy.Row = cornerSetter.Pizza.Cells[xy].Slice.Row.End + 1
		} else if rowDir == -1 {
			xy.Row = cornerSetter.Pizza.Cells[xy].Slice.Row.Start - 1
		}

		if columnDir == 1 {
			xy.Column = cornerSetter.Pizza.Cells[xy].Slice.Column.End + 1
		} else if columnDir == -1 {
			xy.Column = cornerSetter.Pizza.Cells[xy].Slice.Column.Start - 1
		}

		cornerSetter.findBestSliceForCoordinatesRecursive(xy, rowDir, columnDir, top, left, depth)
		return
	}

	lastSlice := cornerSetter.getLastSlice(xy, rowDir, columnDir, cornerSetter.Params.CornerLastSliceDepth)
	addedSlice := (*pizza.Slice)(nil)

	if len(possibleSlices) != 0 {
		if lastSlice != nil {
			for i := 0; i < cornerSetter.Params.CornerAcceptableOffset; i++ {
				for _, sl := range possibleSlices {
					if columnDir != 0 {
						if (Abs(sl.Row.End-lastSlice.Row.End) <= i && top) ||
							(Abs(sl.Row.Start-lastSlice.Row.Start) <= i && !top) {

							if !cornerSetter.Pizza.SafeAddSlice(sl) {
								return
							}

							addedSlice = sl
							break
						}
					} else if rowDir != 0 {
						if (Abs(sl.Column.End-lastSlice.Column.End) <= i && left) ||
							(Abs(sl.Column.Start-lastSlice.Column.Start) <= i && !left) {

							if !cornerSetter.Pizza.SafeAddSlice(sl) {
								return
							}

							addedSlice = sl
							break
						}
					}
				}
			}
		}

		if addedSlice == nil {
			addedSlice = possibleSlices[0]
			for _, sl := range possibleSlices {
				/*if (rowDir != 0 && addedSlice.Row.Length() < sl.Row.Length()) ||
					(columnDir != 0 && addedSlice.Column.Length() < sl.Column.Length()){
					addedSlice = sl
				}*/
				if cornerSetter.sliceValueScorer(addedSlice, depth) < cornerSetter.sliceValueScorer(sl, depth) {
					addedSlice = sl
				}
			}

			if cornerSetter.sliceValueScorer(addedSlice, depth) >= cornerSetter.Params.CornerMinScore {
				if !cornerSetter.Pizza.SafeAddSlice(addedSlice) {
					return
				}
			}
		}
	}

	xy.Row += rowDir
	xy.Column += columnDir

	cornerSetter.findBestSliceForCoordinatesRecursive(xy, rowDir, columnDir, top, left, depth)

}

func (cornerSetter *CornerSetter) getLastSlice(xy pizza.Coordinate, rowDir int, columnDir int, goBackDepth int) *pizza.Slice {

	for i := 0; i < goBackDepth; i++ {
		xy.Row += rowDir
		xy.Column += columnDir

		if xy.Row < 0 || xy.Row >= cornerSetter.Pizza.Row.End ||
			xy.Column < 0 || xy.Column >= cornerSetter.Pizza.Column.End {
			return nil
		}

		if cornerSetter.Pizza.Cells[xy].Slice != nil {
			return cornerSetter.Pizza.Cells[xy].Slice
		}
	}

	return nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (cornerSetter *CornerSetter) ExpandThroughCorners() {

	paramSlices := make([]CornerParams, 0)

	scorePercent := make([]float32, 0)

	for _, param := range paramSlices {
		fmt.Println("\nNow trying: " + param.Name)

		cornerSetter.Pizza.RemoveAllSlice()
		cornerSetter.Params = param
		cornerSetter.setCorners()

		//slicer.FindBiggestParts()
		//slicer.FindSingles()

		//slicer.ExpandThroughEdge()
		cornerSetter.Slicer.FindSmallestParts()
		cornerSetter.Slicer.ExpandThroughDestruction()
		cornerSetter.Slicer.ExpandThroughShrink()

		_, addToScore := cornerSetter.Pizza.Score()
		scorePercent = append(scorePercent, addToScore*100)
		cornerSetter.Pizza.PrintScore()
	}

	fmt.Println()

	for i, param := range paramSlices {
		fmt.Printf(param.Name+": %f\n", scorePercent[i])
	}
}

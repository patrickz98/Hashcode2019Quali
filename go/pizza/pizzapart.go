package pizza

import "fmt"

type PizzaPart struct {
	Pizza   *Pizza
	Slices  []*Slice
	VectorR Vector
	VectorC Vector
}

func (piz PizzaPart) PrintSlices() {

	width := piz.VectorC.Length() * 2 + 1
	height := piz.VectorR.Length() * 2 + 1

	field := make([][]rune, height)

	for iny := range field {
		field[ iny ] = make([]rune, width)
		for inx := range field[ iny ] {
			field[ iny ][ inx ] = ' '
		}
	}

	for iny, yy := range piz.VectorR.Range() {
		for inx, xx := range piz.VectorC.Range() {
			field[ iny * 2 + 1 ][ inx * 2 + 1 ] = piz.Pizza.Cells[ yy ][ xx ].Type
		}
	}

	for _, sli := range piz.Slices {

		t := (sli.Row.Start    - piz.VectorR.Start) * 2 + 1
		b := (sli.Row.End      - piz.VectorR.Start) * 2 + 1
		l := (sli.Column.Start - piz.VectorC.Start) * 2
		r := (sli.Column.End   - piz.VectorC.Start) * 2 + 1

		// piz.PrintSlicesPlain()
		// sli.PrintVector()

		// fmt.Printf("t=%d b=%d\n", t, b)
		// fmt.Printf("l=%d r=%d\n", l, r)

		// break

		horizontalLenth := sli.Column.Length() * 2

		for iny := t; iny < b + 1; iny = iny + 2 {
			field[ iny ][ l ] = '|'
			field[ iny ][ l + horizontalLenth ] = '|'
		}

		for inx := l + 1; inx < r + 1; inx = inx + 2 {
			field[ t - 1 ][ inx ] = '-'
			field[ b + 1 ][ inx ] = '-'
		}
	}

	for iny := range field {
		fmt.Println(string(field[ iny ]))
	}
}

func (piz PizzaPart) PrintSlicesPlain() {

	for _, sli := range piz.Slices {
		sli.PrintVector()
	}
}

func (piz PizzaPart) Cut() []*PizzaPart {

	parts := make([]*PizzaPart, 0)

	if piz.VectorR.CutPossible() && piz.VectorC.CutPossible() {

		r1, r2 := piz.VectorR.Cut()
		c1, c2 := piz.VectorC.Cut()

		plt := piz.CutPeace(r1, c1)
		prt := piz.CutPeace(r1, c2)
		plb := piz.CutPeace(r2, c1)
		prb := piz.CutPeace(r2, c2)

		parts = append(parts, plt)
		parts = append(parts, prt)
		parts = append(parts, plb)
		parts = append(parts, prb)

		return parts
	}

	if piz.VectorR.CutPossible() {

		r1, r2 := piz.VectorR.Cut()

		pt := piz.CutPeace(r1, piz.VectorC)
		pb := piz.CutPeace(r2, piz.VectorC)

		parts = append(parts, pt)
		parts = append(parts, pb)

		return parts
	}

	if piz.VectorC.CutPossible() {

		c1, c2 := piz.VectorC.Cut()

		pl := piz.CutPeace(piz.VectorR, c1)
		pr := piz.CutPeace(piz.VectorR, c2)

		parts = append(parts, pl)
		parts = append(parts, pr)

		return parts
	}

	return parts
}

func (piz PizzaPart) CutPossible() bool {

	return piz.VectorR.CutPossible() || piz.VectorC.CutPossible()
}

func (piz PizzaPart) CutPeace(rVec Vector, cVec Vector) *PizzaPart {

	pizza := &PizzaPart{
		Pizza:   piz.Pizza,
		Slices:  []*Slice{},
		VectorR: rVec,
		VectorC: cVec,

	}

	return pizza
}

func (piz PizzaPart) PrintPart() {

	piz.Pizza.PrintVector(piz.VectorR, piz.VectorC)
}

func (piz *PizzaPart) AddSlice(slice Slice) {

	// fmt.Println("Add Slice")

	row := slice.Row
	column := slice.Column

	for iny := row.Start; iny < row.End+1; iny++ {
		for inx := column.Start; inx < column.End+1; inx++ {

			piz.Pizza.Cells[ iny ][ inx ].Slice = &slice
		}
	}

	piz.Slices = append(piz.Slices, &slice)
}

func (piz PizzaPart) PrintScore() {

	count := 0

	for _, sli := range piz.Slices {
		count += sli.Size()
	}

	total := piz.VectorC.Length() * piz.VectorR.Length()
	fmt.Printf("total cells: %d\n", total)
	fmt.Printf("covered: %d\n", count)
	fmt.Printf("percent: %.2f%%\n", (float32(count) / float32(total)) * 100)
}

func InitPizzaPart(pizza *Pizza) *PizzaPart {

	part := &PizzaPart{
		Pizza: pizza,
		Slices: []*Slice{},
		VectorR: Vector{
			Start: 0,
			End: pizza.Rows - 1,
		},
		VectorC: Vector{
			Start: 0,
			End: pizza.Columns - 1,
		},
	}

	return part
}

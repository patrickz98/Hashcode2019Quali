package pizza

func (pizza Pizza) Cut() []*Pizza {

	parts := make([]*Pizza, 0)

	if pizza.Row.CutPossible() && pizza.Column.CutPossible() {

		r1, r2 := pizza.Row.Cut()
		c1, c2 := pizza.Column.Cut()

		plt := pizza.CutPeace(r1, c1)
		prt := pizza.CutPeace(r1, c2)
		plb := pizza.CutPeace(r2, c1)
		prb := pizza.CutPeace(r2, c2)

		parts = append(parts, plt)
		parts = append(parts, prt)
		parts = append(parts, plb)
		parts = append(parts, prb)

		return parts
	}

	if pizza.Row.CutPossible() {

		r1, r2 := pizza.Row.Cut()

		pt := pizza.CutPeace(r1, pizza.Column)
		pb := pizza.CutPeace(r2, pizza.Column)

		parts = append(parts, pt)
		parts = append(parts, pb)

		return parts
	}

	if pizza.Column.CutPossible() {

		c1, c2 := pizza.Column.Cut()

		pl := pizza.CutPeace(pizza.Row, c1)
		pr := pizza.CutPeace(pizza.Row, c2)

		parts = append(parts, pl)
		parts = append(parts, pr)

		return parts
	}

	return parts
}

func (pizza Pizza) CutPossible() bool {

	return pizza.Row.CutPossible() || pizza.Column.CutPossible()
}

func (pizza Pizza) CutPeace(rVec Vector, cVec Vector) *Pizza {

	part := &Pizza{
		Ingredients: pizza.Ingredients,
		MaxCells:    pizza.MaxCells,
		Cells:       pizza.Cells,
		Row:         rVec,
		Column:      cVec,
	}

	return part
}


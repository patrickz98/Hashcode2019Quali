package slicer

import (
	"../pizza"
	"fmt"
)

func (slicer *Slicer) findBalanced(xy pizza.Coordinate) Slices {

	if slicer.Pizza.HasSliceAt(xy) {
		return nil
	}

	slices := slicer.SliceCache[ xy ]

	balanced := make(Slices, 0)

	for _, sli := range slices {

		balance := sli.IngredientsBalance()

		if balance != 0 {
			continue
		}

		balanced = append(balanced, sli)
	}

	return balanced
}

func (slicer *Slicer) findSlicesWithUnbalanceAt(xy pizza.Coordinate) (balanced Slices, tomato Slices, mushrooms Slices) {

	for _, sli := range slicer.SliceCache[ xy ] {

		tomatoCount, mushroomCount := sli.IngredientsCount()

		if tomatoCount > mushroomCount {
			tomato = append(tomato, sli)
		}

		if tomatoCount < mushroomCount {
			mushrooms = append(mushrooms, sli)
		}

		if tomatoCount == mushroomCount {
			balanced = append(balanced, sli)
		}
	}

	return balanced, tomato, mushrooms
}

func (slicer *Slicer) findSlicesWithUnbalance() (balanced Slices, tomato Slices, mushrooms Slices) {

	for _, xy := range slicer.Pizza.Traversal() {

		balanceSli, tomatoSli, mushroomSli := slicer.findSlicesWithUnbalanceAt(xy)

		balanced = append(balanced, balanceSli...)
		tomato = append(tomato, tomatoSli...)
		mushrooms = append(mushrooms, mushroomSli...)
	}

	return balanced, tomato, mushrooms
}

func (slicer *Slicer) ExpandBalanced() {

	for _, xy := range slicer.Pizza.Traversal() {

		slices := slicer.findBalanced(xy)

		for _, sli := range slices {

			if slicer.overlap(sli) {
				continue
			}

			slicer.AddSlice(sli)
			break
		}
	}
}

func (slicer *Slicer) ExpandBalancedIntelligent() {

	tomato, mushroom := slicer.Pizza.IngredientsCount()
	balance := tomato - mushroom

	fmt.Printf("tomato=%d mushroom=%d\n", tomato, mushroom)
	// fmt.Printf("tomato=%d mushroom=%d\n", tomato % slicer.Pizza.IngredientsCount, mushroom % slicer.Pizza.IngredientsCount)
	fmt.Printf("tomato balance = %f\n",   1 * (float32(tomato)   / float32(slicer.Pizza.Size())))
	fmt.Printf("mushroom balance = %f\n", 1 * (float32(mushroom) / float32(slicer.Pizza.Size())))
	fmt.Printf("balance=%d\n", balance)

	sliBalanced, sliTomato, sliMushrooms := slicer.findSlicesWithUnbalance()

	fmt.Printf("sliTomato = %d\n", len(sliTomato))
	fmt.Printf("sliMushrooms = %d\n", len(sliMushrooms))
	fmt.Printf("sliBalanced = %d\n", len(sliBalanced))

	bal := 0

	queue := InitCoordinateQueue()

	for _, xy := range slicer.Pizza.Traversal() {
		queue.Push(xy)
	}

	for queue.HasItems() {

		xy := *queue.Pop()
		slices := slicer.SliceCache[ xy ]

		for _, sli := range slices {

			tom1, mus1 := sli.IngredientsCount()

			tmp := bal + (tom1 - mus1)
		}
	}
}

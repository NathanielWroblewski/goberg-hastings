package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "sync"

type parallelMapHandler func(int, *Lattice) int

// Lattice is a two-dimensional, square lattice of cells.  Each cell is an
// integer corresponding to a stage.
type Lattice struct {
	Cells  *[]int
	Height int
	Width  int
}

func (lattice *Lattice) parallelMap(fn parallelMapHandler) *Lattice {
	wg := sync.WaitGroup{}
	cells := *lattice.Cells
	results := make([]int, len(cells))

	for index := range cells {
		wg.Add(1)

		go func(i int, lat *Lattice, mut *[]int) {
			defer wg.Done()
			(*mut)[i] = fn(i, lat)
		}(index, lattice, &results)
	}

	wg.Wait()

	return &Lattice{
		Cells:  &results,
		Height: lattice.Height,
		Width:  lattice.Width,
	}
}

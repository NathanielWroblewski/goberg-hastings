package utils

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

type Neighborhoods func(int, *[]int, int, int) *[]int

// MooreNeighborhood returns the cells surrounding a central cell on a
// two-dimensional square lattice.  If the extent is 1 (ρ = 1), the Moore
// neighborhood is comprised of the eight cells surrounding the given central
// cell.
func MooreNeighborhood(index int, lattice *[]int, width int, rho int) *[]int {
	var neighborhood []int

	for x := -rho; x <= rho; x++ {
		for y := -rho; y <= rho; y++ {
			if x == 0 && y == 0 {
				continue
			}

			neighborhood = append(neighborhood, at(index+x+(y*width), lattice))
		}
	}

	return &neighborhood
}

// VonNeumannNeighborhood returns the cells abutting a central cell on a
// two-dimensional square lattice.  If the extent is 1 (ρ = 1), the von Neumann
// neighborhood is comprised of the four cells abutting the given central cell.
func VonNeumannNeighborhood(index int, lattice *[]int, width int, rho int) *[]int {
	var neighborhood []int

	for i := -rho; i <= rho; i++ {
		if i == 0 {
			continue
		}

		neighborhood = append(neighborhood, at(index+i, lattice))
		neighborhood = append(neighborhood, at(index+(i*width), lattice))
	}

	return &neighborhood
}

func wrap(index int, lattice *[]int) int {
	if index > len(*lattice)-1 {
		return wrap(index-len(*lattice), lattice)
	}

	if index < 0 {
		return wrap(len(*lattice)+index, lattice)
	}

	return index
}

func at(index int, lattice *[]int) int {
	return (*lattice)[wrap(index, lattice)]
}

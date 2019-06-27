package models

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import rand "math/rand"

import utils "../utils"

// CellularAutomaton is a two-dimensional square lattice of cells, each of which
// is in a certain stage.  This struct also contains functions for discrete time
// updates in which each cell progresses to its next stage.
type CellularAutomaton struct {
	Lattice      *Lattice
	Parameters   *Parameters
	GetNeighbors utils.Neighborhoods
	ShouldUpdate utils.Rules
}

// Init initializes the lattice by setting each cell to a random stage
func (automaton *CellularAutomaton) Init() *CellularAutomaton {
	return automaton.parallelMap(randomize)
}

// Step evaluates each cell against its local neighborhood and returns the
// appropriate stage each cell
func (automaton *CellularAutomaton) Step() *CellularAutomaton {
	return automaton.parallelMap(evaluate)
}

func (automaton *CellularAutomaton) parallelMap(fn handler) *CellularAutomaton {
	lattice := automaton.Lattice.parallelMap(func(index int, lattice *Lattice) int {
		return fn(index, automaton)
	})

	return &CellularAutomaton{
		Lattice:      lattice,
		Parameters:   automaton.Parameters,
		GetNeighbors: automaton.GetNeighbors,
		ShouldUpdate: automaton.ShouldUpdate,
	}
}

type handler func(int, *CellularAutomaton) int

func (automaton *CellularAutomaton) at(index int) int {
	return (*automaton.Lattice.Cells)[index]
}

func randomize(index int, automaton *CellularAutomaton) int {
	return rand.Intn(automaton.Parameters.Kappa)
}

func evaluate(index int, automaton *CellularAutomaton) int {
	currentState := automaton.at(index)
	nextState := (currentState + 1) % automaton.Parameters.Kappa
	neighborhood := automaton.GetNeighbors(
		index,
		automaton.Lattice.Cells,
		automaton.Lattice.Width,
		automaton.Parameters.Rho,
	)

	if automaton.ShouldUpdate(neighborhood, currentState, nextState, automaton.Parameters.Theta) {
		return nextState
	} else {
		return currentState
	}
}

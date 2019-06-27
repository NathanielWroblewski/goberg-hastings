package utils

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

type Rules func(*[]int, int, int, int) bool

// Cyclic is a predicate encapsulating the conditions which allow a particular
// cell to advance to the next stage.  Cyclic cellular automata advance from
// stage k to k + 1 mod κ by contact with at least a threshold θ number of
// successor stage cells in a prescribed local neighborhood of some extent ρ.
func Cyclic(neighborhood *[]int, currentState int, nextState int, theta int) bool {
	count := 0

	for _, value := range *neighborhood {
		if value == nextState {
			count++
		}
	}

	return count >= theta
}

// GreenbergHastings (GH) is a predicate encapsulating the conditions which
// allow a particular cell to advance to its next stage. GH automata are a
// simpler version of cyclic cellular automata. Under the GH rules, only one
// stage advances by contact with at least a threshold θ number of successor
// colors in its local neighborhood of some extent ρ, and all other stages
// advance automatically.
func GreenbergHastings(neighborhood *[]int, currentState int, nextState int, theta int) bool {
	return currentState > 0 || Cyclic(neighborhood, currentState, nextState, theta)
}

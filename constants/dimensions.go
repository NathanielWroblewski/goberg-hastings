package constants

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

// Height is the default height in pixels of the lattice and corresponding PNG
const Height = 500

// HeightMin is the minimum acceptable number of cells
const HeightMin = 1

// HeightDescription is the command-line flag description for the -height flag
const HeightDescription = "the height in cells of the two-dimensional, square lattice."

// Width is the default width in pixels of the lattice and corresponding PNG
const Width = 1000

// WidthMin is the minimum acceptable number of cells
const WidthMin = 1

// WidthDescription is the command-line flag description for the -width flag
const WidthDescription = "the width in cells of the two-dimensional, square lattice."

// Iterations is the default total count of discrete time steps
const Iterations = 500

// IterationsMin is the minimum number of acceptable iterations
const IterationsMin = 1

// IterationsMax is the maximum number of acceptable iterations. There is no
// hard limit, but the output.png currently uses a template destination name
// with '%06d' padding.
const IterationsMax = 999999

// IterationsDescription is the command-line flag description for the -iters flag
const IterationsDescription = "the total count of discrete time steps."

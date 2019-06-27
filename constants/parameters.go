package constants

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

// Kappa (κ) is the default number of colors/states each cell of the cellular
// automata can be in
const Kappa = len(ColorWheel)

// KappaMin is the minimum acceptable value for Kappa
const KappaMin = 1

// KappaMax is the maximum acceptable value for Kappa
const KappaMax = len(ColorWheel)

// KappaDescription is the command-line flag description for the -kappa flag
const KappaDescription = "κ, the number of states."

// Theta (θ) is the default threshold number of successor colors in a
// neighborhood local to each cell required to advance that cell in stage/color
const Theta = 1

// ThetaMin is the minimum acceptable value for Theta
const ThetaMin = 0

// ThetaDescription is the command-line flag description for the -theta flag
const ThetaDescription = "θ, the threshold number of successor states in a local neighborhood required for cells to advance in state."

// Rho (ρ) is the default extent of the neighborhood
const Rho = 1

// RhoMin is the minimum acceptable value for Rho
const RhoMin = 1

// RhoDescription is the command-line flag description for the -rho flag
const RhoDescription = "ρ, the extent of the neighborhood."

// MooreDescription is the command-line flag description for the -moore flag
const MooreDescription = "if set, use Moore neighborhoods; if unset, use von Neumann neighborhoods"

// CyclicDescription is the command-line flag description for the -cyclic flag
const CyclicDescription = "if set, use cyclic rules; if unset, use Greenberg-Hasting rules"

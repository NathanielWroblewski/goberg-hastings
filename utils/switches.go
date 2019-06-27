package utils

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "flag"
import "fmt"
import "os"

import constants "../constants"

type Switches struct {
	Kappa        int
	Theta        int
	Rho          int
	Width        int
	Height       int
	Iterations   int
	Neighborhood Neighborhoods
	Rules        Rules
}

func ParseSwitches() *Switches {
	kappa := flag.Int("kappa", constants.Kappa, fmt.Sprintf(
		"%s. bounds=[%d, %d]",
		constants.KappaDescription,
		constants.KappaMin,
		constants.KappaMax,
	))
	theta := flag.Int("theta", constants.Theta, fmt.Sprintf(
		"%s. min=%d",
		constants.ThetaDescription,
		constants.ThetaMin,
	))
	rho := flag.Int("rho", constants.Rho, fmt.Sprintf(
		"%s. min=%d",
		constants.RhoDescription,
		constants.RhoMin,
	))
	width := flag.Int("width", constants.Width, fmt.Sprintf(
		"%s. min=%d",
		constants.WidthDescription,
		constants.WidthMin,
	))
	height := flag.Int("height", constants.Height, fmt.Sprintf(
		"%s. min=%d",
		constants.HeightDescription,
		constants.HeightMin,
	))
	iters := flag.Int("iters", constants.Iterations, fmt.Sprintf(
		"%s. bounds=[%d, %d]",
		constants.IterationsDescription,
		constants.IterationsMin,
		constants.IterationsMax,
	))
	moore := flag.Bool("moore", false, constants.MooreDescription)
	cyclic := flag.Bool("cyclic", false, constants.CyclicDescription)

	flag.Parse()

	if *kappa < constants.KappaMin || *kappa > constants.KappaMax ||
		*theta < constants.ThetaMin || *rho < constants.RhoMin ||
		*width < constants.WidthMin || *height < constants.HeightMin ||
		*iters < constants.IterationsMin || *iters > constants.IterationsMax {
		flag.PrintDefaults()
		os.Exit(2)
	}

	neighborhood := VonNeumannNeighborhood
	if *moore {
		neighborhood = MooreNeighborhood
	}

	rule := GreenbergHastings
	if *cyclic {
		rule = Cyclic
	}

	return &Switches{
		Kappa:        *kappa,
		Theta:        *theta,
		Rho:          *rho,
		Width:        *width,
		Height:       *height,
		Iterations:   *iters,
		Neighborhood: neighborhood,
		Rules:        rule,
	}
}

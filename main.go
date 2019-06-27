package main

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import constants "./constants"
import models "./models"
import views "./views"
import utils "./utils"

import "fmt"

func main() {
	switches := utils.ParseSwitches()
	params := models.Parameters{
		Kappa: switches.Kappa,
		Theta: switches.Theta,
		Rho:   switches.Rho,
	}
	cells := make([]int, switches.Height*switches.Width)
	lattice := models.Lattice{
		Cells:  &cells,
		Height: switches.Height,
		Width:  switches.Width,
	}
	automaton := models.CellularAutomaton{
		Lattice:      &lattice,
		Parameters:   &params,
		GetNeighbors: switches.Neighborhood,
		ShouldUpdate: switches.Rules,
	}

	for index := 0; index < switches.Iterations; index++ {
		if index == 0 {
			automaton = *automaton.Init()
		} else {
			automaton = *automaton.Step()
		}

		img := views.PNG(
			automaton.Lattice.Cells,
			switches.Width,
			switches.Height,
			&constants.ColorWheel,
		)
		utils.WritePNG(img, fmt.Sprintf("./output/%06d.png", index))
		fmt.Print(views.Progress(index, switches.Iterations))
	}
}

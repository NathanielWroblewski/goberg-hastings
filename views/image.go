package views

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "image"
import "image/color"

import constants "../constants"

// PNG outputs a lattice as a .png image
func PNG(lattice *[]int, width int, height int, colors *[15]color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for index, stage := range *lattice {
		x := index % width
		y := index / width

		img.Set(x, y, constants.ColorWheel[stage])
	}

	return img
}

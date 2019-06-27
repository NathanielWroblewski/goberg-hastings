package utils

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "os"
import "image"
import png "image/png"

// WritePNG accepts a PNG image and a filepath and writes the file to the path
func WritePNG(img *image.RGBA, dest string) {
	fp, _ := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0600)
	defer fp.Close()
	png.Encode(fp, img)
}

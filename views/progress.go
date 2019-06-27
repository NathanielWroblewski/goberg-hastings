package views

// Copyright (c) 2019 Nathaniel Wroblewski
// I am making my contributions/submissions to this project solely in my personal
// capacity and am not conveying any rights to any intellectual property of any
// third parties.

import "fmt"

// Progress returns a string summarizing the progress of the form:
// 100% (137921 / 137921 bytes)
func Progress(is int, of int) string {
	percent := float64(is) / float64(of) * 100.0

	if percent >= 100.0 {
		return fmt.Sprintf("%3.2f%% (%d / %d iterations)\r\n", percent, is, of)
	} else {
		return fmt.Sprintf("%3.2f%% (%d / %d iterations)\r", percent, is, of)
	}
}

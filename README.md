Goberg-Hastings
===

This repository contains a golang script for generating `.png` images of cyclic
cellular automata or Greenberg-Hastings automata.  The images can be combined
via `ffmpeg` into `.mp4` movies.

Cyclic Cellular Automata
---

Imagine a two-dimensional lattice composed of cells, and each cell of the lattice is painted one of κ colors. The colors are arranged along a color wheel, and the colors advance (k to k + 1 mod κ) by contact with at least a threshold θ number of successor colors in a prescribed local neighborhood (typically either a von Neumann neighborhood or a Moore neighborhood) of some extent ρ. Discrete-time parallel systems of this sort are called cyclic cellular automata (CCA). Initialized in random configurations, these rules exhibit complex self-organization, typically characterized by nucleation of locally periodic spirals and a variety of equilibria that display large-scale stochastic wave fronts. CCA emulate the behavior of a wide range of complex, coherent, periodic wave phenomena in space.

The Greenberg-Hastings (GH) automaton is a simpler version of CCA. Under the GH rules, only one color advances by contact with at least a threshold θ number of successor colors in its local neighborhood ρ, and all other colors advance automatically. The simplest GH model is the automaton with three states, or colors (κ=3): resting, excited, refractory. If a resting cell is in contact with an excited cell, it will become excited on the next time-step; otherwise, it will remain at rest. Once excited, it proceeds automatically through the refractory states until it returns again to a state of rest.

For more information, or to experiment with an interactive version in javascript, [click here](https://nathanielwroblewski.github.io/greenberg-hastings/).

Run
---

Build the binary

```sh
go build -o goberg ./main.go
```

Run the binary
```sh
./goberg -width=1000 -height=500 -kappa=15 -theta=1 -rho=1 -moore=true -cyclic=true -iters=500
```

To make a video, you can use [`ffmpeg`](https://ffmpeg.org/)
```sh
ffmpeg -framerate 30 -pattern_type glob -i './output/*.png' \
  -c:v libx264 -pix_fmt yuv420p ./video.mp4
```

Examples
---

```sh
$ ./goberg -kappa=15 -theta=1 -rho=1 -moore=false -cyclic=true -iters=500
$ ffmpeg -framerate 30 -pattern_type glob -i './output/*.png' \
  -c:v libx264 -pix_fmt yuv420p ./k15-t1-r1-vnn-cca.mp4
```

```sh
$ ./goberg -kappa=8 -theta=8 -rho=4 -moore=true -cyclic=false =iters=200
$ ffmpeg -framerate 15 -pattern_type glob -i './output/*.png' \
  -c:v libx264 -pix_fmt yuv420p ./k8-t8-r4-moore-gh.mp4
```

Copyright (c) 2018 Nathaniel Wroblewski
I am making my contributions/submissions to this project solely in my personal
capacity and am not conveying any rights to any intellectual property of any
third parties.

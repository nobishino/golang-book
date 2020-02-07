package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 600
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>",
		width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, finite := corner(i+1, j)
			if !finite {
				continue
			}
			bx, by, finite := corner(i, j)
			if !finite {
				continue
			}
			cx, cy, finite := corner(i, j+1)
			if !finite {
				continue
			}
			dx, dy, finite := corner(i+1, j+1)
			if !finite {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g, %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (sx, sy float64, finite bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx = width/2 + (x-y)*cos30*xyscale
	sy = width/2 + (x+y)*sin30*xyscale - z*zscale
	finite = !math.IsNaN(sx) && !math.IsInf(sx, 0) && !math.IsNaN(sy) && !math.IsInf(sy, 0)
	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

package svg

import (
	"fmt"
	"io"
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

func Write(w io.Writer) {
	startTag := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>",
		width, height)
	io.WriteString(w, startTag)
	maxZ, minZ := getMaxMinZ()
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
			polygonTag := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g, %g,%g' %s/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getFillColor(i, j, maxZ, minZ))
			io.WriteString(w, polygonTag)
		}
	}
	endTag := fmt.Sprintf("</svg>")
	io.WriteString(w, endTag)
}

func getHeight(i, j int) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return f(x, y)
}

func getMaxMinZ() (max, min float64) {
	max, min = -math.MaxFloat64, math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j)
			if math.IsInf(z, 0) || math.IsNaN(z) {
				continue
			}
			max = math.Max(max, z)
			min = math.Min(min, z)
		}
	}
	return
}

func getFillColor(i, j int, maxZ, minZ float64) string {
	relativeZ := byte(255 * (getHeight(i, j) - minZ) / (maxZ - minZ))
	return fmt.Sprintf("style=\"fill:#%02x00%02x\"", relativeZ, byte(255)-relativeZ)
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
	return sinr(x, y)
}

func sinr(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	omega := 0.3
	r := 0.5
	return r * math.Pow(math.Sin(omega*x)*math.Sin(omega*y), 2.0)
}

func saddle(x, y float64) float64 {
	r := 0.2
	coeff := 0.1
	return r * (math.Pow(coeff*x, 2.) - math.Pow(coeff*y, 2.))
}

func mogul(x, y float64) float64 {
	omega, r := 0.4, 0.2
	return r * (math.Abs(math.Sin(omega*x)) + math.Abs(math.Sin(omega*y)))
}

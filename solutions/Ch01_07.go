package solutions

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func Problem01_12() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type lParam struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err)
		return
	}
	param := lParam{
		5, 0.001, 300, 64, 8,
	}
	for k, v := range r.Form {
		var err error
		fmt.Println(k, v)
		switch k {
		case "cycles":
			param.Cycles, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		case "size":
			param.Size, err = strconv.Atoi(v[0])
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		}
	}

	var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)

	lissajous := func(out io.Writer, param lParam) {
		// fmt.Println(param)
		var (
			cycles  int = param.Cycles  //5     // number of complete x oscillator revolutions
			res         = param.Res     //0.001 // angular resolution
			size        = param.Size    //100   // image canvas covers [-size..+size]
			nframes     = param.Nframes //64    // number of animation frames
			delay       = param.Delay   //8     // delay between frames in 10ms units
		)
		freq := rand.Float64() * 3.0 // relative frequency of y oscillator
		anim := gif.GIF{LoopCount: nframes}
		phase := 0.0 // phase difference
		for i := 0; i < nframes; i++ {
			rect := image.Rect(0, 0, 2*size+1, 2*size+1)
			img := image.NewPaletted(rect, palette)
			for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase)
				img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
					blackIndex)
			}
			phase += 0.1
			anim.Delay = append(anim.Delay, delay)
			anim.Image = append(anim.Image, img)
		}
		gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
	}

	lissajous(w, param)

}

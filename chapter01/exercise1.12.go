package main

// Server1 is a minimal "echo" server.

import (
	"log"
	"net/http"
	"image/color"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100
		nframes = 64
		delay   = 8
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}


func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	cycle := 5
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for _,v := range r.Form {
		cycle,_ = strconv.Atoi(v[0])
	//	fmt.Fprintf(w,"Form[%q] = %q\n", k,v)
	}
	//fmt.Fprintf(w,"%s %s %s\n", r.Method, r.URL, r.Proto)
	//fmt.Fprintf(w, "cycle = %d\n", cycle)
	lissajous(w,cycle)
}
package K6Widgets

import (
	"fmt"
	"image/color"
	"math"

	wolffCanvas "github.com/tdewolff/canvas"
)

var (
	Numeros = map[int][7]int{
		0: {1, 1, 1, 1, 1, 1, 0},
		1: {0, 1, 1, 0, 0, 0, 0},
		2: {1, 1, 0, 1, 1, 0, 1},
		3: {1, 1, 1, 1, 0, 0, 1},
		4: {0, 1, 1, 0, 0, 1, 1},
		5: {1, 0, 1, 1, 0, 1, 1},
		6: {1, 0, 1, 1, 1, 1, 1},
		7: {1, 1, 1, 0, 0, 0, 0},
		8: {1, 1, 1, 1, 1, 1, 1},
		9: {1, 1, 1, 0, 0, 1, 1},
	}

	Colores = []color.Color{
		wolffCanvas.Hex("#333333"),
		wolffCanvas.Hex("#ffffff"),
	}
)

type sevenSegPaint struct {
	k6paint

	ndigits            int
	border, bckgTransp bool
}

func CreateSevenSegPaint(w, h float64, ndig int, br, bckgT bool) *sevenSegPaint {
	aux := sevenSegPaint{}
	aux.width = w
	aux.height = h
	aux.ndigits = ndig
	aux.border = br
	aux.bckgTransp = bckgT

	aux.CrearPintura(w, h)

	aux.drawLayer0Marco()
	aux.drawLayer1Numeros(0)

	return &aux
}

func (cp *sevenSegPaint) Refresh(num int) {
	cp.Reset()

	cp.drawLayer0Marco()
	cp.drawLayer1Numeros(num)
}

func (cp *sevenSegPaint) drawLayer0Marco() {
	p := cp.pincel
	//p.SetZIndex(0)

	if cp.bckgTransp {
		p.SetFillColor(wolffCanvas.Transparent)
	} else {
		p.SetFillColor(wolffCanvas.Black)
	}
	if cp.border {
		p.SetStrokeWidth(2)
		p.SetStroke(wolffCanvas.White)
	}
	p.DrawPath(0, 0, wolffCanvas.Rectangle(float64(cp.width), float64(cp.height)))

	p.ResetStyle()
	p.ResetView()
}

func (cp *sevenSegPaint) drawLayer1Numeros(num int) {
	p := cp.pincel
	//p.SetZIndex(1)

	if cp.bckgTransp {
		p.SetFillColor(wolffCanvas.Transparent)
	} else {
		p.SetFillColor(wolffCanvas.Black)
	}
	p.DrawPath(2, 2, wolffCanvas.Rectangle(float64(cp.width-4), float64(cp.height-4)))

	if !cp.border {
		aux3 := float64(cp.width) / float64(cp.ndigits)
		p.Scale(float64(aux3/80), float64(cp.height/100))
		//p.Scale(float64(cp.height/100), float64(cp.height/100))
	} else {
		aux1 := float64(cp.height * 0.8)
		aux2 := float64(cp.height * 0.1)
		aux3 := (float64(cp.width) - aux2) / float64(cp.ndigits)
		p.Translate(aux2, aux2)
		p.Scale(aux3/80, aux1/100)
	}
	cp.drwaNumber(num, p)

	p.ResetStyle()
	p.ResetView()
}

func (cp *sevenSegPaint) drwaNumber(num int, p *wolffCanvas.Context) {
	if num >= int(math.Pow10(cp.ndigits)) {
		num = int(math.Pow10(cp.ndigits)) - 1
	}
	if num < 1 {
		num = 0
	}

	cum := 0
	for i := 0; i < cp.ndigits; i++ {
		x := (num - cum) / int(math.Pow10(cp.ndigits-i-1))
		cum += x * int(math.Pow10(cp.ndigits-i-1))
		cp.drwaDigit(x, i, p)
	}
}
func (cp *sevenSegPaint) drwaDigit(n, pos int, p *wolffCanvas.Context) {
	dx := pos * 80
	var hm2, hm4, hm6 int = 43, 50, 57
	var wi1, wi2, wd1, wd2 = 8 + dx, 15 + dx, 56 + dx, 63 + dx

	s1 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wi1, hm4+43, wi2, hm6+43, wd1, hm6+43, wd2, hm4+43, wd1, hm2+43, wi2, hm2+43)
	s2 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wd2+1, hm4+1, wd2-6, hm6+1, wd2-6, hm2+42, wd2+1, hm4+42, wd2+8, hm2+42, wd2+8, hm6+1)
	s3 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wd2+1, hm4-42, wd2-6, hm6-42, wd2-6, hm2-1, wd2+1, hm4-1, wd2+8, hm2-1, wd2+8, hm6-42)
	s4 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wi1, hm4-43, wi2, hm6-43, wd1, hm6-43, wd2, hm4-43, wd1, hm2-43, wi2, hm2-43)
	s5 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wi1-1, hm4-42, wi1-8, hm6-42, wi1-8, hm2-1, wi1-1, hm4-1, wi1+6, hm2-1, wi1+6, hm6-42)
	s6 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wi1-1, hm4+1, wi1-8, hm6+1, wi1-8, hm2+42, wi1-1, hm4+42, wi1+6, hm2+42, wi1+6, hm6+1)
	s7 := fmt.Sprintf("M %v   %v  L %v %v  L %v %v  L %v %v L %v %v  L %v %v Z\n",
		wi1, hm4, wi2, hm6, wd1, hm6, wd2, hm4, wd1, hm2, wi2, hm2)

	NumCoordPaths := []string{s1, s2, s3, s4, s5, s6, s7}
	for in, el := range NumCoordPaths {
		p.SetFillColor(Colores[Numeros[n][in]])
		p.SetStroke(Colores[Numeros[n][in]])
		path, _ := wolffCanvas.ParseSVGPath(el)
		p.DrawPath(0, 0, path)
	}
}

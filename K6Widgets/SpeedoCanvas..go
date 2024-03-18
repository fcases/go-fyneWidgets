package K6Widgets

import (
	"fmt"
	"math"

	wolffCanvas "github.com/tdewolff/canvas"
)

var (
	startAngle float64 = -150
	endAngle   float64 = 150
	minValue   float64 = 0
	maxValue   float64 = 100
	k          float64 = (endAngle - startAngle) / (maxValue - minValue)
)

type speedoPaint struct {
	k6paint

	side, radius float64
	margins      int
}

func CreateSpeedoPaint(w, h float64) *speedoPaint {
	aux := speedoPaint{}
	aux.width = w
	aux.height = h
	aux.margins = 8
	aux.side = math.Min(float64(w), float64(h))
	aux.radius = (aux.side - float64(aux.margins)) / 2.0

	aux.CrearPintura(w, h)

	aux.drawLayer0Esferas()
	aux.drawLayer1Flecha(0)
	aux.drawLayer2Capuchon()

	return &aux
}

func (cp *speedoPaint) Refresh(or int) {
	cp.Reset()

	cp.drawLayer0Esferas()
	cp.drawLayer1Flecha(or)
	cp.drawLayer2Capuchon()
}

func (cp *speedoPaint) drawLayer0Esferas() {
	p := cp.pincel
	//p.SetZIndex(0)

	// rectangulo inicial
	p.DrawPath(0, 0, wolffCanvas.Rectangle(float64(cp.width), float64(cp.height)))
	p.Translate(float64(cp.width/2), float64(cp.height/2))

	p.Scale(float64(0.87*cp.radius/100), float64(0.87*cp.radius/100))

	cp.drawPanel(p)
	cp.drawScaleNum(p)
	cp.drawScaleLine(p)

	p.ResetStyle()
	p.ResetView()
}

func (cp *speedoPaint) drawPanel(p *wolffCanvas.Context) {
	// Aros
	lg := wolffCanvas.NewLinearGradient(wolffCanvas.Point{X: 0, Y: 0},
		wolffCanvas.Point{X: float64(cp.width), Y: float64(cp.height)})
	lg.Add(0.1, wolffCanvas.Hex("#2f2f2f"))
	lg.Add(0.5, wolffCanvas.Hex("#808080"))
	lg.Add(0.7, wolffCanvas.Hex("#e8e8e8"))
	lg.Add(0.9, wolffCanvas.Hex("#f0f0f0"))
	p.SetFillGradient(lg)
	p.DrawPath(0, 0, wolffCanvas.Circle(113))
	p.SetFillColor(wolffCanvas.Black)
	p.DrawPath(0, 0, wolffCanvas.Circle(100))
}

func (cp *speedoPaint) drawScaleNum(p *wolffCanvas.Context) {
	monoSpace := wolffCanvas.NewFontFamily("Monospace")
	if err := monoSpace.LoadLocalFont("DroidSansMono", wolffCanvas.FontRegular); err != nil {
		panic(err)
	}
	TamFont := 38 * cp.radius / 100
	face := monoSpace.Face(TamFont, wolffCanvas.White, wolffCanvas.FontRegular)

	for i := int(minValue); i <= int(maxValue); i++ {
		if i%10 == 0 {
			// Letras (pasan del rotate y del scale, asique con rendertext que tiene matriz trafo)
			a := float64(i)*k - 90 + startAngle
			Ax := float64(cp.side)/2 + 75*cp.radius/100*math.Cos(a*math.Pi/180)
			Ay := float64(cp.side)/2 - 2*k + 75*cp.radius/100*math.Sin(-a*math.Pi/180)
			miMatriz := wolffCanvas.Matrix{{1, 0, Ax},
				{0, 1, Ay}}
			p.RenderText(wolffCanvas.NewTextLine(face, fmt.Sprintf("%v", i),
				wolffCanvas.Middle), miMatriz)
		}
	}
	miMatriz := wolffCanvas.Matrix{{1, 0, float64(cp.side) / 2},
		{0, 1, float64(cp.side)/2 + 9*k}}
	p.RenderText(wolffCanvas.NewTextLine(face, "Km/h", wolffCanvas.Middle), miMatriz)
}

func (cp *speedoPaint) drawScaleLine(p *wolffCanvas.Context) {
	// lineas de colores
	p.SetFillColor(wolffCanvas.Transparent)

	radio := float64(52)
	p.SetStrokeWidth(4)
	p.SetStrokeColor(wolffCanvas.Lime)
	pathstr := fmt.Sprintf("M %v 0 A %v %v %v %v %v %v %v\n",
		radio,
		radio, radio, 0, 1, 1,
		radio*math.Cos(240*math.Pi/180),
		radio*math.Sin(240*math.Pi/180))
	path, _ := wolffCanvas.ParseSVGPath(pathstr)
	p.DrawPath(0, 0, path)

	p.SetStrokeColor(wolffCanvas.Red)
	pathstr = fmt.Sprintf("M %v 0 A %v %v %v %v %v %v %v\n",
		radio,
		radio, radio, 0, 0, 0,
		radio*math.Cos(300*math.Pi/180),
		radio*math.Sin(300*math.Pi/180))
	path, _ = wolffCanvas.ParseSVGPath(pathstr)
	p.DrawPath(0, 0, path)

	radio = float64(59)
	p.SetStrokeWidth(1)
	p.SetStrokeColor(wolffCanvas.White)
	pathstr = fmt.Sprintf("M %v 0 A %v %v %v %v %v %v %v\n",
		radio,
		radio, radio, 0, 1, 1,
		radio*math.Cos(240*math.Pi/180),
		radio*math.Sin(240*math.Pi/180))
	path, _ = wolffCanvas.ParseSVGPath(pathstr)
	p.DrawPath(0, 0, path)

	p.SetStrokeColor(wolffCanvas.Red)
	pathstr = fmt.Sprintf("M %v 0 A %v %v %v %v %v %v %v\n",
		radio,
		radio, radio, 0, 0, 0,
		radio*math.Cos(300*math.Pi/180),
		radio*math.Sin(300*math.Pi/180))
	path, _ = wolffCanvas.ParseSVGPath(pathstr)
	p.DrawPath(0, 0, path)

	// escalas
	scaleNums := maxValue - minValue
	angleStep := (endAngle - startAngle) / scaleNums
	p.Rotate(-startAngle) // el angulo cero es a las 12, grados positivos=ccw
	p.SetStrokeColor(wolffCanvas.White)
	for i := 0; i < int(scaleNums)+1; i++ {

		if i >= int(0.8*scaleNums) {
			p.SetStrokeColor(wolffCanvas.Red)
		} else {
			p.SetStrokeColor(wolffCanvas.White)
		}

		if i%10 == 0 {
			p.SetStrokeWidth(3)
			path, _ := wolffCanvas.ParseSVGPath("M 0 60 L 0 70 ")
			p.DrawPath(0, 0, path)
		} else if i%2 == 0 {
			p.SetStrokeWidth(1)
			path, _ := wolffCanvas.ParseSVGPath("M 0 60 L 0 65")
			p.DrawPath(0, 0, path)
		}
		p.Rotate(-angleStep)
	}
}

func (cp *speedoPaint) drawLayer1Flecha(alfa int) {
	p := cp.pincel
	//p.SetZIndex(1)

	p.SetFillColor(wolffCanvas.Transparent)
	p.DrawPath(0, 0, wolffCanvas.Rectangle(float64(cp.width), float64(cp.height)))
	p.Translate(float64(cp.width/2), float64(cp.height/2))
	p.Scale(float64(0.87*cp.radius/100), float64(0.87*cp.radius/100))

	// alfa from 0 to 100, must trasnlated to degrees
	alfa *= int(k)
	p.Rotate(-startAngle - float64(alfa))

	cp.drawIndicator(p)

	p.ResetStyle()
	p.ResetView()
}

func (cp *speedoPaint) drawIndicator(p *wolffCanvas.Context) {
	lg := wolffCanvas.NewLinearGradient(
		wolffCanvas.Point{X: 0, Y: 0},
		wolffCanvas.Point{X: float64(cp.width), Y: float64(cp.height)})
	lg.Add(0, wolffCanvas.Hex("#3c1c00"))
	lg.Add(1, wolffCanvas.Hex("#a05000"))
	p.SetFillGradient(lg)

	p.SetStrokeWidth(1)
	p.SetStrokeColor(wolffCanvas.Red)
	path, _ := wolffCanvas.ParseSVGPath("M -6 0 L 0 70 L 6 0 Z")
	p.DrawPath(0, 0, path)
}

func (cp *speedoPaint) drawLayer2Capuchon() {
	p := cp.pincel
	//p.SetZIndex(2)

	p.SetFillColor(wolffCanvas.Transparent)
	p.Translate(float64(cp.width/2), float64(cp.height/2))
	p.Scale(float64(cp.radius/100), float64(cp.radius/100))
	// Aguja - Capuchon
	rg := wolffCanvas.NewRadialGradient(wolffCanvas.Point{X: float64(cp.width) / 2, Y: float64(cp.height) / 2}, 1.0,
		wolffCanvas.Point{X: float64(cp.width) / 2, Y: float64(cp.height) / 2}, 25.0)
	rg.Add(0.1, wolffCanvas.Hex("#777777"))
	rg.Add(0.50, wolffCanvas.Hex("#aaaaaa"))
	rg.Add(0.85, wolffCanvas.Hex("#eeeeee"))
	rg.Add(0.95, wolffCanvas.Hex("#f0f0f0"))
	p.SetFillGradient(rg)
	p.SetStroke(wolffCanvas.Hex("#cccccc"))
	p.DrawPath(0, 0, wolffCanvas.Circle(12))

	cp.radius = (cp.side - float64(cp.margins)) / 2.0
	p.ResetStyle()
	p.ResetView()
}

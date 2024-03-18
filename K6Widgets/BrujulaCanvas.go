package K6Widgets

import (
	"math"

	wolffCanvas "github.com/tdewolff/canvas"
)

type brujulaPaint struct {
	k6paint

	side, radius float64
	margins      int
}

func CreateBrujulaPaint(w, h float64) *brujulaPaint {
	aux := brujulaPaint{}
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

func (cp *brujulaPaint) Refresh(or int) {
	cp.Reset()

	cp.drawLayer0Esferas()
	cp.drawLayer1Flecha(or)
	cp.drawLayer2Capuchon()
}

func (cp *brujulaPaint) drawLayer0Esferas() {
	p := cp.pincel
	//p.SetZIndex(0)

	// rectangulo inicial
	p.DrawPath(0, 0, wolffCanvas.Rectangle(float64(cp.width), float64(cp.height)))
	p.Translate(float64(cp.width/2), float64(cp.height/2))

	// Aros
	lg := wolffCanvas.NewLinearGradient(wolffCanvas.Point{X: 0, Y: 0},
		wolffCanvas.Point{X: float64(cp.width), Y: float64(cp.height)})
	lg.Add(0.1, wolffCanvas.Hex("#444444"))
	lg.Add(0.4, wolffCanvas.Hex("#cccccc"))
	lg.Add(0.6, wolffCanvas.Hex("#f8f8f8"))
	lg.Add(0.9, wolffCanvas.Hex("#ffffff"))
	p.SetFillGradient(lg)
	p.DrawPath(0, 0, wolffCanvas.Circle(float64(cp.radius)))
	p.ResetStyle()
	cp.radius -= 0.13 * cp.radius
	p.DrawPath(0, 0, wolffCanvas.Circle(float64(cp.radius)))

	p.Scale(float64(cp.radius/100), float64(cp.radius/100))

	// Marcas y Letras
	p.SetStroke(wolffCanvas.White)
	p.SetStrokeWidth(2)

	monoSpace := wolffCanvas.NewFontFamily("Monospace")
	if err := monoSpace.LoadLocalFont("DroidSansMono Bold", wolffCanvas.FontBold); err != nil {
		panic(err)
	}
	TamFont := 48 * cp.radius / 100
	face := monoSpace.Face(TamFont, wolffCanvas.White, wolffCanvas.FontBold)
	pointText := map[int]string{0: "N", 45: "NE", 90: "E", 135: "SE", 180: "S", 225: "SW", 270: "W", 315: "NW"}

	for i := 0; i < 360; i += 15 {
		if i%45 == 0 {
			// Rayita gorda
			path, _ := wolffCanvas.ParseSVGPath("M 0 75 L 0 90")
			p.DrawPath(0, 0, path)

			// Letras (pasan del rotate y del scale, asique con rendertext que tiene matriz trafo)
			a := float64(i - 90)
			b := float64(-i)
			cos := math.Cos(b * math.Pi / 180)
			sin := math.Sin(b * math.Pi / 180)
			Ax := float64(cp.width)/2 + 60*cp.radius/100*math.Cos(a*math.Pi/180)
			Ay := float64(cp.height)/2 + 60*cp.radius/100*math.Sin(-a*math.Pi/180)
			miMatriz := wolffCanvas.Matrix{{cos, -sin, Ax},
				{sin, cos, Ay}}
			p.RenderText(wolffCanvas.NewTextLine(face, pointText[i], wolffCanvas.Middle), miMatriz)
		} else {
			// Rayita corta
			path, _ := wolffCanvas.ParseSVGPath("M 0 85 L 0 90")
			p.DrawPath(0, 0, path)
		}
		p.Rotate(float64(-15))
	}
	p.ResetStyle()
	p.ResetView()
}

func (cp *brujulaPaint) drawLayer1Flecha(alfa int) {
	p := cp.pincel
	//p.SetZIndex(1)

	p.SetFillColor(wolffCanvas.Transparent)
	p.DrawPath(0, 0, wolffCanvas.Rectangle(float64(cp.width), float64(cp.height)))
	p.Translate(float64(cp.width/2), float64(cp.height/2))
	p.Scale(float64(cp.radius/100), float64(cp.radius/100))
	p.Rotate(float64(-alfa))

	// Aguja - Saetas
	lg := wolffCanvas.NewLinearGradient(wolffCanvas.Point{X: 0, Y: 0},
		wolffCanvas.Point{X: float64(cp.width), Y: float64(cp.height)})
	lg.Add(0.2, wolffCanvas.Hex("#FF0000"))
	lg.Add(0.8, wolffCanvas.Hex("#FFFF00"))
	p.SetFillGradient(lg)
	p.SetStroke(wolffCanvas.Hex("#FF0000"))
	path, _ := wolffCanvas.ParseSVGPath("M -7 0 L 0 70 L 7 0 L 0 -70 Z")
	p.DrawPath(0, 0, path)
	//p.Fill()

	// 	Aguja - Punta
	lg2 := wolffCanvas.NewLinearGradient(wolffCanvas.Point{X: 0, Y: 0},
		wolffCanvas.Point{X: float64(cp.width), Y: float64(cp.height)})
	lg2.Add(0.2, wolffCanvas.Hex("#00FF00"))
	lg2.Add(0.8, wolffCanvas.Hex("#FFFF00"))
	p.SetFillGradient(lg2)
	path2, _ := wolffCanvas.ParseSVGPath("M -8 45 L 0 70 L 8 45 L 0 50 Z")
	p.DrawPath(0, 0, path2)

	p.ResetStyle()
	p.ResetView()
}

func (cp *brujulaPaint) drawLayer2Capuchon() {
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

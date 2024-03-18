package K6Widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

type SevenSegWidget struct {
	widget.BaseWidget

	width, height      float32
	theNumber          int
	ndigits            int
	border, bckgTransp bool
	OnTapped           func(pe *fyne.PointEvent)
	OnTappedSecondary  func(pe *fyne.PointEvent)	
	
	binder basicBinder
}

func NewSevenSegWidget(w, h float32, ndig int, funcs ...func(pe *fyne.PointEvent)) *SevenSegWidget {
	wdg := &SevenSegWidget{}
	wdg.width = h * 0.5 * float32(ndig) //- h*0.18
	wdg.height = h
	wdg.theNumber = 0
	wdg.ndigits = ndig
	wdg.border = false
	wdg.bckgTransp = false

	switch len(funcs) {
	case 1:
		wdg.OnTapped = funcs[0]
	case 2:
		wdg.OnTapped = funcs[0]
		wdg.OnTappedSecondary = funcs[1]
	default:
		wdg.OnTapped = func(pe *fyne.PointEvent) { wdg.IncreasTheNumber(2) }
		wdg.OnTappedSecondary = func(pe *fyne.PointEvent) { wdg.IncreasTheNumber(-2) }
	}

	wdg.ExtendBaseWidget(wdg)
	return wdg
}

func NewSevenSegWidgetFull(w, h float32, ndig int, br, bckgT bool, funcs ...func(pe *fyne.PointEvent)) *SevenSegWidget {
	var wdg *SevenSegWidget
	switch len(funcs) {
	case 1:
		wdg = NewSevenSegWidget(w, h, ndig, funcs[0])
	case 2:
		wdg = NewSevenSegWidget(w, h, ndig, funcs[0], funcs[1])
	default:
		wdg = NewSevenSegWidget(w, h, ndig)
	}
	wdg.border = br
	wdg.bckgTransp = bckgT
	return wdg
}

func NewSevenSegWidgetWithData(w, h float32, ndig int, br, bckgT bool, data binding.Int, funcs ...func(pe *fyne.PointEvent)) *SevenSegWidget {
	var wdg *SevenSegWidget
	switch len(funcs) {
	case 1:
		wdg = NewSevenSegWidgetFull(w, h, ndig, br, bckgT, funcs[0])
	case 2:
		wdg = NewSevenSegWidgetFull(w, h, ndig, br, bckgT, funcs[0], funcs[1])
	default:
		wdg = NewSevenSegWidgetFull(w, h, ndig, br, bckgT,)
	}

	wdg.binder.SetCallback(wdg.updateFromData) // This could only be done once, maybe in ExtendBaseWidget?
	wdg.binder.Bind(data)

	return wdg
}

func (wdg *SevenSegWidget) updateFromData(data binding.DataItem) {
	if data == nil {
		return
	}
	intSource, ok := data.(binding.Int)
	if !ok {
		return
	}
	val, err := intSource.Get()
	if err != nil {
		fyne.LogError("Error getting current data value", err)
		return
	}
	wdg.SetTheNumber(val)
}

func (wdg *SevenSegWidget) SetTheNumber(num int) {
	wdg.theNumber = num
	wdg.Refresh()
}

func (wdg *SevenSegWidget) IncreasTheNumber(inc int) {
	wdg.theNumber += inc
	wdg.Refresh()
}

func (wdg *SevenSegWidget) Tapped(pe *fyne.PointEvent) {
	wdg.OnTapped(pe)
}

func (wdg *SevenSegWidget) TappedSecondary(pe *fyne.PointEvent) {
	wdg.OnTappedSecondary(pe)
}

func (wdg *SevenSegWidget) MinSize() fyne.Size {
	return fyne.NewSize(wdg.width, wdg.height)
}

func (wdg *SevenSegWidget) CreateRenderer() fyne.WidgetRenderer {
	rnd := &sevenSegRenderer{miSevenSeg: wdg, objects: []fyne.CanvasObject{}}

	rnd.miPaint = CreateSevenSegPaint(float64(wdg.width), float64(wdg.height), wdg.ndigits, wdg.border, wdg.bckgTransp)
	rnd.miContainer = container.New(layout.NewMaxLayout())
	rnd.miContainer.Resize(fyne.Size{Width: wdg.width, Height: wdg.height})
	rnd.miContainer.Add(rnd.miPaint.GetPaint())
	rnd.objects = []fyne.CanvasObject{rnd.miContainer}

	return rnd
}

type sevenSegRenderer struct {
	miSevenSeg  *SevenSegWidget
	miPaint     *sevenSegPaint
	miContainer *fyne.Container
	objects     []fyne.CanvasObject
}

// Refresh implements fyne.WidgetRenderer.
func (rnd *sevenSegRenderer) Refresh() {
	rnd.miPaint.Refresh(rnd.miSevenSeg.theNumber)
	rnd.miContainer.Objects[0] = rnd.miPaint.GetPaint()
	rnd.miContainer.Refresh()
}

// BackgroundColor implements fyne.WidgetRenderer.
func (*sevenSegRenderer) BackgroundColor() color.Color {
	return color.Color(color.Transparent)
}

// Destroy implements fyne.WidgetRenderer.
func (*sevenSegRenderer) Destroy() {
}

// Layout implements fyne.WidgetRenderer.
func (rnd *sevenSegRenderer) Layout(fyne.Size) {
	//rnd.Refresh()
}

// MinSize implements fyne.WidgetRenderer.
func (rnd *sevenSegRenderer) MinSize() fyne.Size {
	return fyne.NewSize(rnd.miSevenSeg.width, rnd.miSevenSeg.height)
}

// Objects implements fyne.WidgetRenderer.
func (rnd *sevenSegRenderer) Objects() []fyne.CanvasObject {
	return rnd.objects
}

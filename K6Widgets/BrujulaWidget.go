package K6Widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

type BrujulaWidget struct {
	widget.BaseWidget

	width, height     float32
	orientation       int
	OnTapped          func(pe *fyne.PointEvent)
	OnTappedSecondary func(pe *fyne.PointEvent)

	binder basicBinder
}


func NewBrujulaWidget(w, h float32, funcs ...func(pe *fyne.PointEvent)) *BrujulaWidget {
	wdg := &BrujulaWidget{}
	wdg.width = w
	wdg.height = h
	wdg.orientation = 0

	switch len(funcs) {
	case 1:
		wdg.OnTapped = funcs[0]
	case 2:
		wdg.OnTapped = funcs[0]
		wdg.OnTappedSecondary = funcs[1]
	default:
		wdg.OnTapped = func(pe *fyne.PointEvent) { wdg.IncreasOrientetation(2) }
		wdg.OnTappedSecondary = func(pe *fyne.PointEvent) { wdg.IncreasOrientetation(-2) }
	}

	wdg.ExtendBaseWidget(wdg)

	return wdg
}

func NewBrujulaWidgetWithData(w,h float32, data binding.Int, funcs ...func(pe *fyne.PointEvent)) *BrujulaWidget {
	var wdg *BrujulaWidget

	switch len(funcs) {
	case 1:
		wdg=NewBrujulaWidget(w,h,funcs[0]) 
	case 2:
		wdg=NewBrujulaWidget(w,h,funcs[0],funcs[1]) 
	default:
		wdg=NewBrujulaWidget(w,h) 
	}

	wdg.binder.SetCallback(wdg.updateFromData) // This could only be done once, maybe in ExtendBaseWidget?
	wdg.binder.Bind(data)

	return wdg
}

func (wdg *BrujulaWidget) SetOrientetation(orienttn int) {
	wdg.orientation = orienttn
	wdg.Refresh()
}

func (wdg *BrujulaWidget) IncreasOrientetation(orienttn int) {
	wdg.orientation += orienttn

	if wdg.orientation > 360 {
		wdg.orientation -= 360
	}
	if wdg.orientation < 0 {
		wdg.orientation += 360
	}
	wdg.Refresh()
}

func (wdg *BrujulaWidget) Tapped(pe *fyne.PointEvent) {
	wdg.OnTapped(pe)
}

func (wdg *BrujulaWidget) TappedSecondary(pe *fyne.PointEvent) {
	wdg.OnTappedSecondary(pe)
}

func (wdg *BrujulaWidget) MinSize() fyne.Size {
	return fyne.NewSize(wdg.width, wdg.height)
}

func (wdg *BrujulaWidget) updateFromData(data binding.DataItem) {
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
	wdg.SetOrientetation(val)
}

func (wdg *BrujulaWidget) CreateRenderer() fyne.WidgetRenderer {
	rnd := &brujulaRenderer{miBrujula: wdg, objects: []fyne.CanvasObject{}}

	rnd.miPaint = CreateBrujulaPaint(float64(wdg.width), float64(wdg.height))
	rnd.miContainer = container.New(layout.NewMaxLayout())
	rnd.miContainer.Resize(fyne.Size{Width: wdg.width, Height: wdg.height})
	rnd.miContainer.Add(rnd.miPaint.GetPaint())
	rnd.objects = []fyne.CanvasObject{rnd.miContainer}

	return rnd
}

type brujulaRenderer struct {
	miBrujula   *BrujulaWidget
	miPaint     *brujulaPaint
	miContainer *fyne.Container
	objects     []fyne.CanvasObject
}

// Refresh implements fyne.WidgetRenderer.
func (rnd *brujulaRenderer) Refresh() {
	rnd.miPaint.Refresh(rnd.miBrujula.orientation)

	rnd.miContainer.Objects[0] = rnd.miPaint.GetPaint()
	rnd.miContainer.Refresh()
}

// BackgroundColor implements fyne.WidgetRenderer.
func (*brujulaRenderer) BackgroundColor() color.Color {
	return color.Color(color.Transparent)
}

// Destroy implements fyne.WidgetRenderer.
func (*brujulaRenderer) Destroy() {
}

// Layout implements fyne.WidgetRenderer.
func (rnd *brujulaRenderer) Layout(fyne.Size) {
	//rnd.Refresh()
}

// MinSize implements fyne.WidgetRenderer.
func (rnd *brujulaRenderer) MinSize() fyne.Size {
	return fyne.NewSize(rnd.miBrujula.width, rnd.miBrujula.height)
}

// Objects implements fyne.WidgetRenderer.
func (rnd *brujulaRenderer) Objects() []fyne.CanvasObject {
	return rnd.objects
}

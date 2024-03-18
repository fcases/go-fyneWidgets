package K6Widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

type SpeedoWidget struct {
	widget.BaseWidget

	width, height     float32
	orientation       int
	OnTapped          func(pe *fyne.PointEvent)
	OnTappedSecondary func(pe *fyne.PointEvent)

	binder basicBinder
}

func NewSpeedoWidget(w, h float32, funcs ...func(pe *fyne.PointEvent)) *SpeedoWidget {
	wdg := &SpeedoWidget{}
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

func NewSpeedoWidgetWithData(w, h float32, data binding.Int, funcs ...func(pe *fyne.PointEvent)) *SpeedoWidget {
	var wdg *SpeedoWidget
	switch len(funcs) {
	case 1:
		wdg = NewSpeedoWidget(w, h, funcs[0])
	case 2:
		wdg = NewSpeedoWidget(w, h, funcs[0], funcs[1])
	default:
		wdg = NewSpeedoWidget(w, h,)
	}

	wdg.binder.SetCallback(wdg.updateFromData) // This could only be done once, maybe in ExtendBaseWidget?
	wdg.binder.Bind(data)

	return wdg
}

func (wdg *SpeedoWidget) updateFromData(data binding.DataItem) {
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

func (wdg *SpeedoWidget) SetOrientetation(orienttn int) {
	wdg.orientation = orienttn
	if wdg.orientation > 100 {
		wdg.orientation = wdg.orientation  % 100
	}
	if wdg.orientation < 0 {
		wdg.orientation = 0
	}

	wdg.Refresh()
}

func (wdg *SpeedoWidget) IncreasOrientetation(orienttn int) {
	wdg.SetOrientetation(wdg.orientation+orienttn)
}

func (wdg *SpeedoWidget) Tapped(pe *fyne.PointEvent) {
	wdg.OnTapped(pe)
}

func (wdg *SpeedoWidget) TappedSecondary(pe *fyne.PointEvent) {
	wdg.OnTappedSecondary(pe)
}

func (wdg *SpeedoWidget) MinSize() fyne.Size {
	return fyne.NewSize(wdg.width, wdg.height)
}

func (wdg *SpeedoWidget) CreateRenderer() fyne.WidgetRenderer {
	rnd := &speedoRenderer{miSpeedo: wdg, objects: []fyne.CanvasObject{}}

	rnd.miPaint = CreateSpeedoPaint(float64(wdg.width), float64(wdg.height))
	rnd.miContainer = container.New(layout.NewMaxLayout())
	rnd.miContainer.Resize(fyne.Size{Width: wdg.width, Height: wdg.height})
	rnd.miContainer.Add(rnd.miPaint.GetPaint())

	rnd.mi7S = NewSevenSegWidgetFull(0.13*wdg.height, 0.13*wdg.height, 3, true, true)
	rnd.miContainer.Add(container.New(&speedoLayout{}, rnd.mi7S))

	rnd.objects = []fyne.CanvasObject{rnd.miContainer}

	return rnd
}

type speedoRenderer struct {
	miSpeedo    *SpeedoWidget
	miPaint     *speedoPaint
	miContainer *fyne.Container
	mi7S        *SevenSegWidget
	objects     []fyne.CanvasObject
}

// Refresh implements fyne.WidgetRenderer.
func (rnd *speedoRenderer) Refresh() {
	rnd.mi7S.SetTheNumber(rnd.miSpeedo.orientation)

	rnd.miPaint.Refresh(rnd.miSpeedo.orientation)
	rnd.miContainer.Objects[0] = rnd.miPaint.GetPaint()

	rnd.miContainer.Refresh()
}

// BackgroundColor implements fyne.WidgetRenderer.
func (*speedoRenderer) BackgroundColor() color.Color {
	return color.Color(color.Transparent)
}

// Destroy implements fyne.WidgetRenderer.
func (*speedoRenderer) Destroy() {
}

// Layout implements fyne.WidgetRenderer.
func (rnd *speedoRenderer) Layout(fyne.Size) {
	//rnd.Refresh()
}

// MinSize implements fyne.WidgetRenderer.
func (rnd *speedoRenderer) MinSize() fyne.Size {
	return fyne.NewSize(rnd.miSpeedo.width, rnd.miSpeedo.height)
}

// Objects implements fyne.WidgetRenderer.
func (rnd *speedoRenderer) Objects() []fyne.CanvasObject {
	return rnd.objects
}

type speedoLayout struct {
}

func (sl *speedoLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) > 0 {
		return objects[0].MinSize()
	} else {
		return fyne.Size{Width: 0, Height: 0}
	}
}

func (sl *speedoLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) > 0 {
		objects[0].Move(fyne.Position{
			X: containerSize.Width/2 - objects[0].MinSize().Width/2,
			Y: 0.74 * containerSize.Height})
	}
}

package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"

	myWidgets "p9fynesvg/K6Widgets"
)

func main() {
	a := app.New()

	// Create window named compass widget (in semi-spannish, brujula=compass)
	w := a.NewWindow("K6 Widgets v2")

	/*
	// Create the widgets
	myB := myWidgets.NewBrujulaWidget(370, 370)
	myS1 := myWidgets.NewSevenSegWidget(0, 60, 3)
	myS2 := myWidgets.NewSevenSegWidgetFull(0, 60, 3, true, true)
	myS3 := myWidgets.NewSevenSegWidgetFull(0, 40, 3, false, true)
	myS4 := myWidgets.NewSevenSegWidgetFull(0, 40, 3, true, false)
	mySW := myWidgets.NewSpeedoWidget(370, 370)
	*/
	// Create the widgets with Binding.Int	
	myBI:=binding.NewInt(); myBI.Set(0)
	myB := myWidgets.NewBrujulaWidgetWithData(370, 370,myBI)
	myS1 := myWidgets.NewSevenSegWidgetWithData(0, 60, 3,false,false,myBI)
	myS2 := myWidgets.NewSevenSegWidgetWithData(0, 60, 3, true, true,myBI)
	myS3 := myWidgets.NewSevenSegWidgetWithData(0, 40, 3, false, true,myBI)
	myS4 := myWidgets.NewSevenSegWidgetWithData(0, 40, 3, true, false,myBI)
	mySW := myWidgets.NewSpeedoWidgetWithData(370, 370,myBI)

	// set them on the window
	cont1 := container.NewCenter(container.NewVBox(
		container.NewCenter(myS2),
		container.NewCenter(myS4),
		container.NewCenter(myS1),
		container.NewCenter(myS3),
	))
	cont2 := container.NewCenter(container.NewHBox(
		myB, cont1, mySW,
	))
	w.SetContent(cont2)

	// 'any key' sets 0, press 'space' to stop the app
	w.Canvas().SetOnTypedKey(
		func(ke *fyne.KeyEvent) {
			if ke.Name == fyne.KeySpace {
				a.Quit()
			} else {
				/*
				myB.SetOrientetation(0)
				myS1.SetTheNumber(0)
				myS2.SetTheNumber(0)
				myS3.SetTheNumber(0)
				myS4.SetTheNumber(0)
				mySW.SetOrientetation(0)
				*/
				myBI.Set(0)
			}
		},
	)

	//go actualizar(myB, myS1, myS2, myS3, myS4, mySW)
	go actualizar(myBI)

	// .. and run
	w.Show()
	a.Run()
}
/*
func actualizar(wdg1 *myWidgets.BrujulaWidget, wdg2 *myWidgets.SevenSegWidget,
	wdg3 *myWidgets.SevenSegWidget, wdg4 *myWidgets.SevenSegWidget,
	wdg5 *myWidgets.SevenSegWidget, wdg6 *myWidgets.SpeedoWidget) {
*/
func actualizar(myBI binding.Int) {
	for i := 0; i < 360; i++ {
		time.Sleep(time.Millisecond * 150)
		myBI.Set(i)
//		wdg1.SetOrientetation(i)
//		wdg2.SetTheNumber(i)
//		wdg3.SetTheNumber(i)
//		wdg4.SetTheNumber(i)
//		wdg5.SetTheNumber(i)
//		wdg6.SetOrientetation(i % 100)
	}
}

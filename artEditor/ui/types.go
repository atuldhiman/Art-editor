package ui

import (
	"creative.io/artEditor/apptype"
	"creative.io/artEditor/pxcanvas"
	"creative.io/artEditor/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}

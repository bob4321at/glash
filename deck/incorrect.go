package deck

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func IncorrectUi() *fyne.Container {
	IncorrectText := canvas.NewText("Incorrect", color.Black)
	IncorrectText.Alignment = fyne.TextAlignCenter
	IncorrectText.TextSize = 64

	IncorrectUi := container.NewGridWithColumns(1,
		IncorrectText,
		widget.NewButton("Next", func() {
			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
	)

	return IncorrectUi
}

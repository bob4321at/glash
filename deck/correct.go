package deck

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CorrectUi() *fyne.Container {
	CorrectText := canvas.NewText("Correct", color.Black)
	CorrectText.Alignment = fyne.TextAlignCenter
	CorrectText.TextSize = 64

	CorrectUi := container.NewGridWithColumns(1,
		CorrectText,
		widget.NewButton("Next", func() {
			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
	)

	return CorrectUi
}

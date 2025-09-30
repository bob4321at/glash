package deck

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CorrectUi(Back string) *fyne.Container {
	CorrectText := canvas.NewText("Correct", color.Black)
	CorrectText.Alignment = fyne.TextAlignCenter
	CorrectText.TextSize = 128
	Back_Ui := canvas.NewText(Back, color.Black)
	Back_Ui.Alignment = fyne.TextAlignCenter
	Back_Ui.TextSize = 96

	CorrectUi := container.NewGridWithColumns(1,
		widget.NewButton("Next", func() {
			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
		CorrectText,
		Back_Ui,
	)

	return CorrectUi
}

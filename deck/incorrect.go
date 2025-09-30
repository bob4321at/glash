package deck

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func IncorrectUi(Back string) *fyne.Container {
	Back_Ui := canvas.NewText(Back, color.Black)
	Back_Ui.Alignment = fyne.TextAlignCenter
	Back_Ui.TextSize = 96
	IncorrectText := canvas.NewText("Incorrect", color.Black)
	IncorrectText.Alignment = fyne.TextAlignCenter
	IncorrectText.TextSize = 64

	IncorrectUi := container.NewGridWithColumns(1,
		widget.NewButton("Next", func() {
			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
		IncorrectText,
		Back_Ui,
	)

	return IncorrectUi
}

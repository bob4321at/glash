package deck

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Right_Answers int
var Total_Answers int

func EndCardUi() fyne.CanvasObject {
	End_Text := canvas.NewText("End Scores", color.Black)
	End_Text.Alignment = fyne.TextAlignCenter
	End_Text.TextSize = 32

	circle := canvas.NewImageFromFile("./circle.png")

	card_canvas := canvas.NewRectangle(color.RGBA{230, 230, 230, 255})
	card_canvas.CornerRadius = 20

	Score_Text := canvas.NewText(strconv.Itoa(Right_Answers)+" out of "+strconv.Itoa(Total_Answers), color.Black)
	Score_Text.Alignment = fyne.TextAlignCenter
	Score_Text.TextSize = 128

	EndCardUi := container.NewBorder(
		End_Text,
		widget.NewButton("Exit", func() {
			Window_Ref.SetContent(CardUi(Selected_Deck.Name))
		}),
		nil,
		nil,
		container.NewGridWithRows(1,
			container.NewStack(
				card_canvas,
				circle,
				Score_Text,
			),
		),
	)

	fmt.Println(Right_Answers, Total_Answers)

	return EndCardUi
}

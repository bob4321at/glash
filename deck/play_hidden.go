package deck

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PlayHiddenUi() *fyne.Container {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	last_index := r.Intn(len(Selected_Deck.Cards))
	QuestionWordLabel := canvas.NewText(Selected_Deck.Cards[last_index].Word, color.Black)
	QuestionWordLabel.TextSize = 100
	QuestionWordLabel.Alignment = fyne.TextAlignCenter

	QuestionUi := container.NewGridWithColumns(1,
		QuestionWordLabel,
	)

	var button_ui_ref *widget.Button

	button_ui := widget.NewButton("flip", func() {
		QuestionWordLabel.Text = Selected_Deck.Cards[last_index].Back + " / " + Selected_Deck.Cards[last_index].Answer
		button_ui_ref.Text = "Continue"
		button_ui_ref.OnTapped = func() {
			Window_Ref.SetContent(PlayHiddenUi())
		}
	})

	button_ui_ref = button_ui

	PlayUi := container.NewBorder(
		container.NewVBox(widget.NewButton("Done", func() {
			Window_Ref.SetContent(
				CardUi(Selected_Deck.Name),
			)
		})),
		button_ui,
		nil,
		nil,
		container.NewGridWithColumns(1,
			QuestionUi,
		),
	)
	return PlayUi
}

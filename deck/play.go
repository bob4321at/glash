package deck

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetAllAnswers(deck Deck) []string {
	answers := []string{}

	for _, card := range deck.Cards {
		answers = append(answers, card.Answer)
	}

	return answers
}

func AnswerButton(Answer string, card_index int) {
	if Selected_Deck.Cards[card_index].Answer == Answer {
		Window_Ref.SetContent(
			CorrectUi(),
		)
	} else {
		Window_Ref.SetContent(
			IncorrectUi(),
		)
	}
}

func PlayUi() *fyne.Container {
	answers := GetAllAnswers(*Selected_Deck)

	last_index := int(rand.Float64() * float64(len(Selected_Deck.Cards)))
	QuestionWordLabel := canvas.NewText(Selected_Deck.Cards[last_index].Word, color.Black)
	QuestionWordLabel.TextSize = 100
	QuestionWordLabel.Alignment = fyne.TextAlignCenter

	fmt.Print(last_index)

	QuestionUi := container.NewGridWithColumns(1,
		QuestionWordLabel,
	)

	Answer := [4]string{}

	for i := range 4 {
		Answer[i] = answers[int(rand.Float64()*float64(len(answers)))]
	}

	Answer[int(rand.Float64()*float64(len(Answer)))] = Selected_Deck.Cards[last_index].Answer

	PlayUi := container.NewBorder(
		container.NewVBox(widget.NewButton("exit", func() {
			Window_Ref.SetContent(
				CardUi(Selected_Deck.Name),
			)
		})),
		nil,
		nil,
		nil,
		container.NewGridWithColumns(1,
			QuestionUi,

			container.NewGridWithColumns(2,
				container.NewGridWithRows(2,
					widget.NewButton(Answer[0], func() {
						AnswerButton(Answer[0], last_index)
					}),
					widget.NewButton(Answer[1], func() {
						AnswerButton(Answer[1], last_index)
					}),
				),
				container.NewGridWithRows(2,
					widget.NewButton(Answer[2], func() {
						AnswerButton(Answer[2], last_index)
					}),
					widget.NewButton(Answer[3], func() {
						AnswerButton(Answer[3], last_index)
					}),
				),
			),
		),
	)
	return PlayUi
}

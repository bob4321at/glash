package deck

import (
	"fmt"
	"image/color"
	"main/utils"
	"math/rand"
	"time"

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
	Total_Answers += 1
	if Selected_Deck.Cards[card_index].Answer == Answer {
		Right_Answers += 1

		var selected_card_ptr *Card
		var selected_card_difficulty int
		var selected_card_index int

		for difficulty := range Priorities {
			for index := range Priorities[difficulty] {
				card := Priorities[difficulty][index]
				if Priorities[difficulty][index].ID == card_index {
					selected_card_ptr = card
					selected_card_difficulty = difficulty
					selected_card_index = index
				}
			}
		}

		if selected_card_difficulty < 2 {
			utils.RemoveArrayElement(selected_card_index, &Priorities[selected_card_difficulty])
			Priorities[selected_card_difficulty+1] = append(Priorities[selected_card_difficulty+1], selected_card_ptr)
		}

		Window_Ref.SetContent(
			CorrectUi(Selected_Deck.Cards[card_index].Back),
		)
	} else {
		Window_Ref.SetContent(
			IncorrectUi(Selected_Deck.Cards[card_index].Back),
		)
	}
}

var Priorities = [3][]*Card{}

func PlayUi() *fyne.Container {
	answers := GetAllAnswers(*Selected_Deck)

	fmt.Println(Priorities)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var chosen_card *Card

	for chosen_card == nil {
		difficulty := r.Float64()

		if difficulty < 0.55 {
			if len(Priorities[0]) != 0 {
				chosen_card = Priorities[0][rand.Intn(len(Priorities[0]))]
			}
		} else if difficulty > 0.8 {
			if len(Priorities[2]) != 0 {
				chosen_card = Priorities[2][rand.Intn(len(Priorities[2]))]
			}
		} else {
			if len(Priorities[1]) != 0 {
				chosen_card = Priorities[1][rand.Intn(len(Priorities[1]))]
			}
		}
	}

	QuestionWordLabel := canvas.NewText(chosen_card.Word, color.Black)
	QuestionWordLabel.TextSize = 100
	QuestionWordLabel.Alignment = fyne.TextAlignCenter

	QuestionUi := container.NewGridWithColumns(1,
		QuestionWordLabel,
	)

	Answer := [4]string{}

	for i := range 4 {
		Answer[i] = answers[rand.Intn(len(answers))]
	}

	Answer[r.Intn(len(Answer))] = chosen_card.Answer

	PlayUi := container.NewBorder(
		container.NewVBox(widget.NewButton("Done", func() {
			Window_Ref.SetContent(
				EndCardUi(),
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
						AnswerButton(Answer[0], chosen_card.ID)
					}),
					widget.NewButton(Answer[1], func() {
						AnswerButton(Answer[1], chosen_card.ID)
					}),
				),
				container.NewGridWithRows(2,
					widget.NewButton(Answer[2], func() {
						AnswerButton(Answer[2], chosen_card.ID)
					}),
					widget.NewButton(Answer[3], func() {
						AnswerButton(Answer[3], chosen_card.ID)
					}),
				),
			),
		),
	)
	return PlayUi
}

package deck

import (
	"fmt"
	"main/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Card struct {
	Word   string
	Answer string
	ID     int
}

func NewCard(Word, Answer string, ID int) Card {
	card := Card{
		Word:   Word,
		Answer: Answer,
		ID:     ID,
	}

	return card
}

type CardWidget struct {
	widget.BaseWidget
	WordUi     *widget.Entry
	AnswerUi   *widget.Entry
	SaveButton *widget.Button
	RemoveCard *widget.Button

	Word   string
	Answer string

	OnTapped func()
}

type CardWidgetRenderer struct {
	icon   *canvas.Image
	label  *canvas.Text
	shadow *fyne.CanvasObject

	objects []fyne.CanvasObject
	card    *CardWidget
}

func NewMyListItemWidget(Word, Answer string) *CardWidget {
	var Temp_Word_Ui_Ref *string
	var Temp_Answer_Ui_Ref *string
	var Before_Temp_Word_Ui_Ref *string
	var Before_Temp_Answer_Ui_Ref *string

	item := &CardWidget{
		WordUi:   widget.NewEntry(),
		AnswerUi: widget.NewEntry(),
		SaveButton: widget.NewButton("Save", func() {
			for i := range Selected_Deck.Cards {
				card := &Selected_Deck.Cards[i]
				fmt.Println(Current_Card_ID)
				if card.ID == Current_Card_ID {
					if *Temp_Word_Ui_Ref == "" {
						Temp_Word_Ui_Ref = Before_Temp_Word_Ui_Ref
					}
					if *Temp_Answer_Ui_Ref == "" {
						Temp_Answer_Ui_Ref = Before_Temp_Answer_Ui_Ref
					}

					card.Word = *Temp_Word_Ui_Ref
					card.Answer = *Temp_Answer_Ui_Ref
				}
			}
		}),
		RemoveCard: widget.NewButton("Remove", func() {
			for i := range Selected_Deck.Cards {
				card := &Selected_Deck.Cards[i]
				if card.Answer == Answer && card.Word == Word {
					utils.RemoveArrayElement(i, &Selected_Deck.Cards)
					return
				}
			}
		}),
	}

	Temp_Word_Ui_Ref = &item.WordUi.Text
	Temp_Answer_Ui_Ref = &item.AnswerUi.Text
	Before_Temp_Word_Ui_Ref = &item.WordUi.PlaceHolder
	Before_Temp_Answer_Ui_Ref = &item.AnswerUi.PlaceHolder

	item.WordUi.PlaceHolder = Word
	item.AnswerUi.PlaceHolder = Answer

	item.ExtendBaseWidget(item)

	return item
}

func (item *CardWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(item.WordUi, item.AnswerUi, item.SaveButton, item.RemoveCard)
	return widget.NewSimpleRenderer(c)
}

var Current_Card_ID int

func CardUi(name string) *fyne.Container {
	var Selected_Card *Card

	Current_Card_Widget := container.NewVBox(widget.NewLabel(""))

	Cards_To_Render := widget.NewList(
		func() int {
			return len(Selected_Deck.Cards)
		}, func() fyne.CanvasObject {
			return widget.NewLabel("")
		}, func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(Selected_Deck.Cards[i].Word)
		})
	Cards_To_Render.OnSelected = func(id int) {
		Selected_Card = &Selected_Deck.Cards[id]
		Current_Card_ID = Selected_Card.ID
		Current_Card_Widget.RemoveAll()
		Current_Card_Widget.Add(NewMyListItemWidget(Selected_Card.Word, Selected_Card.Answer))
		Current_Card_Widget.Refresh()
	}
	Cards_To_Render_Container := container.NewScroll(
		Cards_To_Render,
	)
	Cards_To_Render_Container.SetMinSize(fyne.NewSize(200, 600))

	New_Card_Name := widget.NewEntry()
	New_Card_Answer := widget.NewEntry()

	BottomUi := container.NewVBox(
		New_Card_Name,
		New_Card_Answer,
		widget.NewButton("Add Card", func() {
			if New_Card_Name != nil && New_Card_Answer != nil {
				Selected_Deck.Cards = append(Selected_Deck.Cards, NewCard(New_Card_Name.Text, New_Card_Answer.Text, len(Selected_Deck.Cards)))
				New_Card_Name.Text = ""
				New_Card_Answer.Text = ""
				New_Card_Name.Refresh()
				New_Card_Answer.Refresh()
				Cards_To_Render.Refresh()
			}
		}),

		widget.NewButton("Leave", func() {
			Window_Ref.SetContent(DeckUi())
		}),
	)

	Left_Side_Ui := container.NewVBox(
		container.NewVBox(
			widget.NewLabel(name),
			Cards_To_Render_Container,
		),
		widget.NewButton("Save Deck", func() {
			Selected_Deck.Serialize(Selected_Deck.Name + ".fcard")
			Cards_To_Render.Refresh()
		}),
		BottomUi,
		widget.NewButton("Play", func() {
			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
	)

	Right_Side_Ui := container.NewVBox(
		Current_Card_Widget,
	)

	CardUi := container.NewGridWithColumns(
		2,
		Left_Side_Ui,
		Right_Side_Ui,
	)

	return CardUi
}

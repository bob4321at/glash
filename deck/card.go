package deck

import (
	"main/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Card struct {
	Word   string
	Answer string
	Back   string
	ID     int
}

func NewCard(Word, Answer string, Back string, ID int) Card {
	card := Card{
		Word:   Word,
		Answer: Answer,
		Back:   Back,
		ID:     ID,
	}

	return card
}

type CardWidget struct {
	widget.BaseWidget
	WordUi     *widget.Entry
	AnswerUi   *widget.Entry
	BackUi     *widget.Entry
	SaveButton *widget.Button
	RemoveCard *widget.Button

	Word   string
	Answer string
	Back   string

	OnTapped func()
}

type CardWidgetRenderer struct {
	icon   *canvas.Image
	label  *canvas.Text
	shadow *fyne.CanvasObject

	objects []fyne.CanvasObject
	card    *CardWidget
}

func NewMyListItemWidget(Word, Answer, Back string) *CardWidget {
	var Temp_Word_Ui_Ref *string
	var Temp_Answer_Ui_Ref *string
	var Temp_Back_Ui_Ref *string
	var Before_Temp_Word_Ui_Ref *string
	var Before_Temp_Answer_Ui_Ref *string
	var Before_Temp_Back_Ui_Ref *string

	item := &CardWidget{
		WordUi:   widget.NewEntry(),
		AnswerUi: widget.NewEntry(),
		BackUi:   widget.NewEntry(),
		SaveButton: widget.NewButton("Save", func() {
			for i := range Selected_Deck.Cards {
				card := &Selected_Deck.Cards[i]
				if card.ID == Current_Card_ID {
					if *Temp_Word_Ui_Ref == "" {
						Temp_Word_Ui_Ref = Before_Temp_Word_Ui_Ref
					}
					if *Temp_Answer_Ui_Ref == "" {
						Temp_Answer_Ui_Ref = Before_Temp_Answer_Ui_Ref
					}
					if *Temp_Back_Ui_Ref == "" {
						Temp_Back_Ui_Ref = Before_Temp_Back_Ui_Ref
					}

					card.Word = *Temp_Word_Ui_Ref
					card.Answer = *Temp_Answer_Ui_Ref
					card.Back = *Temp_Back_Ui_Ref
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
			for i := range Selected_Deck.Cards {
				card := &Selected_Deck.Cards[i]
				card.ID = i
			}
		}),
	}

	Temp_Word_Ui_Ref = &item.WordUi.Text
	Temp_Answer_Ui_Ref = &item.AnswerUi.Text
	Temp_Back_Ui_Ref = &item.BackUi.Text
	Before_Temp_Word_Ui_Ref = &item.WordUi.PlaceHolder
	Before_Temp_Answer_Ui_Ref = &item.AnswerUi.PlaceHolder
	Before_Temp_Back_Ui_Ref = &item.BackUi.PlaceHolder

	item.WordUi.PlaceHolder = Word
	item.AnswerUi.PlaceHolder = Answer
	item.BackUi.PlaceHolder = Back

	item.ExtendBaseWidget(item)

	return item
}

func (item *CardWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewVBox(item.WordUi, item.BackUi, item.AnswerUi, item.SaveButton, item.RemoveCard)
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
		Current_Card_Widget.Add(NewMyListItemWidget(Selected_Card.Word, Selected_Card.Answer, Selected_Card.Back))
		Current_Card_Widget.Refresh()
	}
	Cards_To_Render_Container := container.NewScroll(
		Cards_To_Render,
	)
	Cards_To_Render_Container.SetMinSize(fyne.NewSize(200, 600))

	New_Card_Name := widget.NewEntry()
	New_Card_Answer := widget.NewEntry()
	New_Card_Back := widget.NewEntry()

	BottomUi := container.NewVBox(
		New_Card_Name,
		New_Card_Answer,
		New_Card_Back,
		widget.NewButton("Add Card", func() {
			if New_Card_Name != nil && New_Card_Answer != nil {
				Selected_Deck.Cards = append(Selected_Deck.Cards, NewCard(New_Card_Name.Text, New_Card_Answer.Text, New_Card_Back.Text, len(Selected_Deck.Cards)))
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
			for i := range Selected_Deck.Cards {
				card := &Selected_Deck.Cards[i]
				card.ID = i
			}
			Selected_Deck.Serialize(Selected_Deck.Name + ".fcard")
			Cards_To_Render.Refresh()
		}),
		BottomUi,
		widget.NewButton("Play", func() {
			Right_Answers = 0
			Total_Answers = 0

			Priorities = [3][]*Card{}

			for i := range Selected_Deck.Cards {
				Priorities[0] = append(Priorities[0], &Selected_Deck.Cards[i])
			}

			Window_Ref.SetContent(
				PlayUi(),
			)
		}),
		widget.NewButton("Hidden Mode", func() {
			Right_Answers = 0
			Total_Answers = 0
			Window_Ref.SetContent(
				PlayHiddenUi(),
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

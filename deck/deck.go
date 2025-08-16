package deck

import (
	"encoding/json"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Deck struct {
	Name  string
	Cards []Card
}

func (deck *Deck) Serialize(filename string) {
	bytes, err := json.Marshal(deck)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("./decks/" + filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write(bytes)
}

func NewDeck(name string) Deck {
	deck := Deck{}

	deck.Name = name

	return deck
}

var Window_Ref fyne.Window

var Decks []Deck
var Selected_Deck *Deck

var Decks_To_Render = widget.NewList(
	func() int {
		return len(Decks)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("")
	}, func(i int, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(Decks[i].Name)
	})

func DeckUi() *fyne.Container {
	Selected_Deck_Text := widget.NewLabel("")

	Decks_To_Render.OnSelected = func(id int) {
		Selected_Deck = &Decks[id]
		Selected_Deck_Text.SetText(Selected_Deck.Name)
	}
	Decks_To_Render_Container := container.NewScroll(
		Decks_To_Render,
	)

	Decks_To_Render_Container.SetMinSize(fyne.NewSize(200, 400))

	New_Deck_Name := widget.NewEntry()
	New_Deck_Name.SetMinRowsVisible(10)

	DeckUi := container.NewGridWithColumns(
		2,
		Decks_To_Render_Container,
		Selected_Deck_Text,
		container.NewVBox(
			New_Deck_Name,
			widget.NewButton("Make Deck", func() {
				Decks = append(Decks, NewDeck(New_Deck_Name.Text))
				Decks_To_Render_Container.Refresh()
			}),
		),
		widget.NewButton("open", func() {
			Window_Ref.SetContent(
				CardUi(Selected_Deck.Name),
			)
		}),
	)

	return DeckUi
}

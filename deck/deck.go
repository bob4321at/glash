package deck

import (
	"encoding/json"
	"main/utils"
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
	home_dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(deck)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(home_dir + "/Documents/decks/" + filename)
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
				Decks[len(Decks)-1].Serialize(New_Deck_Name.Text + ".fcard")
				Decks_To_Render_Container.Refresh()
			}),
		),
		widget.NewButton("open", func() {
			Window_Ref.SetContent(
				CardUi(Selected_Deck.Name),
			)
		}),
		widget.NewLabel(""),
		widget.NewButton("remove", func() {
			if Selected_Deck != nil {
				home_dir, err := os.UserHomeDir()
				if err != nil {
					panic(err)
				}

				os.Remove(home_dir + "/Documents/decks/" + Selected_Deck.Name + ".fcard")

				for i := range Decks {
					deck := &Decks[i]
					if deck == Selected_Deck {
						utils.RemoveArrayElement(i, &Decks)
						Decks_To_Render_Container.Refresh()
						return
					}
				}
			}
		}),
	)

	return DeckUi
}

package main

import (
	"encoding/json"
	"main/deck"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

var window fyne.Window

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	window = myApp.NewWindow("learn words")
	window.Resize(fyne.NewSize(400, 600))

	Decks_Dir, err := os.ReadDir("./decks")
	if err != nil {
		panic(err)
	}

	for _, file := range Decks_Dir {
		bytes, err := os.ReadFile("./decks/" + file.Name())
		if err != nil {
			panic(err)
		}

		temp_data := deck.Deck{}
		if err := json.Unmarshal(bytes, &temp_data); err != nil {
			panic(err)
		}

		deck.Decks = append(deck.Decks, temp_data)
	}

	deck.Window_Ref = window

	window.SetContent(container.NewVBox(deck.DeckUi()))
	window.ShowAndRun()
}

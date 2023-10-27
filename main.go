package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Globale Variable für die Daten
var data = []string{}

func main() {
	a := app.NewWithID("com.lennart.einkaufszettel")
	w := a.NewWindow("Einkaufszettel")
	w.Resize(fyne.NewSize(300, 600))

	w.SetFixedSize(true)
	w.CenterOnScreen()

	//Entry
	entry := widget.NewEntry()

	//Liste
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle.Italic = true
			o.(*widget.Label).TextStyle.Bold = true
			o.(*widget.Label).SetText(data[i])
		})

	//On select item in list, delete item from list
	list.OnSelected = func(id widget.ListItemID) {
		list.Unselect(id)
		data = append(data[:id], data[id+1:]...)
		list.Refresh()
	}

	//Button
	buttonAdd := widget.NewButton("hinzufügen", func() {
		add(entry.Text)
		entry.SetText("") // Setzt den Entry-Text zurück, nachdem ein Element hinzugefügt wurde
		list.Refresh()    // Aktualisiert die Liste, um das neue Element anzuzeigen
	})
	buttonDel := widget.NewButton("alles entfernen", func() {
		deleteAll()
		list.Refresh() // Aktualisiert die Liste, um alle Elemente zu entfernen
	})

	// container
	vBoxOben := container.NewVBox(entry, buttonAdd)
	borderBox := container.NewBorder(vBoxOben, buttonDel, nil, nil, list)
	w.SetContent(borderBox)
	w.ShowAndRun()

}

// Funktion zum Hinzufügen eines neuen Elements zur Liste
func add(item string) {
	if item != "" { // Überprüft, ob der Eintrag nicht leer ist
		data = append(data, item)
	}
}

// Funktion zum Löschen aller Elemente aus der Liste
func deleteAll() {
	data = []string{}
}

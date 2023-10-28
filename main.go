package main

import (
	"bufio"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
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
			return widget.NewLabelWithStyle("template", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: true})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
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
	buttonSave := widget.NewButton("speichern", func() {
		save(&w)

	})
	buttonLoad := widget.NewButton("laden", func() {
		load(&w)
	})

	buttonAdd.SetIcon(theme.MoveDownIcon())
	buttonDel.SetIcon(theme.DeleteIcon())
	buttonSave.SetIcon(theme.DocumentSaveIcon())
	buttonLoad.SetIcon(theme.FileIcon())

	// container
	vBoxOben := container.NewVBox(entry, buttonAdd)
	hBoxUnten := container.NewHBox(buttonDel, buttonSave, buttonLoad)
	borderBox := container.NewBorder(vBoxOben, hBoxUnten, nil, nil, list)
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
	data = nil
}

func save(w *fyne.Window) {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, *w)
			return
		}
		if writer == nil { // Benutzer hat den Dialog abgebrochen
			return
		}
		defer writer.Close()

		bufWriter := bufio.NewWriter(writer)
		for _, item := range data {
			_, err := bufWriter.WriteString(item + "\n")
			if err != nil {
				dialog.ShowError(err, *w)
				return
			}
		}
		err = bufWriter.Flush()
		if err != nil {
			dialog.ShowError(err, *w)
		}
	}, *w)
	saveDialog.SetFileName("einkaufszettel.txt")
	saveDialog.Show()
}
func load(w *fyne.Window) {
	loadDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, *w)
			return
		}
		if reader == nil { // Benutzer hat den Dialog abgebrochen
			return
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			dialog.ShowError(err, *w)
		}
	}, *w)
	loadDialog.SetFileName("einkaufszettel.txt")
	loadDialog.Show()
}

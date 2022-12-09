package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
)

var windows []*gtk.Window
var windowsCount int

func NewWindow(labelText string) {
    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Panic(err)
    }
    windows = append(windows, win)
    windowsCount++

    win.Connect("destroy", func(w *gtk.Window) {
        w.Destroy()
        windowsCount--

        if windowsCount == 0 { gtk.MainQuit() }
    });

    box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
    if err != nil {
        log.Panic(err)
    }

    label, err := gtk.LabelNew(labelText)
    if err != nil {
        log.Panic(err)
    }
    box.Add(label)

    yesButton, err := gtk.ButtonNewWithLabel("Да")
    if err != nil {
        log.Panic(err)
    }

    yesButton.Connect("clicked", func() {
        windowsCount++
        for _, window := range windows {
            window.Destroy()
        }

        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Panic(err)
        }

        win.Connect("destroy", gtk.MainQuit)

        box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 12)
        if err != nil {
            log.Panic(err)
        }

        label, err := gtk.LabelNew("К сожалению, сейчас у нас нет слона, так что мы не сможем Вам его продать.")
        if err != nil {
            log.Panic(err)
        }
        box.Add(label)

        button, err := gtk.ButtonNewWithLabel("Жаль")
        if err != nil {
            log.Panic(err)
        }

        button.Connect("clicked", gtk.MainQuit)

        box.Add(button)

        win.Add(box)

        win.ShowAll()
    })

    box.Add(yesButton)

    noButton, err := gtk.ButtonNewWithLabel("Нет")
    if err != nil {
        log.Panic(err)
    }

    noButton.Connect("clicked", func() {
        NewWindow("Может, Вы всё-таки хотите ли купить слона?")
    })

    box.Add(noButton)

    win.Add(box)

    win.ShowAll()
}

func main() {
    gtk.Init(nil)

    NewWindow("Хотите ли купить слона?")

    gtk.Main()
}

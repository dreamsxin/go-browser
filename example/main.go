package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dreamsxin/go-browser/frame"
)

func main() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	app := frame.MakeApp("My App")
	app.SetIconFromFile(filepath.Join(dir, "/moon.png"))
	fmt.Println(filepath.Join(dir, "/moon.png"))

	wv := app.NewWindow("Simple program!", 500, 400).
		SetBackgroundColor(50, 50, 50, 0.8).
		Move(20, 100).
		SetMinSize(500, 400).
		SetMaxSize(800, 700).
		// SetDecorated(false).
		LoadHTML(`<body style="color:#dddddd; background: transparent">
      <h1>Hello world</h1>
      <p>Test test test...</p>
      </body>`, "http://localhost:1015/panel/").
		SetStateEvent(func(state frame.State) {
			if state.Hidden {
				fmt.Println("Main window closed")
			}
		}).
		SetInvoke(func(msg string) {
			fmt.Println(":::", msg)
		}).
		Show()

	go func() {
		// wv.Eval("document.querySelector('html').style.background = 'rgba(0,0,0,0.2)';")
		// wv2.Eval("document.querySelector('html').style.background = '#0000aa99';")
		wv.Eval("window.external.invoke('Wow! This is external invoke!')")
		wv.SetTitle("new title")
		wv.Eval("thisIsError2")

		go func() {
			fmt.Println("~~~~~~=========<<<<<< + >>>>>=========~~~~~~~")
			fmt.Println(wv.GetSize())
			fmt.Println(wv.GetInnerSize())
			fmt.Println(wv.GetPosition())
			fmt.Println(wv.GetScreenSize())
			fmt.Println(wv.GetScreenScaleFactor())
			wv.Eval("window.external.invoke('Window 1: This is external invoke')")

		}()
	}()
	// w, h := wv.GetScreen().Size()
	// fmt.Println("Screen size:", w, h)

	app.WaitAllWindowClose()
	// select {}
	fmt.Println("Application terminated")
}

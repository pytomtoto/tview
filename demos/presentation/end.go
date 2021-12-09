package main

import (
	"fmt"

	"github.com/pytomtoto/tcell/v2"
	"github.com/pytomtoto/tview"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/pytomtoto/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}

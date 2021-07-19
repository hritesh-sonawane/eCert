package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)


func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v", w, h)
	pdf.AddPage()

	// Basic Text Stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("Arial", "B", 30)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0) //red
	pdf.Text(0, lineHt, "Hello world!")
	pdf.MoveTo(0, lineHt*2.0)

	pdf.SetFont("Times", "", 15)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*1.5, "Here is some text. If it is too long it will be word wrapped automatically. Hunger is the best spice. If there is a new line it will be\nwrapped as well (unlike other ways of writing text in gofpdf).", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
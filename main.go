package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/signintech/gopdf"
	"log"
	"strconv"
)

//go:embed Ubuntu-L.ttf
var ttf []byte

func main() {
	flag.Parse()
	args := flag.Args()

	numOfPages := 1
	if len(args) > 0 {
		a0 := args[0]
		i, err := strconv.Atoi(a0)
		if err != nil {
			log.Fatal(err)
		}
		numOfPages = i
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if err := pdf.AddTTFFontData("Ubuntu-L", ttf); err != nil {
		log.Fatal(err)
	}

	if err := pdf.SetFont("Ubuntu-L", "", 38); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numOfPages; i++ {
		pdf.AddPage()
		pdf.Cell(nil, fmt.Sprintf("A%d", i+1))
	}

	pdf.WritePdf("sample.pdf")
}

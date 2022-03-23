package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/jroimartin/gocui"
)

// scrapper
// https://github.com/gocolly/colly

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-20, maxY/2, maxX/2+20, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// c := colly.NewCollector()
		c := colly.NewCollector(colly.MaxDepth(2))

		// Find and visit all links
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			err := e.Request.Visit(e.Attr("href"))
			if err != nil {
				return
			}
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		err := c.Visit("https://www.auray.me/")
		if err != nil {
			return err
		}

		_, err2 := fmt.Fprintln(v, "Hello world!")
		if err2 != nil {
			return err2
		}
	}
	return nil
}

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

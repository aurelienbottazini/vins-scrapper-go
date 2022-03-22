package scrapper

import "testing"

type DumbScrapper struct {
}

func (t DumbScrapper) Scrap()               {}
func (t DumbScrapper) CheckForNew()         {}
func (t DumbScrapper) Progress() int        { return 3 }
func (t DumbScrapper) ScrapName() string    { return "dumb" }
func (t DumbScrapper) ScrapState() int      { return Done }
func (t DumbScrapper) ScrapMessage() string { return "dumb message" }

func TestScrapperInterface(t *testing.T) {
	// t.Fatal("not implemented")
	var d Scrapper = DumbScrapper{}
	p := d.Progress()
	if p != 3 {
		t.Errorf("d.Progress() = %d; want 3", p)
	}

	s := d.ScrapState()
	if s != Done {
		t.Errorf("d.ScrapState() = %d; want Done (2)", s)
	}
}

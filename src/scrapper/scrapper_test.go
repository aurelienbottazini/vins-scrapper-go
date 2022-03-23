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
	var d Scrapper = DumbScrapper{}
	p := d.Progress()
	want := 3
	if p != want {
		t.Errorf("d.Progress() = %d; want %d", p, want)
	}

	s := d.ScrapState()
	want = Done
	if s != want {
		t.Errorf("d.ScrapState() = %d; want %d", s, want)
	}
}

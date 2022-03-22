package scrapper

type Status int

const (
	Idle    = iota
	Running = iota
	Done    = iota
)

type Scrapper interface {
	Scrap()
	CheckForNew()
	Progress() int
	ScrapName() string
	ScrapState() int
	ScrapMessage() string
}

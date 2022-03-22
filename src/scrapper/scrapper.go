package scrapper

type Scrapper interface {
	Scrap()
	CheckForNew()
	Progress() int
	ScrapName() string
	ScrapState() string
	ScrapMessage() string
}

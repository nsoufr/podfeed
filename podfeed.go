package podfeed

import "encoding/xml"

func Parse(blob []byte) (pd Podcast, err error) {
	err = xml.Unmarshal(blob, &pd)
	if err != nil {
		return
	}

	return
}

type Image struct {
	Href  string `xml:"href,attr"`
	Url   string `xml:"url"`
	Title string `xml:"title"`
}

type Owner struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Category struct {
	Text string `xml:"text,attr"`
}

type Podcast struct {
	Title       string   `xml:"channel>title"`
	Subtitle    string   `xml:"channel>subtitle"`
	Description string   `xml:"channel>description"`
	Link        string   `xml:"channel>link"`
	Language    string   `xml:"channel>language"`
	Author      string   `xml:"channel>author"`
	Image       Image    `xml:"channel>image"`
	Owner       Owner    `xml:"channel>owner"`
	Category    Category `xml:"channel>category"`
}

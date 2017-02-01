/*
 Author: Nando Sousa <nandosousafr@gmail.com>
 This package allows to parse Podcast feeds.
*/

package podfeed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type Time struct {
	Value time.Time
}

const rfc2822 = "Mon, 2 Jan 2006 15:04:05 MST"

func (t *Time) UnmarshalText(data []byte) (err error) {
	str := string(data)

	tm, err := time.Parse(rfc2822, str)
	if err != nil {
			return
	}

	*t = Time{tm}
	return
}

func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.Value.String()), nil
}

func Parse(blob []byte) (pd Podcast, err error) {
	err = xml.Unmarshal(blob, &pd)
	if err != nil {
		return
	}

	return
}

func Fetch(url string) (pd Podcast, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}

	defer res.Body.Close()

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return Parse(buff)
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
	Items       []Item   `xml:"channel>item"`
}

func (p Podcast) ReleasesByWeekday() (map[string]int, error) {
	res := map[string]int{}

	for _, episode := range p.Items {
		res[episode.PubDate.Value.Weekday().String()]++
	}

	return res, nil
}

type Item struct {
	Title       string    `xml:"title"`
	PubDate     Time      `xml:"pubDate"`
	Link        string    `xml:"link"`
	Duration    string    `xml:"duration"`
	Author      string    `xml:"author"`
	Summary     string    `xml:"summary"`
	Subtitle    string    `xml:"subtitle"`
	Description string    `xml:"description"`
	Enclosure   Enclosure `xml:"enclosure"`
	Image       Image     `xml:"image"`
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

type Enclosure struct {
	Type string `xml:"type,attr"`
	Url  string `xml:"url,attr"`
}

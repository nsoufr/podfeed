package podfeed

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	blob := loadFixture()

	pod, err := Parse(blob)
	if err != nil {
		t.Fatal(err)
	}

	var asserts = []struct {
		in  string
		out string
	}{
		{pod.Title, "CapyCast"},
		{pod.Subtitle, "Perfil do time de produto da Resultados Digitais.…"},
		{pod.Description, "Perfil do time de produto da Resultados Digitais. Aqui compartilhamos sobre experiências com tecnologia construindo o RDStation."},
		{pod.Link, "http://shipit.resultadosdigitais.com.br"},
		{pod.Language, "pt"},
		{pod.Author, "Ship It"},
		{pod.Image.Url, "http://i1.sndcdn.com/avatars-000211243469-j7jvez-original.jpg"},
		{pod.Image.Href, "http://i1.sndcdn.com/avatars-000211243469-j7jvez-original.jpg"},
		{pod.Image.Title, "Ship It"},
		{pod.Owner.Name, "Ship It"},
		{pod.Owner.Email, "shipit.rd@gmail.com"},
		{pod.Category.Text, "Technology"},
		{pod.Items[0].Title, "Capycast #4 Solopreneur, Entrepreneur, Intrapreneur"},
		{pod.Items[0].PubDate.Value.String(), "2016-04-09 00:00:00 +0000 +0000"},
		{pod.Items[0].Link, "https://soundcloud.com/shipit-rd/capycast-4-solopreneur-entrepreneur-intrapreneur"},
		{pod.Items[0].Duration, "00:43:24"},
		{pod.Items[0].Author, "Ship It"},
		{pod.Items[0].Summary, "itunes:summary tag"},
		{pod.Items[0].Subtitle, "subtitle"},
		{pod.Items[0].Description, "description"},
		{pod.Items[0].Enclosure.Type, "audio/mpeg"},
		{pod.Items[0].Enclosure.Url, "http://feeds.soundcloud.com/stream/258388007-shipit-rd-capycast-4-solopreneur-entrepreneur-intrapreneur.mp3"},
		{fmt.Sprintf("%v", pod.Items[3].Enclosure.Length), "292585861"},
		{pod.Items[0].Image.Href, "http://i1.sndcdn.com/artworks-000157500330-xdws90-original.jpg"},
	}

	for _, assert := range asserts {
		if assert.in != assert.out {
			t.Errorf("got %q, want %q", assert.in, assert.out)
		}
	}
}

func TestFetch(t *testing.T) {
	blob := loadFixture()
	ctx := context.Background()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(blob)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	pod, err := Fetch(ctx, ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if pod.Title != "CapyCast" {
		t.Errorf("expected %s, got %s", pod.Title, "CapyCast")
	}
}

func TestPodcast_ReleasesByWeekday(t *testing.T) {
	podcast := loadPodcast()
	got, err := podcast.ReleasesByWeekday()
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]int{
		"Saturday": 1,
		"Monday":   3,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func loadFixture() []byte {
	blob, _ := os.ReadFile("./fixtures/podcast.rss")
	return blob
}

func loadPodcast() Podcast {
	buff := loadFixture()
	pc, _ := Parse(buff)
	return pc
}

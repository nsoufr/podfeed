package podfeed

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	blob, err := ioutil.ReadFile("./test/podcast.rss")
	if err != nil {
		t.Fatal(err)
	}

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
		{pod.Items[0].PubDate, "Sat, 09 Apr 2016 00:00:00 +0000"},
		{pod.Items[0].Link, "https://soundcloud.com/shipit-rd/capycast-4-solopreneur-entrepreneur-intrapreneur"},
		{pod.Items[0].Duration, "00:43:24"},
		{pod.Items[0].Author, "Ship It"},
		{pod.Items[0].Summary, "itunes:summary tag"},
		{pod.Items[0].Subtitle, "subtitle"},
		{pod.Items[0].Description, "description"},
		{pod.Items[0].Enclosure.Type, "audio/mpeg"},
		{pod.Items[0].Enclosure.Url, "http://feeds.soundcloud.com/stream/258388007-shipit-rd-capycast-4-solopreneur-entrepreneur-intrapreneur.mp3"},
		{pod.Items[0].Image.Href, "http://i1.sndcdn.com/artworks-000157500330-xdws90-original.jpg"},
	}

	for _, assert := range asserts {
		if assert.in != assert.out {
			t.Errorf("got %q, want %q", assert.in, assert.out)
		}
	}
}

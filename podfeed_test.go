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
	}

	for _, assert := range asserts {
		if assert.in != assert.out {
			t.Errorf("got %q, want %q", assert.in, assert.out)
		}
	}
}

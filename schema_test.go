package vast2

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVASTDefault(t *testing.T) {
	vast := VAST{}
	data, err := xml.Marshal(vast)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VAST version=""></VAST>`)
}

func TestVASTWithAttrs(t *testing.T) {
	vast := VAST{Version: "2.0"}
	data, err := xml.Marshal(vast)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VAST version="2.0"></VAST>`)
}

func TestVASTWithAds(t *testing.T) {
	vast := VAST{Version: "2.0"}

	vast.Ad = append(vast.Ad, Ad{})
	data, err := xml.Marshal(vast)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VAST version="2.0"><Ad id=""></Ad></VAST>`)

	vast.Ad = append(vast.Ad, Ad{})
	data, err = xml.Marshal(vast)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VAST version="2.0"><Ad id=""></Ad><Ad id=""></Ad></VAST>`)
}

func TestAdDefault(t *testing.T) {
	ad := Ad{}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Ad id=""></Ad>`)
}

func TestAdWithAttrs(t *testing.T) {
	ad := Ad{Id: "123"}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Ad id="123"></Ad>`)
}

func TestAdWithNestedObjects(t *testing.T) {
	ad := Ad{Id: "123", InLine: &InLine{}, Wrapper: &Wrapper{}}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)

	res := `<Ad id="123">` +
		`<InLine>` +
		`<AdTitle></AdTitle><AdSystem></AdSystem><Creatives></Creatives>` +
		`</InLine>` +
		`<Wrapper>` +
		`<VASTAdTagURI></VASTAdTagURI><AdSystem></AdSystem><Creatives></Creatives>` +
		`</Wrapper>` +
		`</Ad>`
	assert.Equal(t, string(data), res)
}

func TestInLineDefault(t *testing.T) {
	inLine := InLine{}
	data, err := xml.Marshal(inLine)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<InLine><AdTitle></AdTitle><AdSystem></AdSystem><Creatives></Creatives></InLine>`)
}

func TestInLineWithAttrs(t *testing.T) {
	inLine := InLine{
		AdTitle:     "AdName",
		Description: "Desc",
		Survey:      "http://survey.com",
		Error:       "http://err.com",
	}

	data, err := xml.Marshal(inLine)
	assert.Nil(t, err)

	res := `<InLine>` +
		`<AdTitle>AdName</AdTitle><Description>Desc</Description>` +
		`<Survey>http://survey.com</Survey><Error>http://err.com</Error>` +
		`<AdSystem></AdSystem><Creatives></Creatives>` +
		`</InLine>`
	assert.Equal(t, string(data), res)
}

func TestInLineWithNestedObjects(t *testing.T) {
	inLine := InLine{
		Impression: []Impression{Impression{}, Impression{}},
		Extensions: &Extensions{},
	}
	data, err := xml.Marshal(inLine)
	assert.Nil(t, err)

	res := `<InLine>`+
	`<AdTitle></AdTitle><AdSystem></AdSystem>`+
	`<Impression></Impression><Impression></Impression>`+
	`<Creatives></Creatives><Extensions></Extensions>`+
	`</InLine>`
	assert.Equal(t, string(data), res)
}

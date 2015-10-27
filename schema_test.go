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

	res := `<InLine>` +
		`<AdTitle></AdTitle><AdSystem></AdSystem>` +
		`<Impression></Impression><Impression></Impression>` +
		`<Creatives></Creatives><Extensions></Extensions>` +
		`</InLine>`
	assert.Equal(t, string(data), res)
}

func TestAdSystemDefault(t *testing.T) {
	adSystem := AdSystem{}
	data, err := xml.Marshal(adSystem)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<AdSystem></AdSystem>`)
}

func TestAdSystemWithAttrs(t *testing.T) {
	adSystem := AdSystem{Version: "2.1", Data: "ad_server"}
	data, err := xml.Marshal(adSystem)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<AdSystem version="2.1">ad_server</AdSystem>`)
}

func TestImpressionDefault(t *testing.T) {
	imp := Impression{}
	data, err := xml.Marshal(imp)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Impression></Impression>`)
}

func TestImpressionWithAttrs(t *testing.T) {
	imp := Impression{Id: "1", Data: "http://imp.com"}
	data, err := xml.Marshal(imp)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Impression id="1">http://imp.com</Impression>`)
}

func TestCreativesDefault(t *testing.T) {
	creatives := Creatives{}
	data, err := xml.Marshal(creatives)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Creatives></Creatives>`)
}

func TestCreativesWithNestedObjects(t *testing.T) {
	creatives := Creatives{Creative: []Creative{Creative{}, Creative{}}}
	data, err := xml.Marshal(creatives)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Creatives><Creative></Creative><Creative></Creative></Creatives>`)
}

func TestCreativeDefault(t *testing.T) {
	creative := Creative{}
	data, err := xml.Marshal(creative)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Creative></Creative>`)
}

func TestCreativeWithAttrs(t *testing.T) {
	creative := Creative{Id: "a1", Sequence: 2, AdID: "ID3"}
	data, err := xml.Marshal(creative)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Creative id="a1" sequence="2" AdID="ID3"></Creative>`)
}

func TestCreativeWithNestedObjects(t *testing.T) {
	creative := Creative{Linear: &Linear{}, CompanionAds: &CompanionAds{}, NonLinearAds: &NonLinearAds{}}
	data, err := xml.Marshal(creative)
	assert.Nil(t, err)

	res := `<Creative>` +
		`<Linear>` +
		`<Duration></Duration><MediaFiles></MediaFiles>` +
		`</Linear>` +
		`<CompanionAds></CompanionAds>` +
		`<NonLinearAds></NonLinearAds>` +
		`</Creative>`
	assert.Equal(t, string(data), res)
}

func TestLinearDefault(t *testing.T) {
	linear := Linear{}
	data, err := xml.Marshal(linear)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Linear><Duration></Duration><MediaFiles></MediaFiles></Linear>`)
}

func TestLinearWithAttrs(t *testing.T) {
	linear := Linear{Duration: "2:00", AdParameters: "a=1"}
	data, err := xml.Marshal(linear)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Linear><Duration>2:00</Duration><AdParameters>a=1</AdParameters><MediaFiles></MediaFiles></Linear>`)
}

func TestLinearWithNestedObjects(t *testing.T) {
	linear := Linear{
		TrackingEvents: &TrackingEvents{},
		VideoClicks:    &VideoClicks{},
		MediaFiles:     MediaFiles{},
	}

	data, err := xml.Marshal(linear)
	assert.Nil(t, err)

	res := `<Linear>` +
		`<Duration></Duration>` +
		`<TrackingEvents></TrackingEvents>` +
		`<VideoClicks></VideoClicks>` +
		`<MediaFiles></MediaFiles>` +
		`</Linear>`
	assert.Equal(t, string(data), res)
}

func TestTrackingEventsDefault(t *testing.T) {
	events := TrackingEvents{}
	data, err := xml.Marshal(events)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<TrackingEvents></TrackingEvents>`)
}

func TestTrackingEventsWithNestedObjects(t *testing.T) {
	events := TrackingEvents{Tracking: []Tracking{Tracking{}, Tracking{}}}
	data, err := xml.Marshal(events)
	assert.Nil(t, err)

	res := `<TrackingEvents>` +
		`<Tracking event=""></Tracking><Tracking event=""></Tracking>` +
		`</TrackingEvents>`
	assert.Equal(t, string(data), res)
}

func TestTrackingDefault(t *testing.T) {
	tracking := Tracking{}
	data, err := xml.Marshal(tracking)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Tracking event=""></Tracking>`)
}

func TestTrackingWithAttrs(t *testing.T) {
	tracking := Tracking{Event: "close", Data: "http://url.com"}
	data, err := xml.Marshal(tracking)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Tracking event="close">http://url.com</Tracking>`)
}

func TestVideoClicksDefault(t *testing.T) {
	clicks := VideoClicks{}
	data, err := xml.Marshal(clicks)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VideoClicks></VideoClicks>`)
}

func TestVideoClicksWithAttrs(t *testing.T) {
	clicks := VideoClicks{
		ClickThrough:  "http://clk.com",
		ClickTracking: "http://tr.com",
	}
	data, err := xml.Marshal(clicks)
	assert.Nil(t, err)

	res := `<VideoClicks>` +
		`<ClickThrough>http://clk.com</ClickThrough>` +
		`<ClickTracking>http://tr.com</ClickTracking>` +
		`</VideoClicks>`
	assert.Equal(t, string(data), res)
}

func TestVideoClicksWithNestedObjects(t *testing.T) {
	clicks := VideoClicks{CustomClick: &CustomClick{}}
	data, err := xml.Marshal(clicks)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<VideoClicks><CustomClick></CustomClick></VideoClicks>`)
}

func TestCustomClickDefault(t *testing.T) {
	click := CustomClick{}
	data, err := xml.Marshal(click)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<CustomClick></CustomClick>`)
}

func TestCustomClickWithAttrs(t *testing.T) {
	click := CustomClick{Id: "1", Data: "http://clk.com"}
	data, err := xml.Marshal(click)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<CustomClick id="1">http://clk.com</CustomClick>`)
}

func TestMediaFilesDefault(t *testing.T) {
	files := MediaFiles{}
	data, err := xml.Marshal(files)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<MediaFiles></MediaFiles>`)
}

func TestMediaFilesWithNestedObjects(t *testing.T) {
	files := MediaFiles{MediaFile: []MediaFile{MediaFile{}, MediaFile{}}}
	data, err := xml.Marshal(files)
	assert.Nil(t, err)

	res := `<MediaFiles>` +
		`<MediaFile delivery="" type="" width="0" height="0"></MediaFile>` +
		`<MediaFile delivery="" type="" width="0" height="0"></MediaFile>` +
		`</MediaFiles>`
	assert.Equal(t, string(data), res)
}

func TestMediaFileDefault(t *testing.T) {
	file := MediaFile{}
	data, err := xml.Marshal(file)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<MediaFile delivery="" type="" width="0" height="0"></MediaFile>`)
}

func TestMediaFileWithAttrs(t *testing.T) {
	file := MediaFile{
		Id:                  "123",
		Delivery:            "progressive",
		Type:                "video/mp4",
		Bitrate:             1,
		Width:               300,
		Height:              250,
		Scalable:            true,
		MaintainAspectRatio: true,
		ApiFramework:        "js",
		Data:                "http://site.com/video.mp4",
	}
	data, err := xml.Marshal(file)
	assert.Nil(t, err)

	res := `<MediaFile id="123" delivery="progressive" ` +
		`type="video/mp4" bitrate="1" width="300" height="250" scalable="true" ` +
		`maintainAspectRatio="true" apiFramework="js">http://site.com/video.mp4</MediaFile>`
	assert.Equal(t, string(data), res)
}

func TestCompanionAdsDefault(t *testing.T) {
	ads := CompanionAds{}
	data, err := xml.Marshal(ads)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<CompanionAds></CompanionAds>`)
}

func TestCompanionAdsWithNestedObjects(t *testing.T) {
	ads := CompanionAds{Companion: []Companion{Companion{}, Companion{}}}
	data, err := xml.Marshal(ads)
	assert.Nil(t, err)

	res := `<CompanionAds>` +
		`<Companion width="0" height="0"></Companion>` +
		`<Companion width="0" height="0"></Companion>` +
		`</CompanionAds>`

	assert.Equal(t, string(data), res)
}

func TestCompanionDefault(t *testing.T) {
	companion := Companion{}
	data, err := xml.Marshal(companion)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<Companion width="0" height="0"></Companion>`)
}

func TestCompanionWithAttrs(t *testing.T) {
	companion := Companion{
		Id:                    "ab1",
		Width:                 300,
		Height:                250,
		ExpandedWidth:         350,
		ExpandedHeight:        300,
		ApiFramework:          "js",
		IFrameResource:        "http://iframe.com",
		HTMLResource:          "<body></body>",
		CompanionClickThrough: "http://clk.com",
		AltText:               "ad text",
		AdParameters:          "a=1",
	}
	data, err := xml.Marshal(companion)
	assert.Nil(t, err)

	res := `<Companion id="ab1" width="300" height="250" ` +
		`expandedWidth="350" expandedHeight="300" apiFramework="js">` +
		`<IFrameResource>http://iframe.com</IFrameResource>` +
		`<HTMLResource>&lt;body&gt;&lt;/body&gt;</HTMLResource>` +
		`<CompanionClickThrough>http://clk.com</CompanionClickThrough>` +
		`<AltText>ad text</AltText><AdParameters>a=1</AdParameters>` +
		`</Companion>`
	assert.Equal(t, string(data), res)
}

func TestCompanionWithNestedObjects(t *testing.T) {
	companion := Companion{
		StaticResource: &StaticResource{},
		TrackingEvents: &TrackingEvents{},
	}
	data, err := xml.Marshal(companion)
	assert.Nil(t, err)

	res := `<Companion width="0" height="0">` +
		`<StaticResource creativeType=""></StaticResource>` +
		`<TrackingEvents></TrackingEvents>` +
		`</Companion>`

	assert.Equal(t, string(data), res)
}

func TestStaticResourceDefault(t *testing.T) {
	resource := StaticResource{}
	data, err := xml.Marshal(resource)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<StaticResource creativeType=""></StaticResource>`)
}

func TestStaticResourceWithAttrs(t *testing.T) {
	resource := StaticResource{CreativeType: "image/png", Data: "http://img.png"}
	data, err := xml.Marshal(resource)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<StaticResource creativeType="image/png">http://img.png</StaticResource>`)
}

func TestNonLinearAdsDefault(t *testing.T) {
	ads := NonLinearAds{}
	data, err := xml.Marshal(ads)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<NonLinearAds></NonLinearAds>`)
}

func TestNonLinearAdsWithNestedObjects(t *testing.T) {
	ads := NonLinearAds{NonLinear: []NonLinear{NonLinear{}, NonLinear{}}}
	data, err := xml.Marshal(ads)
	assert.Nil(t, err)

	res := `<NonLinearAds>` +
		`<NonLinear width="0" height="0"></NonLinear>` +
		`<NonLinear width="0" height="0"></NonLinear>` +
		`</NonLinearAds>`

	assert.Equal(t, string(data), res)
}

func TestNonLinearDefault(t *testing.T) {
	ad := NonLinear{}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)
	assert.Equal(t, string(data), `<NonLinear width="0" height="0"></NonLinear>`)
}

func TestNonLinearWithAttrs(t *testing.T) {
	ad := NonLinear{
		Id:                    "a1",
		Width:                 300,
		Height:                250,
		ExpandedWidth:         350,
		ExpandedHeight:        300,
		Scalable:              true,
		MaintainAspectRatio:   true,
		ApiFramework:          "js",
		IFrameResource:        "http://iframe.com",
		HTMLResource:          "<body></body>",
		AdParameters:          "a=1",
		NonLinearClickThrough: "http://clk.com",
	}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)
	res := `<NonLinear id="a1" width="300" height="250" ` +
		`expandedWidth="350" expandedHeight="300" scalable="true" ` +
		`maintainAspectRatio="true" apiFramework="js">` +
		`<IFrameResource>http://iframe.com</IFrameResource>` +
		`<HTMLResource>&lt;body&gt;&lt;/body&gt;</HTMLResource>` +
		`<AdParameters>a=1</AdParameters>` +
		`<NonLinearClickThrough>http://clk.com</NonLinearClickThrough>` +
		`</NonLinear>`

	assert.Equal(t, string(data), res)
}

func TestNonLinearWithNestedObjects(t *testing.T) {
	ad := NonLinear{
		StaticResource: &StaticResource{},
		TrackingEvents: &TrackingEvents{},
	}
	data, err := xml.Marshal(ad)
	assert.Nil(t, err)

	res := `<NonLinear width="0" height="0">` +
		`<StaticResource creativeType=""></StaticResource>` +
		`<TrackingEvents></TrackingEvents>` +
		`</NonLinear>`
	assert.Equal(t, string(data), res)
}

func TestExtensions(t *testing.T) {
	ext := Extensions{
		Extension: []Extension{
			Extension{
				Type: "price",
				Data: []byte("<Price>2.5</Price>"),
			},
			Extension{
				Type: "model",
				Data: []byte("<Model>cpm</Model>"),
			},
		},
	}
	data, err := xml.Marshal(ext)
	assert.Nil(t, err)

	res := `<Extensions>` +
		`<Extension type="price"><Price>2.5</Price></Extension>` +
		`<Extension type="model"><Model>cpm</Model></Extension>` +
		`</Extensions>`
	assert.Equal(t, string(data), res)
}

func TestWrapperDefault(t *testing.T) {
	wrapper := Wrapper{}
	data, err := xml.Marshal(wrapper)
	assert.Nil(t, err)

	res := `<Wrapper>` +
		`<VASTAdTagURI></VASTAdTagURI>` +
		`<AdSystem></AdSystem><Creatives></Creatives>` +
		`</Wrapper>`
	assert.Equal(t, string(data), res)
}

func TestWrapperWithAttrs(t *testing.T) {
	wrapper := Wrapper{
		VASTAdTagURI: "http://tag.com",
		Error:        "http://error.com",
	}
	data, err := xml.Marshal(wrapper)
	assert.Nil(t, err)

	res := `<Wrapper>` +
		`<VASTAdTagURI>http://tag.com</VASTAdTagURI>` +
		`<Error>http://error.com</Error>` +
		`<AdSystem></AdSystem><Creatives></Creatives>` +
		`</Wrapper>`
	assert.Equal(t, string(data), res)
}

func TestWrapperWithNestedObjects(t *testing.T) {
	wrapper := Wrapper{
		AdSystem:   AdSystem{},
		Creatives:  Creatives{},
		Extensions: &Extensions{},
	}
	data, err := xml.Marshal(wrapper)
	assert.Nil(t, err)

	res := `<Wrapper>` +
		`<VASTAdTagURI></VASTAdTagURI>` +
		`<AdSystem></AdSystem><Creatives></Creatives>` +
		`<Extensions></Extensions>` +
		`</Wrapper>`
	assert.Equal(t, string(data), res)
}

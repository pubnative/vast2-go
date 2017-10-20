package vast2

type VAST struct {
	Version string `xml:"version,attr"`
	Ad      []Ad   `xml:"Ad"`
}

type Ad struct {
	ID      string   `xml:"id,attr"`
	InLine  *InLine  `xml:"InLine,omitempty"`
	Wrapper *Wrapper `xml:"Wrapper,omitempty"`
}

type InLine struct {
	AdTitle     string       `xml:"AdTitle"`
	Description string       `xml:"Description,omitempty"`
	Survey      string       `xml:"Survey,omitempty"`
	Error       string       `xml:"Error,omitempty"`
	AdSystem    AdSystem     `xml:"AdSystem"`
	Impression  []Impression `xml:"Impression,omitempty"`
	Creatives   Creatives    `xml:"Creatives"`
	Extensions  *Extensions  `xml:"Extensions,omitempty"`
}

type AdSystem struct {
	Version string `xml:"version,attr,omitempty"`
	Data    string `xml:",chardata"`
}

type Impression struct {
	ID   string `xml:"id,attr,omitempty"`
	Data string `xml:",chardata"`
}

type Creatives struct {
	Creative []Creative `xml:"Creative"`
}

type Creative struct {
	ID           string        `xml:"id,attr,omitempty"`
	Sequence     int           `xml:"sequence,attr,omitempty"`
	AdID         string        `xml:"AdID,attr,omitempty"`
	Linear       *Linear       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAds `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty"`
}

type Linear struct {
	Duration       string          `xml:"Duration"`
	AdParameters   string          `xml:"AdParameters,omitempty"`
	TrackingEvents *TrackingEvents `xml:"TrackingEvents,omitempty"`
	VideoClicks    *VideoClicks    `xml:"VideoClicks,omitempty"`
	MediaFiles     MediaFiles      `xml:"MediaFiles"`
}

type TrackingEvents struct {
	Tracking []Tracking `xml:"Tracking,omitempty"`
}

type Tracking struct {
	Event string `xml:"event,attr"`
	Data  string `xml:",chardata"`
}

type VideoClicks struct {
	ClickThrough  string       `xml:"ClickThrough,omitempty"`
	ClickTracking string       `xml:"ClickTracking,omitempty"`
	CustomClick   *CustomClick `xml:"CustomClick,omitempty"`
}

type CustomClick struct {
	ID   string `xml:"id,attr,omitempty"`
	Data string `xml:",chardata"`
}

type MediaFiles struct {
	MediaFile []MediaFile `xml:"MediaFile"`
}

type MediaFile struct {
	ID                  string `xml:"id,attr,omitempty"`
	Delivery            string `xml:"delivery,attr"`
	Type                string `xml:"type,attr"`
	Bitrate             int    `xml:"bitrate,attr,omitempty"`
	Width               int    `xml:"width,attr"`
	Height              int    `xml:"height,attr"`
	Scalable            bool   `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio bool   `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework        string `xml:"apiFramework,attr,omitempty"`
	Data                string `xml:",chardata"`
}

type CompanionAds struct {
	Companion []Companion `xml:"Companion,omitempty"`
}

type Companion struct {
	ID                    string          `xml:"id,attr,omitempty"`
	Width                 int             `xml:"width,attr"`
	Height                int             `xml:"height,attr"`
	ExpandedWidth         int             `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        int             `xml:"expandedHeight,attr,omitempty"`
	ApiFramework          string          `xml:"apiFramework,attr,omitempty"`
	IFrameResource        string          `xml:"IFrameResource,omitempty"`
	HTMLResource          string          `xml:"HTMLResource,omitempty"`
	CompanionClickThrough string          `xml:"CompanionClickThrough,omitempty"`
	AltText               string          `xml:"AltText,omitempty"`
	AdParameters          string          `xml:"AdParameters,omitempty"`
	StaticResource        *StaticResource `xml:"StaticResource,omitempty"`
	TrackingEvents        *TrackingEvents `xml:"TrackingEvents,omitempty"`
}

type StaticResource struct {
	CreativeType string `xml:"creativeType,attr"`
	Data         string `xml:",chardata"`
}

type NonLinearAds struct {
	NonLinear      []NonLinear     `xml:"NonLinear,omitempty"`
	TrackingEvents *TrackingEvents `xml:"TrackingEvents,omitempty"`
}

type NonLinear struct {
	ID                    string          `xml:"id,attr,omitempty"`
	Width                 int             `xml:"width,attr"`
	Height                int             `xml:"height,attr"`
	ExpandedWidth         int             `xml:"expandedWidth,attr,omitempty"`
	ExpandedHeight        int             `xml:"expandedHeight,attr,omitempty"`
	Scalable              bool            `xml:"scalable,attr,omitempty"`
	MaintainAspectRatio   bool            `xml:"maintainAspectRatio,attr,omitempty"`
	ApiFramework          string          `xml:"apiFramework,attr,omitempty"`
	IFrameResource        string          `xml:"IFrameResource,omitempty"`
	HTMLResource          string          `xml:"HTMLResource,omitempty"`
	AdParameters          string          `xml:"AdParameters,omitempty"`
	NonLinearClickThrough string          `xml:"NonLinearClickThrough,omitempty"`
	StaticResource        *StaticResource `xml:"StaticResource,omitempty"`
}

type Extensions struct {
	Extension []Extension `xml:"Extension"`
}

type Extension struct {
	Type string `xml:"type,attr,omitempty"`
	Data []byte `xml:",innerxml"`
}

type Wrapper struct {
	VASTAdTagURI string       `xml:"VASTAdTagURI"`
	Error        string       `xml:"Error,omitempty"`
	AdSystem     AdSystem     `xml:"AdSystem"`
	Impression   []Impression `xml:"Impression,omitempty"`
	Creatives    Creatives    `xml:"Creatives"`
	Extensions   *Extensions  `xml:"Extensions,omitempty"`
}

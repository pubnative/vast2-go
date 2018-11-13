package vast2

// https://www.iab.com/wp-content/uploads/2015/11/VAST-2_0-FINAL.pdf

// Could be present in CompanionAds, Linear, or NonLinearAds
const (
	TrackingEventCreativeView     = "creativeView"
)

// Tracking Events
// Present in Linear, and NonLinearAds
// VAST Version 2 Final Page 13, 14.
const (
	TrackingEventStart            = "start"
	TrackingEventMindPoint        = "midpoint"
	TrackingEventFirstQuartile    = "firstQuartile"
	TrackingEventThirdQuartile    = "thirdQuartile"
	TrackingEventComplete         = "complete"
	TrackingEventMute             = "mute"
	TrackingEventUnmute           = "unmute"
	TrackingEventPause            = "pause"
	TrackingEventRewing           = "rewind"
	TrackingEventResume           = "resume"
	TrackingEventFullscreen       = "fullscreen"
	TrackingEventExpand           = "expand"
	TrackingEventCollapse         = "collapse"
	TrackingEventAcceptInvitation = "acceptInvitation"
	TrackingEventClose            = "close"
)


// MediaFilesDelivery
// Present Only in Linear vasts
// VAST Version 2 Final Page 11.
const (
	LinearMediaFileDeliveryStreaming = "streaming"
	LinearMediaFileDeliveryProgressive = "progressive"
)

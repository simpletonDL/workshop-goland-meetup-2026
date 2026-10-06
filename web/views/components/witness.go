package components

import (
	"strconv"
	"strings"
	"workshop/internal/domain/sighting"
)

var quotes = []string{
	"He went that way.",
	"Asked me for the Wi-Fi password. I'm a buffalo.",
	"Nice lanyard. Smelled of coffee.",
	"He tried to explain goroutines to me. I left.",
	"Never seen him. Moo.",
	"He kept asking where 'the venue' was.",
	"Let him ride on my back for a bit. Four stars.",
	"A lion asked me if he was tasty. I said no comment.",
	"He said he was 'mostly protein shakes'.",
	"He waved. I didn't wave back. I have hooves.",
	"Looked lost. Also looked like he'd skipped lunch.",
	"He offered me a sticker. I ate it.",
}

// quote gives a witness its statement, the same one on every render.
func quote(n int) string {
	return quotes[n%len(quotes)]
}

// firstPhoto returns the first still image of a sighting, if it has one.
func firstPhoto(media []sighting.Media) (sighting.Media, bool) {
	for _, m := range media {
		if m.Type == "StillImage" && m.URL != "" {
			return m, true
		}
	}
	return sighting.Media{}, false
}

// thumb asks iNaturalist and Flickr for a small copy of a photo rather than the original.
func thumb(url string) string {
	switch {
	case strings.HasPrefix(url, "https://inaturalist-open-data.s3.amazonaws.com/"):
		return strings.Replace(url, "/original.", "/small.", 1)
	case strings.HasPrefix(url, "https://live.staticflickr.com/") && strings.HasSuffix(url, "_b.jpg"):
		return strings.TrimSuffix(url, "_b.jpg") + "_n.jpg"
	}
	return url
}

// witnessAlt describes a witness's photo, crediting whoever took it.
func witnessAlt(n int, m sighting.Media) string {
	alt := "Buffalo witness #" + strconv.Itoa(n)
	if m.Creator != "" {
		alt += ", photographed by " + m.Creator
	}
	return alt
}

package pages

import (
	"strconv"
	"workshop/internal/domain/sighting"
	"workshop/web/views/components"
)

// chapters is Ainsley's trip stop by stop, each with how many buffalo saw him there.
func chapters(sightings []sighting.Sighting) []components.Chapter {
	seen := sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return string(s.Location.CountryCode)
	})

	return []components.Chapter{
		{
			ID:      "start",
			Short:   "Start",
			Day:     "Day 0",
			Place:   "Somewhere in Africa",
			Title:   "Where's Ainsley?",
			Story:   "He popped out for a flat white. Since then he's been spotted by " + buffalo(len(sightings)) + " and chased by most of the food chain.",
			Ainsley: "Has anyone seen a coffee shop?",
			Pose:    "stand",
			Sky:     "mustard",
			Ground:  "orange",
			Link:    &components.ChapterLink{Href: "#kenya", Label: "Follow the trail ↓"},
			Cast: []components.Animal{
				{Emoji: "🦅", Spot: "sky", Move: "circle", Say: "I'll wait."},
				{Emoji: "🐃", Spot: "right", Move: "bob", Say: "Moo?"},
				{Emoji: "🌴", Spot: "back-left", Move: "sway", Size: "l"},
			},
		},
		{
			ID:        "kenya",
			Short:     "Kenya",
			Day:       "Day 1",
			Place:     "Maasai Mara, Kenya",
			Title:     "He asked the lions for directions",
			Story:     "They were very helpful. They asked about his dietary requirements, whether he was free-range, and how he felt about being shared.",
			Witnesses: seenHere(seen["KE"]),
			Ainsley:   "I'm mostly protein shakes, honestly.",
			Pose:      "stand",
			Sky:       "coral",
			Ground:    "mustard",
			Cast: []components.Animal{
				{Emoji: "🦁", Spot: "left", Move: "chomp", Size: "l", Say: "Is he gluten-free?"},
				{Emoji: "🦁", Spot: "right", Move: "sneak", Say: "Dibs on the lanyard."},
				{Emoji: "🦓", Spot: "high-right", Move: "bob", Size: "s", Say: "Run, mate."},
				{Emoji: "🌳", Spot: "back-left", Move: "sway", Size: "l"},
			},
		},
		{
			ID:        "tanzania",
			Short:     "Tanzania",
			Day:       "Day 4",
			Place:     "Serengeti, Tanzania",
			Title:     "He climbed Kilimanjaro. It was a giraffe.",
			Story:     "The view was lovely. The giraffe has since filed a formal complaint, and the vultures have started a waiting list.",
			Witnesses: seenHere(seen["TZ"]),
			Ainsley:   "In my defence, it was very tall.",
			Pose:      "stand",
			Sky:       "lilac",
			Ground:    "apricot",
			Cast: []components.Animal{
				{Emoji: "🦒", Spot: "left", Move: "bob", Size: "l", Flip: true, Say: "Get. Off. My. Neck."},
				{Emoji: "🦅", Spot: "sky", Move: "circle", Say: "Table for one?"},
				{Emoji: "🐘", Spot: "right", Move: "sneak", Size: "l", Say: "I'll remember this."},
				{Emoji: "🏔️", Spot: "back-right", Size: "l"},
			},
		},
		{
			ID:        "botswana",
			Short:     "Botswana",
			Day:       "Day 9",
			Place:     "Okavango Delta, Botswana",
			Title:     "He crossed the river on some logs",
			Story:     "The logs had teeth. A hippo watched the whole thing and did nothing, which is the most dangerous thing a hippo can do.",
			Witnesses: seenHere(seen["BW"]),
			Ainsley:   "Why is this log smiling at me?",
			Pose:      "stand",
			Sky:       "azure",
			Ground:    "teal",
			Cast: []components.Animal{
				{Emoji: "🐊", Spot: "front-left", Move: "peek", Flip: true, Say: "Lovely day for a swim!"},
				{Emoji: "🐊", Spot: "front-right", Move: "peek", Say: "Water's lovely. Come in."},
				{Emoji: "🦛", Spot: "right", Move: "bob", Size: "l", Say: "This is MY river."},
				{Emoji: "🦩", Spot: "high-left", Move: "bob", Size: "s", Say: "Not getting involved."},
			},
		},
		{
			ID:        "uganda",
			Short:     "Uganda",
			Day:       "Day 16",
			Place:     "Bwindi Forest, Uganda",
			Title:     "He gave a talk on Go generics to some gorillas",
			Story:     "Attendance: 100%. Questions: none. Bananas thrown: 37. The feedback forms were mostly bananas too.",
			Witnesses: seenHere(seen["UG"]),
			Ainsley:   "So, type parameters. Any questions? No? Great.",
			Pose:      "mic",
			Sky:       "grass",
			Ground:    "jungle",
			Cast: []components.Animal{
				{Emoji: "🦍", Spot: "left", Move: "shake", Size: "l", Flip: true, Say: "Ooh ooh. (Boring.)"},
				{Emoji: "🐒", Spot: "right", Move: "bob", Say: "Is there pizza after?"},
				{Emoji: "🐍", Spot: "high-left", Move: "sneak", Size: "s", Say: "Ssso, generics?"},
				{Emoji: "🍌", Spot: "high-right", Move: "throw", Size: "s"},
				{Emoji: "🍌", Spot: "high-right", Move: "lob", Size: "s"},
				{Emoji: "🌴", Spot: "back-right", Move: "sway", Size: "l"},
			},
		},
		{
			ID:        "namibia",
			Short:     "Namibia",
			Day:       "Day 23",
			Place:     "Namib Desert, Namibia",
			Title:     "He found an oasis. It was a mirage.",
			Story:     "The snake, however, was very real and very keen on hugs. Ainsley has left the desert a strongly worded review.",
			Witnesses: seenHere(seen["NA"]),
			Ainsley:   "One star. No shade. No Wi-Fi.",
			Pose:      "stand",
			Sky:       "orange",
			Ground:    "sand",
			Cast: []components.Animal{
				{Emoji: "🐍", Spot: "front-left", Move: "sneak", Flip: true, Say: "Hugsss?"},
				{Emoji: "🦂", Spot: "front-right", Move: "chomp", Size: "s", Say: "Pinch pinch."},
				{Emoji: "🐪", Spot: "right", Move: "sneak", Size: "l", Say: "Not my desert, mate."},
				{Emoji: "🌵", Spot: "back-left", Move: "sway", Size: "l"},
			},
		},
		{
			ID:        "south-africa",
			Short:     "South Africa",
			Day:       "Day 38",
			Place:     "Boulders Beach, South Africa",
			Title:     "He reached the end of Africa",
			Story:     "The penguins offered him a job. The great white offered him a hug. He's still thinking about both offers.",
			Witnesses: seenHere(seen["ZA"]),
			Ainsley:   "Is this a hiring event?",
			Pose:      "stand",
			Sky:       "aqua",
			Ground:    "ocean",
			Cast: []components.Animal{
				{Emoji: "🐧", Spot: "left", Move: "jump", Flip: true, Say: "Ainsley! Ainsley!"},
				{Emoji: "🐧", Spot: "front-left", Move: "jump", Size: "s", Flip: true, Say: "He's one of us now."},
				{Emoji: "🦈", Spot: "front-right", Move: "swim", Say: "Hug?"},
				{Emoji: "🦭", Spot: "right", Move: "bob", Say: "Got any fish?"},
			},
		},
		{
			ID:      "now",
			Short:   "Now",
			Day:     "Day " + strconv.Itoa(len(sightings)),
			Place:   "Everywhere, apparently",
			Title:   "The buffalo have him now",
			Story:   buffalo(len(sightings)) + " in " + countries(len(seen)) + " swear they've seen him. We've pinned every statement on a map. Help us find him.",
			Ainsley: "Honestly? The buffalo are lovely.",
			Pose:    "stand",
			Sky:     "mustard",
			Ground:  "coral",
			Link:    &components.ChapterLink{Href: "#search", Label: "Open Search HQ →"},
			Cast: []components.Animal{
				{Emoji: "🐃", Spot: "left", Move: "bob", Size: "l", Flip: true, Say: "He's ours now."},
				{Emoji: "🐃", Spot: "right", Move: "bob", Size: "l", Say: "Moo."},
				{Emoji: "🐃", Spot: "front-left", Move: "sneak", Flip: true},
				{Emoji: "🐃", Spot: "front-right", Move: "sneak", Say: "Group photo!"},
			},
		},
	}
}

// buffalo counts buffalo, which are buffalo whether there's one or many.
func buffalo(n int) string {
	if n == 0 {
		return "no buffalo at all"
	}
	return strconv.Itoa(n) + " buffalo"
}

// countries counts countries.
func countries(n int) string {
	if n == 1 {
		return "1 country"
	}
	return strconv.Itoa(n) + " countries"
}

// seenHere says how many buffalo reported seeing Ainsley in one country.
func seenHere(n int) string {
	switch n {
	case 0:
		return "Not one buffalo saw him here. Suspicious."
	case 1:
		return "1 buffalo saw him here. It isn't talking."
	}
	return strconv.Itoa(n) + " buffalo saw him here. None of them helped."
}

// statements introduces the witness list.
func statements(n int) string {
	switch n {
	case 0:
		return "Nobody has come forward."
	case 1:
		return "One buffalo has come forward. Its statement is below."
	}
	return strconv.Itoa(n) + " buffalo have come forward. Their statements are below, in the order they were taken. Click a pin on the map to jump to one."
}

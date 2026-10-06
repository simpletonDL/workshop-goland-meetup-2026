package components

type (
	// Chapter is one stop on Ainsley's trip: a caption on the page and the
	// scene the stage behind it switches to.
	Chapter struct {
		ID        string
		Short     string
		Day       string
		Place     string
		Title     string
		Story     string
		Witnesses string
		Ainsley   string
		Pose      string
		Sky       string
		Ground    string
		Cast      []Animal
		Link      *ChapterLink
	}
	// ChapterLink sends the reader on from a chapter's caption.
	ChapterLink struct {
		Href  string
		Label string
	}
	// Animal is one emoji in a scene. Spot, Move and Size name the
	// .spot-*, .move-* and .size-* classes in trip.css.
	Animal struct {
		Emoji string
		Spot  string
		Move  string
		Size  string
		Say   string
		Flip  bool
	}
)

// animalClass joins the classes that place and animate an animal.
func animalClass(a Animal) []string {
	classes := []string{"animal", "spot-" + a.Spot}
	if a.Move != "" {
		classes = append(classes, "move-"+a.Move)
	}
	if a.Size != "" {
		classes = append(classes, "size-"+a.Size)
	}
	if a.Flip {
		classes = append(classes, "is-flipped")
	}
	return classes
}

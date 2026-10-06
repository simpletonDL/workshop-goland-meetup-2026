# Where's Ainsley? Design

Ainsley is lost in Africa. The site is a scroll-driven trip: he stays in the middle of the
screen while the locations change around him, and the animals in each one try to eat him or
otherwise make his day worse. The sightings (all African buffalo) are the witnesses.

**Style: cartoon brutalism.** Flat, bright colours, thick ink outlines, hard offset shadows,
everything looking like a sticker. Modern, but silly. Copy is in English, deadpan and short.

## Palette

Defined on `:root` in `assets/css/app.css`. Text is always `--ink`, which passes contrast on
every colour below. The trip's chapters name these tokens for their sky and ground.

| Token | Hex | Used for |
|-------|-----|----------|
| `--ink` | `#121212` | Text, every outline, every shadow |
| `--paper` | `#FFF6E0` | Page background, witness cards |
| `--white` | `#FFFFFF` | Captions, bubbles, nav pills |
| `--mustard` | `#FFDB58` | Header, buttons, intro sky, eyebrow text on ink |
| `--coral` | `#FF6B6B` | Markers, Kenya sky, accent stat |
| `--teal` | `#4ECDC4` | Search HQ background, river |
| `--orange` | `#FF9F43` | Intro ground, desert sky |
| `--apricot` | `#FFB86B` | Serengeti ground |
| `--lilac` | `#B8A1FF` | Serengeti sky |
| `--azure` | `#7FC8F8` | Okavango sky |
| `--grass` / `--jungle` | `#7BD389` / `#3BB273` | Forest sky / ground |
| `--sand` | `#FFE1A8` | Desert ground, photo placeholder |
| `--aqua` / `--ocean` | `#A0E7E5` / `#4D96FF` | Cape sky / sea |

## Type

Rubik from Google Fonts, 400/500/700/900, as `--font`. Headings are 900 with tight leading
(1.05, the hero 0.95); labels and buttons 700; body 400 at 16 to 18px.

## Shape and spacing

- Outline: `--border` (3px solid ink) on cards, pills, the map, photos. Small things use 2.5px.
- Shadow: hard, never blurred: `--shadow-sm` 3px, `--shadow` 5px, `--shadow-lg` 8px.
- Radius: `--radius` 14px for cards, `--radius-sm` 8px for photos, `--pill` for buttons and tags.
- Spacing: `--space-1` to `--space-8` (4, 8, 12, 16, 24, 32, 48, 72px); `--gutter` for page edges.
- Buttons lift on hover (translate up-left, shadow grows) and press flat on click.
- Stat stickers sit slightly rotated and straighten on hover.

## Components

- **Header**: mustard, sticky, ink bottom border. Buffalo-head logo (`assets/images/logo.svg`,
  also the favicon) and the wordmark, which hides under 640px. Nav: white pills.
- **Trip** (`components/trip.templ`, `assets/css/trip.css`, `assets/js/trip.js`): a sticky stage
  under the header with a dotted halftone sky, a sun with a spinning dashed ring, a hill and a
  ground drawn as outlined SVG waves. Ainsley is a cut-out photo with a white sticker edge and a
  hard shadow; he bobs nervously, sweats, and can be poked (it's a button). Chapters scroll over
  the stage as white caption cards (left column on desktop, top on mobile). The chapter in the
  middle of the screen sets the sky, the ground, Ainsley's pose (`stand` or `mic`) and which cast
  of animals is shown. The trail of pills at the top is the progress indicator.
- **Animals**: native emoji with a hard drop shadow. Placed by spot (`left`, `right`,
  `high-left`, `high-right`, `sky`, `front-left`, `front-right`, `back-left`, `back-right`),
  animated by move (`chomp`, `sneak`, `bob`, `circle`, `sway`, `throw`, `lob`, `jump`, `shake`,
  `peek`, `swim`), sized `s` or `l`. Speech bubbles are white with an ink outline and pop in after
  the animal. A new chapter is one entry in `chapters()` in `pages/home.go`.
- **Search HQ**: teal section, four stat stickers, then the Leaflet map framed like a card,
  coral cluster markers with ink outlines. Ainsley's face sits on a random buffalo and hops to
  another when clicked.
- **Witness statements**: a scrollable white case file holding every sighting as a card: a
  square photo (88px, `object-fit: cover`), "Witness #n", a made-up quote, then when, where and
  who took the statement. A map pin's link highlights its card in mustard.
- **Tracker**: once the trip has scrolled away, a pill in the bottom-right corner with Ainsley's
  face and a peeking lion, linking back to the trip, so he's always on screen.
- **Footer**: ink, paper text, mustard links.

## Rules

- Ainsley is always visible: on the stage during the trip, in the tracker after it.
- No gradients and no blurred shadows; the dotted sky is the only texture.
- Every motion stops under `prefers-reduced-motion`; the captions read fine without JavaScript.

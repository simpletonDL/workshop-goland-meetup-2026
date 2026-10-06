// block: map — copy to web/assets/js/map.js. See map.md.
//
// Draws every [data-map] on the page with Leaflet, reading its points from
// the JSON the component wrote out. Markers always go on a canvas, since
// the home page has thousands; clusters and heat are opt-in per map and
// used only when their plugin loaded. Without Leaflet the fallback shows.

(function () {
	const TILES = 'https://tile.openstreetmap.org/{z}/{x}/{y}.png';
	const CREDIT = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors';
	const still = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

	function points(id) {
		const el = document.getElementById(id);
		try {
			return (el && JSON.parse(el.textContent)) || [];
		} catch (err) {
			return [];
		}
	}

	// Built from nodes, never an HTML string: every field here is GBIF's text.
	function popup(point) {
		const root = document.createElement('div');
		root.className = 'map__popup';
		root.appendChild(document.createElement('strong')).textContent = point.title;
		root.appendChild(document.createElement('span')).textContent =
			[point.date, point.place, point.count].filter(Boolean).join(' · ');
		if (point.by) {
			root.appendChild(document.createElement('span')).textContent = 'Statement taken by ' + point.by;
		}
		if (point.anchor) {
			const link = root.appendChild(document.createElement('a'));
			link.href = '#' + point.anchor;
			link.textContent = 'Read the statement';
		}
		return root;
	}

	const ESCAPES = [
		'Gone! He was last seen near',
		'Too slow. He legged it to',
		'He heard a lion and ran to',
		'A buffalo gave him a lift to',
		'He asked for directions and ended up in',
	];

	function pick(list) {
		return list[Math.floor(Math.random() * list.length)];
	}

	// Ainsley's face sits on a random buffalo, and hops to another when clicked.
	function lastSeen(map, el, data) {
		const img = document.createElement('img');
		img.src = el.dataset.lastSeen;
		img.alt = '';
		const icon = L.divIcon({ className: 'map__ainsley', html: img, iconSize: [56, 56], iconAnchor: [28, 28] });

		const start = pick(data);
		const marker = L.marker([start.lat, start.lng], { icon: icon, zIndexOffset: 1000, title: 'Ainsley, last seen here' })
			.addTo(map);

		marker.on('click', () => {
			const next = pick(data);
			marker.setLatLng([next.lat, next.lng]);
			const note = document.createElement('div');
			note.className = 'map__popup';
			note.appendChild(document.createElement('strong')).textContent = 'Ainsley';
			note.appendChild(document.createElement('span')).textContent = pick(ESCAPES) + ' ' + next.place + '.';
			L.popup({ offset: [0, -24] }).setLatLng([next.lat, next.lng]).setContent(note).openOn(map);
			map.panTo([next.lat, next.lng], { animate: !still });
		});
	}

	// Colours come from the element's CSS, so the palette lives in one place.
	function colour(el, prop, fallback) {
		return getComputedStyle(el).getPropertyValue(prop).trim() || fallback;
	}

	function init(el) {
		const data = points(el.dataset.map);
		if (!data.length || typeof L === 'undefined') {
			el.classList.add('is-unavailable');
			return;
		}

		const map = L.map(el, {
			preferCanvas: true,
			scrollWheelZoom: false,
			zoomSnap: 0.25,
			zoomAnimation: !still,
			fadeAnimation: !still,
			markerZoomAnimation: !still,
			worldCopyJump: true,
		});
		L.tileLayer(el.dataset.tiles || TILES, {
			attribution: el.dataset.attribution || CREDIT,
			maxZoom: 18,
			minZoom: 1,
		}).addTo(map);

		const latLngs = data.map((p) => [p.lat, p.lng]);
		let layer;
		if ('heat' in el.dataset && typeof L.heatLayer === 'function') {
			layer = L.heatLayer(latLngs, { radius: 18, blur: 15, minOpacity: 0.3 });
		} else {
			const fill = colour(el, '--map-marker', '#c2410c');
			const stroke = colour(el, '--map-marker-stroke', '#fff');
			layer = 'cluster' in el.dataset && typeof L.markerClusterGroup === 'function'
				? L.markerClusterGroup({ chunkedLoading: true, showCoverageOnHover: false })
				: L.featureGroup();
			data.forEach((p) => {
				L.circleMarker([p.lat, p.lng], { radius: 6, weight: 2, color: stroke, fillColor: fill, fillOpacity: 0.8 })
					// A function, so the popup is only built when it opens.
					.bindPopup(() => popup(p))
					.addTo(layer);
			});
		}
		layer.addTo(map);

		if (el.dataset.lastSeen) {
			lastSeen(map, el, data);
		}

		map.fitBounds(L.latLngBounds(latLngs), {
			padding: [24, 24],
			animate: false,
			maxZoom: Number(el.dataset.maxZoom) || 12,
		});

		// Wheel zoom waits for a click or focus, so the page scrolls past freely.
		map.on('focus', () => map.scrollWheelZoom.enable());
		map.on('blur', () => map.scrollWheelZoom.disable());
	}

	document.querySelectorAll('[data-map]').forEach(init);
})();

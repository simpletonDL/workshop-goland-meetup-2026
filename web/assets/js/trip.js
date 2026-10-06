// Switches the stage to whichever chapter is in the middle of the screen,
// lets the reader poke Ainsley, and brings in the corner tracker once the
// stage has scrolled away. Without it the first scene stays and every
// caption still reads in order.

(function () {
	const trip = document.querySelector('.trip');
	if (!trip) return;

	const chapters = trip.querySelectorAll('.chapter');
	const casts = trip.querySelectorAll('.cast');
	const trail = trip.querySelector('.trail');
	const links = trip.querySelectorAll('.trail a');
	const ainsley = trip.querySelector('.ainsley');
	const say = trip.querySelector('.ainsley-say');
	const tracker = document.querySelector('.tracker');

	const POKES = [
		'Ow.',
		'I only went out for a flat white!',
		'Does anyone have Wi-Fi?',
		'Is this the GoLand meetup?',
		"I'm not food. I'm a developer.",
		'Has anyone seen a gopher?',
		'My talk starts in five minutes!',
		'Please stop poking me.',
		'I think the buffalo like me.',
		'Can somebody call JetBrains?',
	];

	// Restarts a CSS animation by taking its class off and putting it back.
	function replay(el, className) {
		el.classList.remove(className);
		void el.offsetWidth;
		el.classList.add(className);
	}

	function speak(text) {
		say.textContent = text;
		replay(say, 'is-popping');
	}

	function show(chapter) {
		const id = chapter.dataset.scene;
		if (trip.dataset.scene === id) return;

		trip.dataset.scene = id;
		trip.dataset.pose = chapter.dataset.pose || 'stand';
		trip.style.setProperty('--sky', 'var(--' + chapter.dataset.sky + ')');
		trip.style.setProperty('--ground', 'var(--' + chapter.dataset.ground + ')');
		casts.forEach((cast) => cast.classList.toggle('is-active', cast.dataset.cast === id));

		links.forEach((link) => {
			if (link.hash === '#' + id) {
				link.setAttribute('aria-current', 'step');
				trail.scrollTo({ left: link.offsetLeft - trail.clientWidth / 2 + link.offsetWidth / 2 });
			} else {
				link.removeAttribute('aria-current');
			}
		});

		speak(chapter.dataset.ainsley || '');
	}

	const watcher = new IntersectionObserver((entries) => {
		entries.forEach((entry) => {
			if (entry.isIntersecting) show(entry.target);
		});
	}, { rootMargin: '-50% 0px -50% 0px' });
	chapters.forEach((chapter) => watcher.observe(chapter));

	ainsley.addEventListener('click', () => {
		const lines = POKES.filter((line) => line !== say.textContent);
		speak(lines[Math.floor(Math.random() * lines.length)]);
		replay(ainsley, 'is-poked');
	});
	ainsley.addEventListener('animationend', () => ainsley.classList.remove('is-poked'));

	if (tracker) {
		new IntersectionObserver(([entry]) => {
			tracker.classList.toggle('is-visible', !entry.isIntersecting);
		}).observe(trip);
	}
})();

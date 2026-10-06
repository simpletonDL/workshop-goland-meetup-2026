package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

type SightingsLister interface {
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

func Home(logger *slog.Logger, lister SightingsLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sightings, err := lister.List(ctx, sighting.ListFilter{})
		if err != nil {
			logger.ErrorContext(ctx, "handlers: homepage: listing sightings: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Sightings couldn't be loaded.")
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// The status is already sent once rendering starts, so a failure
		// here can only be logged.
		if err := pages.Home(sightings).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: homepage: rendering: "+err.Error())
		}
	}
}

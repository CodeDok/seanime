package mediastream

import (
	"net/http"
	"net/url"
	"path/filepath"
	"seanime/internal/events"
	"seanime/internal/mediastream/videofile"

	"github.com/labstack/echo/v4"
)

func (r *Repository) ServeEchoExtractedSubtitles(c echo.Context) error {

	if !r.IsInitialized() {
		r.wsEventManager.SendEvent(events.MediastreamShutdownStream, "Module not initialized")
		return echo.NewHTTPError(http.StatusServiceUnavailable, "module not initialized")
	}

	if !r.TranscoderIsInitialized() {
		r.wsEventManager.SendEvent(events.MediastreamShutdownStream, "Transcoder not initialized")
		return echo.NewHTTPError(http.StatusServiceUnavailable, "transcoder not initialized")
	}

	// Get the parameter group
	subFilePath := c.Param("*")

	// Get current media
	mediaContainer, found := r.playbackManager.currentMediaContainer.Get()
	if !found {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "no file has been loaded")
	}

	retPath := videofile.GetFileSubsCacheDir(r.cacheDir, mediaContainer.Hash)

	if retPath == "" {
		return echo.NewHTTPError(http.StatusNotFound, "could not find subtitles")
	}

	r.logger.Trace().Msgf("mediastream: Serving subtitles from %s", retPath)

	return c.File(filepath.Join(retPath, subFilePath))
}

func (r *Repository) ServeEchoExtractedAttachments(c echo.Context) error {
	if !r.IsInitialized() {
		r.wsEventManager.SendEvent(events.MediastreamShutdownStream, "Module not initialized")
		return echo.NewHTTPError(http.StatusServiceUnavailable, "module not initialized")
	}

	if !r.TranscoderIsInitialized() {
		r.wsEventManager.SendEvent(events.MediastreamShutdownStream, "Transcoder not initialized")
		return echo.NewHTTPError(http.StatusServiceUnavailable, "transcoder not initialized")
	}

	// Get the parameter group
	subFilePath := c.Param("*")

	// Get current media
	mediaContainer, found := r.playbackManager.currentMediaContainer.Get()
	if !found {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "no file has been loaded")
	}

	retPath := videofile.GetFileAttCacheDir(r.cacheDir, mediaContainer.Hash)

	if retPath == "" {
		return echo.NewHTTPError(http.StatusNotFound, "could not find attachments")
	}

	subFilePath, _ = url.PathUnescape(subFilePath)

	return c.File(filepath.Join(retPath, subFilePath))
}

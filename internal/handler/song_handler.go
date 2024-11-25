package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createSong(ctx *gin.Context) {
	logrus.Debug("Entering createSong handler")
	var input domain.SongList
	if err := ctx.BindJSON(&input); err != nil {
		logrus.Warnf("Failed to bind JSON in createSong: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	logrus.Infof("Creating song with input: %+v", input)
	id, err := h.services.Song.Create(input)
	if err != nil {
		logrus.Errorf("Failed to create song: %v", err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully created song with ID: %d", id)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSong(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf("Failed to convert song ID to integer: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	logrus.Infof("Deleting song with ID: %d", songId)
	err = h.services.Song.Delete(songId)
	if err != nil {
		logrus.Errorf("Failed to delete song with ID %d: %v", songId, err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully deleted song with ID: %d", songId)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}

func (h *Handler) updateSong(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf("Failed to convert song ID to integer: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	var input domain.SongUpdateInput
	if err := ctx.BindJSON(&input); err != nil {
		logrus.Warnf("Failed to bind JSON in updateSong: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	logrus.Infof("Updating song with ID: %d", songId)
	err = h.services.Song.Update(songId, input)
	if err != nil {
		logrus.Errorf("Failed to update song with ID %d: %v", songId, err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully updated song with ID: %d", songId)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Status": "OK",
	})
}

func (h *Handler) getSongsList(ctx *gin.Context) {
	logrus.Debug("Entering getSongsList handler")

	input := domain.PaginatedSongInput{
		GroupName:   ctx.DefaultQuery("group", ""),
		Title:       ctx.DefaultQuery("song", ""),
		Text:        ctx.DefaultQuery("text", ""),
		ReleaseDate: ctx.DefaultQuery("release_date", ""),
		Link:        ctx.DefaultQuery("link", ""),
	}

	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "5")

	logrus.Infof("Fetching songs list with filters: %+v, page: %s, page size: %s", input, page, pageSize)
	filteredSongs, err := h.services.Song.GetSongsList(input, page, pageSize)
	if err != nil {
		logrus.Errorf("Failed to fetch songs list: %v", err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully fetched %d songs", len(filteredSongs))
	response := domain.PaginatedSongResponse{
		Data:       filteredSongs,
		TotalCount: len(filteredSongs),
		Page:       input.Page,
		PageSize:   input.PageSize,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) getSongText(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf("Failed to convert song ID to integer: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	logrus.Infof("Fetching song text for song ID: %d", songId)

	input := domain.PaginatedSongTextInput{
		SongId: songId,
	}

	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "5")

	logrus.Infof("Fetching song text with pagination: page=%s, page_size=%s", page, pageSize)
	songText, err := h.services.Song.GetSongText(input, page, pageSize)
	if err != nil {
		logrus.Errorf("Failed to fetch song text for song ID %d: %v", songId, err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully fetched song text for song ID: %d", songId)
	ctx.JSON(http.StatusOK, songText)
}

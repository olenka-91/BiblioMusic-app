package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createSong(ctx *gin.Context) {
	var input domain.SongList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request")
	}

	id, err := h.services.Song.Create(input)
	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error()) //"Internal server error")
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getSongsList(ctx *gin.Context) {
	input := domain.PaginatedSongInput{
		GroupName:   ctx.DefaultQuery("group", ""),
		Title:       ctx.DefaultQuery("song", ""),
		Text:        ctx.DefaultQuery("text", ""),
		ReleaseDate: ctx.DefaultQuery("release_date", ""),
		Link:        ctx.DefaultQuery("link", ""),
		Page:        0,
		PageSize:    0,
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "5"))
	if err != nil || pageSize < 1 {
		pageSize = 5
	}
	logrus.Debug("pageSize=", pageSize, " page=", page)
	input.Page = page
	input.PageSize = pageSize

	filteredSongs, err := h.services.Song.GetSongsList(input)
	if err != nil {
		newErrorResponce(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := domain.PaginatedSongResponse{
		Data:       filteredSongs,
		TotalCount: len(filteredSongs),
		Page:       page,
		PageSize:   pageSize,
	}

	ctx.JSON(http.StatusOK, response)
}

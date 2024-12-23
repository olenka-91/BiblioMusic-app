package handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/domain"
	"github.com/sirupsen/logrus"
)

func enrichSongData(input domain.AddSongRequest) (domain.SongDetail, error) {
	baseURL := os.Getenv("EXTERNAL_HTTP_ADDR")
	if baseURL == "" {
		logrus.Warnf("Failed to get EXTERNAL_HTTP_ADDR")
		return domain.SongDetail{}, errors.New("EXTERNAL_HTTP_ADDR is not configured")
	}

	logrus.Debugf("Get baseURL for EXTERNAL API successfuly %s", baseURL)

	req, err := http.NewRequest("GET", baseURL+"/info", nil)
	if err != nil {
		return domain.SongDetail{}, err
	}

	q := req.URL.Query()
	q.Add("group", input.Group)
	q.Add("song", input.Song)
	req.URL.RawQuery = q.Encode()

	logrus.Debugf("Request for EXTERNAL API is %s", req.URL.RawQuery)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.SongDetail{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.SongDetail{}, fmt.Errorf("external API returned status code %d", resp.StatusCode)
	}

	var songDetail domain.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return domain.SongDetail{}, err
	}

	return songDetail, nil
}

// CreateSong godoc
// @Summary Добавление песни в БД
// @Description Добавление новой песни в БД с запросом доп информации на сторонний сервер
// @Tags songs
// @Accept json
// @Produce json
// @Param song body domain.AddSongRequest true "Song data"
// @Success 200 {object} Response "Song created successfully"
// @Failure 400 {object} ErrorResponce "Invalid input"
// @Failure 500 {object} ErrorResponce "Internal server error"
// @Router /song [post]
func (h *Handler) createSong(ctx *gin.Context) {
	logrus.Debug("Entering createSong handler")
	var input domain.AddSongRequest
	if err := ctx.BindJSON(&input); err != nil {
		logrus.Warnf("Failed to bind JSON in createSong: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	songDetail, err := enrichSongData(input)
	if err != nil {
		logrus.Warnf("Failed to bind JSON in createSong: %v", err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully received extra data from EXTERNAL API: Text=%s, ReleaseDate=%s, Link=%s",
		domain.StringValue(songDetail.Text),
		domain.StringValue(songDetail.ReleaseDate),
		domain.StringValue(songDetail.Link))

	id, err := h.services.Song.Create(input, songDetail)
	if err != nil {
		logrus.Errorf("Failed to create song: %v", err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully created song with ID: %d", id)
	ctx.JSON(http.StatusOK, Response{
		ID: id,
	})
}

// DeleteSong godoc
// @Summary Удаление пепсни по ID
// @Description Удаление записи из таблицы песен БД по ее ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Success 200 {object} StatusResponce "Successfully deleted song"
// @Failure 400 {object} ErrorResponce "Invalid song ID"
// @Failure 500 {object} ErrorResponce "Internal server error"
// @Router /songs/{id} [delete]
func (h *Handler) deleteSong(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf("Failed to convert song ID to integer: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	err = h.services.Song.Delete(songId)
	if err != nil {
		logrus.Errorf("Failed to delete song with ID %d: %v", songId, err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully deleted song with ID: %d", songId)
	ctx.JSON(http.StatusOK, StatusResponce{
		Status: "OK",
	})
}

// UpdateSong godoc
// @Summary Обновление песни по ее ID
// @Description Обновление информации о песне по ее ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body domain.SongUpdateInput true "Updated song data"
// @Success 200 {object} StatusResponce "Song updated successfully"
// @Failure 400 {object} ErrorResponce "Invalid input"
// @Failure 500 {object} ErrorResponce "Internal server error"
// @Router /songs/{id} [patch]
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

	err = h.services.Song.Update(songId, input)
	if err != nil {
		logrus.Errorf("Failed to update song with ID %d: %v", songId, err)
		newErrorResponce(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	logrus.Infof("Successfully updated song with ID: %d", songId)
	ctx.JSON(http.StatusOK, StatusResponce{
		Status: "OK",
	})
}

// GetSongsList godoc
// @Summary Получение списка песен
// @Description Получение списка песен с опциональной фильтрацией результатов по полям: group, song, text, release_date, link и пагинацией.
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song title"
// @Param text query string false "Song text"
// @Param release_date query string false "Release date"
// @Param link query string false "Song link"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(5)
// @Success 200 {object} domain.PaginatedSongResponse "List of songs"
// @Failure 500 {object} ErrorResponce "Internal server error"
// @Router /songs [get]
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

// GetSongText godoc
// @Summary Получение текста песни
// @Description Получение текста песни по ее ID с пагинацией по куплетам
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(5)
// @Success 200 {object} domain.PaginatedSongTextResponse "Song text data"
// @Failure 400 {object} ErrorResponce "Invalid song ID"
// @Failure 500 {object} ErrorResponce "Internal server error"
// @Router /songs/{id}/text [get]
func (h *Handler) getSongText(ctx *gin.Context) {
	songId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.Warnf("Failed to convert song ID to integer: %v", err)
		newErrorResponce(ctx, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

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

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{services: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/song", h.createSong)
	router.GET("/info", h.getSongsList)
	router.GET("/songs/:id/text", h.getSongText)
	//router.PATCH("/songs/:id", h.updateSong)
	//router.DELETE("/songs/:id", h.deleteSong)

	return router

}

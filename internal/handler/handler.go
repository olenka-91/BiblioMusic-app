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

	song := router.Group("/info")
	{
		song.POST("/group/:groupName/song/:songName", h.createSong)
		// song.POST("/", h.createSong)
		// song.GET("/", h.getAllReminds)
		// song.GET("/:id", h.getRemindByID)
		// song.PUT("/:id", h.updateRemind)
		// song.DELETE("/:id", h.deleteRemind)
	}

	return router

}

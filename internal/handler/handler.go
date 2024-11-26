package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/olenka-91/BIBLIOMUSIC-APP/internal/service"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/olenka-91/BIBLIOMUSIC-APP/docs"
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
	router.GET("/songs", h.getSongsList)
	router.GET("/songs/:id/text", h.getSongText)
	router.PATCH("/songs/:id", h.updateSong)
	router.DELETE("/songs/:id", h.deleteSong)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router

}

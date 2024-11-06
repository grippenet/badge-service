package server

import(
	"github.com/gin-gonic/gin"
	"github.com/grippenet/badge-service/pkg/types"
)

type HttpServer struct {
	conf types.HttpConfig
	services types.BadgeServices
}

func NewHttpServer(cfg types.HttpConfig, services types.BadgeServices) *HttpServer {
	return &HttpServer{
		conf: cfg,
		services: services,
	}
}

func healthCheckHandle(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func (h *HttpServer) Start() error {
	router := gin.Default()
	router.GET("/", healthCheckHandle)

	if(h.services.Pioneer != nil) {
		handler := pioneerHandler{service: h.services.Pioneer}
		router.POST("/pioneer", handler.Handle)
	}

	return router.Run()
}
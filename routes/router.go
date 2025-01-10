package routes

import (
	"github.com/gin-gonic/gin"
)

func BootstrapRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	HealthCheck(v1)
	User(v1)
	Relation(v1)
	return router
}

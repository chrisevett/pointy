package routers

import (
	controllers "chrisevett/pointy/v1/controllers/v1/title"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Server entry point
func Run() {
	getRoutes()
	// todo: set this via the environment/config
	router.Run(":5000")
}

func getRoutes() {
	v1 := router.Group("/v1")
	// add
	addTitleRoutes(v1)
}

// todo add controllers
// todo add favicon, make some pixel art for it
func addTitleRoutes(rg *gin.RouterGroup) {
	titles := rg.Group("/titles")

	titles.GET("/", controllers.TitleGetAll)
	titles.GET("/:id", controllers.TitleGetOne)
	titles.POST("/", controllers.TitleCreate)
	titles.PUT("/:id", controllers.TitleUpdate)
	titles.DELETE("/:id", controllers.TitleUpdate)

}

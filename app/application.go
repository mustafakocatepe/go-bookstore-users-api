package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() //Default returns an Engine instance with the Logger and Recovery middleware already attached.
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}

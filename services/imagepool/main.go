package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/eskpil/salmon/services/imagepool/mycontext"
	"github.com/gin-gonic/gin"

	imageController "github.com/eskpil/salmon/services/imagepool/controllers/images"
	metadataController "github.com/eskpil/salmon/services/imagepool/controllers/metadata"
)

func main() {
	ctx, err := mycontext.NewContext()

	if err != nil {
		log.Fatalf("Failed to create a new context: %v\n", err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/metadata/", metadataController.CreateOne(ctx))
	r.GET("/api/metadata/:metadataId/", metadataController.GetOne(ctx))

	r.POST("/api/images/:metadataId/", imageController.CreateOne(ctx))
	r.GET("/api/images/:metadataId/", imageController.GetOne(ctx))

	r.Run("0.0.0.0:8091")
}

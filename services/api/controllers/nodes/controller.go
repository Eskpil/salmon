package nodes

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/eskpil/salmon/services/api/mycontext"
	"github.com/gin-gonic/gin"

	nodeService "github.com/eskpil/salmon/services/api/services/nodes"
)

func GetById(s *mycontext.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")

		node, err := nodeService.GetById(ctx, s, id)

		if err != nil {
			log.Error(err)
			c.String(403, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, node)
	}
}

func GetAll(m *mycontext.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		nodes, err := nodeService.GetAll(ctx, m)

		if err != nil {
			log.Error(err)
			c.String(403, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, nodes)
	}
}

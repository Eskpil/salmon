package machines

import (
	"context"
	"net/http"
	"time"

	"github.com/eskpil/salmon/services/api/mycontext"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	machineService "github.com/eskpil/salmon/services/api/services/machines"
)

func GetAll(m *mycontext.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		machines, err := machineService.GetAll(ctx, m)

		if err != nil {
			log.Error(err)
			c.String(403, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, machines)
	}
}

func GetById(m *mycontext.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := c.Param("id")

		machine, err := machineService.GetById(ctx, m, id)

		if err != nil {
			log.Error(err)
			c.String(403, err.Error())
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, machine)
	}
}

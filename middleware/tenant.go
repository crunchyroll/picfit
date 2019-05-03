package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thoas/picfit/logger"
	"net/http"
	"github.com/thoas/picfit/storage"
	"github.com/thoas/picfit/application"
	"github.com/thoas/picfit/config"
	"github.com/thoas/picfit/kvstore"
)

func TenantParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenant := c.Request.Header.Get("X-Tenant")
		log := logger.FromContext(c)
		if tenant != "" {
			log.Infof("X-Tenant header: %s", tenant)

			path := "bin/config_" + tenant + ".json";
			cfg, err := config.Load(path)
			if err != nil {
				log.Error("can't load new config")
				c.String(http.StatusInternalServerError, "can't load new config")
				c.Abort()
				return
			}

			ctx, err := application.Load(cfg)
			if err != nil {
				log.Error("can't load new context")
				c.String(http.StatusInternalServerError, "can't load new context")
				c.Abort()
				return
			}

			destination := storage.DestinationFromContext(ctx)
			c.Set("dstStorage", destination)
			log.Infof("dest: %s", destination)

			source := storage.SourceFromContext(ctx)
			c.Set("srcStorage", source)
			log.Infof("source: %s", source)

			redis := kvstore.FromContext(ctx)
			c.Set("kvstore", redis)
			log.Infof("redis: %s", redis)

		} else {
			log.Error("empty X-Tenant header")
			c.String(http.StatusBadRequest, "empty X-Tenant header")
			c.Abort()
			return
		}

		c.Next()
	}
}

package httpapp

import (
	processhandle "app/cmd/merge-blob/delivery/http/process"
	streamhandle "app/cmd/merge-blob/delivery/http/stream"
	"app/internal/connection"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Register() http.Handler {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	processhandle.Register(r)
	streamhandle.Register(r)

	log.Printf(
		"Server h-learning-merge-blob-service starting success! URL: http://%s:%s",
		connection.GetConnect().MergeBlobSevice.Host,
		connection.GetConnect().MergeBlobSevice.Port,
	)

	return r
}

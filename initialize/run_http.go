package initialize

import (
	httpapp "app/cmd/merge-blob-service/delivery/http"
	"app/internal/connection"
	"fmt"
	"log"
	"net/http"
	"time"
)

func runHttpServer() {
	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s",
			connection.GetConnectionInfo().MergeBlobSevice.Host,
			connection.GetConnectionInfo().MergeBlobSevice.Port,
		),
		Handler:      httpapp.Register(),
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 300 * time.Second,
	}
	log.Fatalln(s.ListenAndServe())
}

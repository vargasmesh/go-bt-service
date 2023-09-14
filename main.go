package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	httpServer.ListenAndServe()
}

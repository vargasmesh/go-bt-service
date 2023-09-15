package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vargasmesh/go-bt-service/internal/server"
)

func main() {
	router := gin.New()
	treeServer := server.NewTreeServer()

	router.GET("/tree", func(c *gin.Context) {
		c.JSON(http.StatusOK, treeServer.GetPreOrderTree())
	})

	router.POST("/tree", func(c *gin.Context) {
		var value int
		if err := c.ShouldBindJSON(&value); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		treeServer.Insert(value)
		c.JSON(http.StatusOK, fmt.Sprintf("Inserted %d", value))
	})

	ctx, cancel := context.WithCancel(context.Background())

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	go treeServer.Run(ctx)
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	cancel()
}

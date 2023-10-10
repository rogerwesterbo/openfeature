package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

const defaultMessage = "Hello!"

func main() {
	cancelChan := make(chan os.Signal, 1)
	stop := make(chan struct{})
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	_, _ = fmt.Fprintln(os.Stdout, "🚀 Starting server...")

	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("Test server 😊"))
		return
	})
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, defaultMessage)
		return
	})

	go func() {
		_ = engine.Run("localhost:7357")
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	<-cancelChan
	_, _ = fmt.Fprintln(os.Stdout, "\n⛔ Abort signal detected\nShutting down gracefully 😎")
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

const defaultMessage = "Hello hello!"
const newWelcomeMessage = "Hello, welcome to this OpenFeature-enabled website!"

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("API_PORT", 7357)
	viper.SetDefault("API_HOST", "")

	viper.SetDefault("FLAGD_HOST", "localhost")
	viper.SetDefault("FLAGD_PORT", 8013)
}

func main() {
	cancelChan := make(chan os.Signal, 1)
	stop := make(chan struct{})
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	_, _ = fmt.Fprintln(os.Stdout, "🚀 Starting server...")

	// Use flagd as the OpenFeature provider
	flagdHost := viper.GetString("FLAGD_HOST")
	flagdPort := viper.GetUint16("FLAGD_PORT")
	provider := flagd.NewProvider(
		flagd.WithHost(flagdHost),
		flagd.WithPort(flagdPort))
	err := openfeature.SetProvider(provider)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error setting OpenFeature provider: %s\n", err)
		os.Exit(1)
	}

	// Initialize OpenFeature openFeatureClient
	openFeatureClient := openfeature.NewClient("GoStartApp")

	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("Test server 😊"))
		return
	})
	engine.GET("/hello", func(c *gin.Context) {
		// Evaluate welcome-message feature flag
		welcomeMessage, _ := openFeatureClient.BooleanValue(
			context.Background(), "welcome-message", false, openfeature.EvaluationContext{},
		)

		if welcomeMessage {
			c.JSON(http.StatusOK, newWelcomeMessage)
			return
		} else {
			c.JSON(http.StatusOK, defaultMessage)
			return
		}
	})

	go func() {
		apiHost := viper.GetString("API_HOST")
		apiPort := viper.GetString("API_PORT")

		if apiHost != "" && apiPort != "" {
			apiAddress := fmt.Sprintf("%s:%s", apiHost, apiPort)
			_ = engine.Run(apiAddress)
		} else {
			_ = engine.Run()
		}

		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	<-cancelChan
	_, _ = fmt.Fprintln(os.Stdout, "\n⛔ Abort signal detected\nShutting down gracefully 😎")
}

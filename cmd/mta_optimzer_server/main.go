package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service"
)

func main() {
	var env string
	env, err := GetEnvironment()
	if err != nil {
		return
	}
	config := LoadConfig(env)
	thresholdValue := os.Getenv("X")
	if thresholdValue == "" {
		err := os.Setenv("X", "1")
		if err != nil {
			return
		}
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableKeepAlives:     true,
		},
	}
	mtaServiceInitializer := mta_hosting_optimizer_service.NewMtaHostingOptimizerService(httpClient, config.DataServerUrl)
	mtaHostingOptimizerService := gin.Default()

	mtaHostingOptimizerService.GET("/", func(context *gin.Context) {
		resp, err := mtaServiceInitializer.GetUnderUtilizedHostName(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, resp.HostNames)
		}
	})
	err = mtaHostingOptimizerService.Run(fmt.Sprintf(":%d", config.MtaServicePort))
	if err != nil {
		fmt.Printf("error in running mtaHostingOptimizerService")
	}
}

func GetEnvironment() (string, error) {
	var ok bool
	var envLoadErr error
	envName, ok := os.LookupEnv("ENVIRONMENT")

	if !ok {
		envLoadErr = fmt.Errorf("env var `ENVIRONMENT` is not set")
	}
	return envName, envLoadErr
}

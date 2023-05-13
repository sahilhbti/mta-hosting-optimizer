package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service"
)

func main() {
	err := os.Setenv("X", "1")
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dataServiceInitializer := data_service.NewDataService()
		dataService := gin.Default()
		dataService.GET("/", func(context *gin.Context) {
			serverDetail, err := dataServiceInitializer.GetServerDetails(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, serverDetail)
			}
		})
		err := dataService.Run(":443")
		if err != nil {
			fmt.Printf("error in running DataService")
		}
	}()

	go func() {
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
		mtaServiceInitializer := mta_hosting_optimizer_service.NewMtaHostingOptimizerService(httpClient)
		mtaHostingOptimizerService := gin.Default()

		mtaHostingOptimizerService.GET("/", func(context *gin.Context) {
			resp, err := mtaServiceInitializer.GetUnderUtilizedHostName(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, resp.HostNames)
			}
		})
		err := mtaHostingOptimizerService.Run(":9518")
		if err != nil {
			fmt.Printf("error in running mtaHostingOptimizerService")
		}
	}()
	wg.Wait()
}

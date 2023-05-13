package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service"
)

func main() {
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
	err := dataService.Run(":80")
	if err != nil {
		fmt.Printf("error in running DataService")
	}
}

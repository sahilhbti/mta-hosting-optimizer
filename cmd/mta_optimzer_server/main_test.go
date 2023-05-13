//go:build integration && !unit_test

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service"
)

func setupServer() *gin.Engine {
	gin.EnableJsonDecoderUseNumber()
	mtaHostingOptimizerService := gin.Default()
	env, err := GetEnvironment()
	if err != nil {
		return nil
	}
	config := LoadConfig(env)
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

	// version 1 async routes
	v1 := mtaHostingOptimizerService.Group("")
	{
		v1.GET("/", func(context *gin.Context) {
			resp, err := mtaServiceInitializer.GetUnderUtilizedHostName(context)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, resp.HostNames)
			}
		})
	}

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
	go func() {
		err = dataService.Run(":9090")
		if err != nil {
			fmt.Println(err)
		}
	}()

	return mtaHostingOptimizerService
}

func runTestServer() *httptest.Server {

	return httptest.NewServer(setupServer())
}

func Test_get_under_utilized_host_name(t *testing.T) {
	err := os.Setenv("X", "1")
	if err != nil {
		return
	}
	err = os.Setenv("ENVIRONMENT", "test")
	if err != nil {
		return
	}
	ts := runTestServer()
	defer ts.Close()
	t.Run("it should return ok ", func(t *testing.T) {
		resp, err := http.Get(ts.URL)
		if err != nil {
			fmt.Println(err)
		}
		if err != nil || resp == nil || resp.StatusCode != 200 {
			t.Errorf(" error = %v", err)
		}
		defer func() {
			if resp != nil && resp.Body != nil {
				_ = resp.Body.Close()
			}
		}()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		var hostNames []string
		err = json.Unmarshal(body, &hostNames)
		if err != nil {
			t.Errorf(" error = %v", err)
		}
		if len(hostNames) == 2 && ((hostNames[0] == "mta-prod-1" && hostNames[1] == "mta-prod-3") || (hostNames[1] == "mta-prod-1" && hostNames[0] == "mta-prod-3")) {
			fmt.Println("integration test passed")
			return
		}
		t.Errorf(" error = %v", err)
	})

}

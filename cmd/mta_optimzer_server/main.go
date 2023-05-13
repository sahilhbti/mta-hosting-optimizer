package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service"
)

var (
	_, b, _, _ = runtime.Caller(0)
)

type Config struct {
	Environment    string `yaml:"Environment"`
	DataServerUrl  string `yaml:"DataServerUrl"`
	MtaServicePort int    `yaml:"MtaServicePort"`
}

func main() {
	var env string
	env, err := GetEnvironment()
	if err != nil {
		log.Fatalf("error in getting env %s", err)
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

func LoadConfig(env string) *Config {
	var config Config
	fileName := "/mta-" + env + ".yml"
	configPath := filepath.Join(b, "..")
	configFilepath := configPath + fileName
	yamlFile, err := ioutil.ReadFile(configFilepath)
	if err != nil || yamlFile == nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &config

}

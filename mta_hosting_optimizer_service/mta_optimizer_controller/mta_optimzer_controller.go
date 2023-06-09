package mta_optimizer_controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail"
)

//go:generate mockgen -source=$PWD/mta_optimzer_controller.go -destination=$PWD/mocks/mta_optimzer_controller.go -package=mocks
type IMtaOptimizerController interface {
	GetUnderUtilizedHost(*gin.Context, *GetUnderUtilizedHostRequest) (*GetUnderUtilizedHostResponse, error)
}

type GetUnderUtilizedHostRequest struct {
}

type GetUnderUtilizedHostResponse struct {
	HostNames []string
}

type MtaOptimizerController struct {
	HttpClient    *http.Client
	ServersDetail server_detail.IServersDetail
}

func NewMtaOptimizerController(httpClient *http.Client, url string) *MtaOptimizerController {
	return &MtaOptimizerController{
		HttpClient: httpClient,
		ServersDetail: &server_detail.ServersDetail{
			HttpClient: httpClient,
			Url:        url,
		},
	}
}

func (m *MtaOptimizerController) GetUnderUtilizedHost(ctx *gin.Context, request *GetUnderUtilizedHostRequest) (*GetUnderUtilizedHostResponse, error) {
	var hostNames []string
	hostNameToActiveCount := make(map[string]int)
	serversDetailResp, err := m.ServersDetail.GetServersDetail(ctx, &server_detail.GetServersDetailRequest{})
	if err != nil {
		return nil, fmt.Errorf("error in getting server detail %s", err)
	}
	serversDetail := serversDetailResp.ServerData
	thresholdValue, err := strconv.Atoi(os.Getenv("X"))
	if err != nil {
		return nil, fmt.Errorf("error in parsing threshold %s", err)
	}
	fmt.Println(thresholdValue)
	for _, eachServerDetail := range serversDetail {
		if eachServerDetail.Active == true {
			hostNameToActiveCount[eachServerDetail.HostName]++
		} else {
			hostNameToActiveCount[eachServerDetail.HostName] = hostNameToActiveCount[eachServerDetail.HostName] + 0
		}
	}
	for hostName, activeIpAddressesCount := range hostNameToActiveCount {
		if activeIpAddressesCount <= thresholdValue {
			hostNames = append(hostNames, hostName)
		}
	}
	return &GetUnderUtilizedHostResponse{
		HostNames: hostNames,
	}, nil
}

package mta_optimizer_controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail"
)

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

func NewMtaOptimizerController(httpClient *http.Client) *MtaOptimizerController {
	return &MtaOptimizerController{
		HttpClient: httpClient,
		ServersDetail: &server_detail.ServersDetail{
			HttpClient: httpClient,
		},
	}
}

func (m *MtaOptimizerController) GetUnderUtilizedHost(ctx *gin.Context, request *GetUnderUtilizedHostRequest) (*GetUnderUtilizedHostResponse, error) {
	var hostNames []string
	hostNameToActiveCount := make(map[string]int)
	serversDetailresp, err := m.ServersDetail.GetServersDetail(ctx, &server_detail.GetServersDetailRequest{})
	if err != nil {
		return nil, fmt.Errorf("error in getting server detail %s", err)
	}
	serversDetail := serversDetailresp.ServerData
	thresholdValue, err := strconv.Atoi(os.Getenv("X"))
	if err != nil {
		return nil, fmt.Errorf("error in parsing threshold", err)
	}
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

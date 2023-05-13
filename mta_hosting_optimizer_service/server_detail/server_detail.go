package server_detail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service/models"
)

//go:generate mockgen -source=$PWD/server_detail.go -destination=$PWD/mocks/server_detail.go -package=mocks
type IServersDetail interface {
	GetServersDetail(ctx *gin.Context, request *GetServersDetailRequest) (*GetServersDetailResponse, error)
}

type GetServersDetailRequest struct {
}

type GetServersDetailResponse struct {
	ServerData []*models.ServerData
}

type ServersDetail struct {
	HttpClient *http.Client
}

func (s *ServersDetail) GetServersDetail(ctx *gin.Context, request *GetServersDetailRequest) (*GetServersDetailResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://ec2-3-111-218-77.ap-south-1.compute.amazonaws.com", nil)
	if err != nil {
		return nil, fmt.Errorf("error in building request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error parser request: %w", err)
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading response: %w", err)
	}
	var serverData []*models.ServerData
	err = json.Unmarshal(body, &serverData)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling response: %w", err)
	}
	return &GetServersDetailResponse{
		ServerData: serverData,
	}, nil
}

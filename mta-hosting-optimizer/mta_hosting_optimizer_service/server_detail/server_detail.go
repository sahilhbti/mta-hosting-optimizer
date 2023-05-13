package server_detail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service/models"
)

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
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:9517/", nil)
	if err != nil {
		return nil, fmt.Errorf("error in building request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error parser request: %w", err)
	}
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
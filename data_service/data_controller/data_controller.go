package data_controller

import "github.com/mta-hosting-optimizer/data_service/models"


//go:generate mockgen -source=$PWD/data_controller.go -destination=$PWD/mocks/data_controller.go -package=mocks
type IDataController interface {
	GetServerData(*GetServerDatRequest) (*GetServerDataResponse, error)
}

type GetServerDatRequest struct {
}

type GetServerDataResponse struct {
	ServerData []*models.ServerData
}

type DataController struct {
}

func NewDataController() *DataController {
	return &DataController{}
}

func (m *DataController) GetServerData(request *GetServerDatRequest) (*GetServerDataResponse, error) {
	return &GetServerDataResponse{
		ServerData: []*models.ServerData{
			{
				Ip:       "127.0.0.1",
				HostName: "mta-prod-1",
				Active:   true,
			},
			{
				Ip:       "127.0.0.2",
				HostName: "mta-prod-1",
				Active:   false,
			},
			{
				Ip:       "127.0.0.3",
				HostName: "mta-prod-2",
				Active:   true,
			},
			{
				Ip:       "127.0.0.4",
				HostName: "mta-prod-2",
				Active:   true,
			},
			{
				Ip:       "127.0.0.5",
				HostName: "mta-prod-2",
				Active:   false,
			},
			{
				Ip:       "127.0.0.6",
				HostName: "mta-prod-3",
				Active:   false,
			},
		},
	}, nil
}

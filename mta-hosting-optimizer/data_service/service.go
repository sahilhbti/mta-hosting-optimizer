package data_service

import (
	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/data_service/data_controller"
	"github.com/mta-hosting-optimizer/data_service/models"
)

type DataService struct {
	dataController data_controller.IDataController
}

func NewDataService() *DataService {
	dataController := data_controller.NewDataController()
	return &DataService{
		dataController: dataController,
	}

}

func (service *DataService) GetServerDetails(ctx *gin.Context) ([]*models.ServerData, error) {
	serverDataResp, err := service.dataController.GetServerData(&data_controller.GetServerDatRequest{})
	if err != nil {
		return nil, err
	}
	return serverDataResp.ServerData, err
}

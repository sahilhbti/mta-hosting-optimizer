package mta_hosting_optimizer_service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/models"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/mta_optimizer_controller"
)

type MtaHostingOptimizerService struct {
	mtaOptimizerController mta_optimizer_controller.IMtaOptimizerController
}

func NewMtaHostingOptimizerService(httpClient *http.Client, url string) *MtaHostingOptimizerService {
	mtaController := mta_optimizer_controller.NewMtaOptimizerController(httpClient, url)
	return &MtaHostingOptimizerService{
		mtaOptimizerController: mtaController,
	}
}

func (service *MtaHostingOptimizerService) GetUnderUtilizedHostName(ctx *gin.Context) (*models.GetUnderUtilizedHostNameResponse, error) {
	resp, err := service.mtaOptimizerController.GetUnderUtilizedHost(ctx, &mta_optimizer_controller.GetUnderUtilizedHostRequest{})
	if err != nil {
		return nil, err
	}
	return &models.GetUnderUtilizedHostNameResponse{
		HostNames: resp.HostNames,
	}, nil
}

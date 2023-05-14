//go:build unit_test && !integration

package mta_hosting_optimizer_service

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/models"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/mta_optimizer_controller"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/mta_optimizer_controller/mocks"
)

type mockGetUnderUtilizedHost struct {
	resp *mta_optimizer_controller.GetUnderUtilizedHostResponse
	err  error
}

func TestMtaHostingOptimizerService_GetUnderUtilizedHostName(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type fields struct {
		mtaOptimizerController mta_optimizer_controller.IMtaOptimizerController
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name                     string
		fields                   fields
		args                     args
		want                     *models.GetUnderUtilizedHostNameResponse
		mockGetUnderUtilizedHost *mockGetUnderUtilizedHost
		wantErr                  bool
	}{
		{
			name: "error in getting response",
			args: args{
				ctx: c,
			},
			mockGetUnderUtilizedHost: &mockGetUnderUtilizedHost{
				err: errors.New("error in getting response"),
			},
			wantErr: true,
		},
		{
			name: "got response",
			args: args{
				ctx: c,
			},
			mockGetUnderUtilizedHost: &mockGetUnderUtilizedHost{
				resp: &mta_optimizer_controller.GetUnderUtilizedHostResponse{
					HostNames: []string{
						"host-name-1",
					},
				},
			},
			want: &models.GetUnderUtilizedHostNameResponse{
				HostNames: []string{
					"host-name-1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockMtaOptimizerController := mocks.NewMockIMtaOptimizerController(ctrl)
			if tt.mockGetUnderUtilizedHost != nil {
				mockMtaOptimizerController.EXPECT().GetUnderUtilizedHost(gomock.Any(), gomock.Any()).Return(tt.mockGetUnderUtilizedHost.resp, tt.mockGetUnderUtilizedHost.err)
			}
			service := &MtaHostingOptimizerService{
				mtaOptimizerController: mockMtaOptimizerController,
			}
			got, err := service.GetUnderUtilizedHostName(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnderUtilizedHostName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnderUtilizedHostName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMtaHostingOptimizerService(t *testing.T) {
	type args struct {
		httpClient *http.Client
		url        string
	}
	tests := []struct {
		name string
		args args
		want *MtaHostingOptimizerService
	}{
		{
			name: "test",
			args: args{
				httpClient: nil,
				url:        "test-url",
			},
			want: &MtaHostingOptimizerService{
				mtaOptimizerController: mta_optimizer_controller.NewMtaOptimizerController(nil, "test-url"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMtaHostingOptimizerService(tt.args.httpClient, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMtaHostingOptimizerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

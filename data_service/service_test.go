package data_service

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/mta-hosting-optimizer/data_service/data_controller"
	"github.com/mta-hosting-optimizer/data_service/data_controller/mocks"
	"github.com/mta-hosting-optimizer/data_service/models"
)

type mockGetServerData struct {
	resp *data_controller.GetServerDataResponse
	err  error
}

func TestDataService_GetServerDetails(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type fields struct {
		dataController data_controller.IDataController
	}
	fmt.Println("x")
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		want              []*models.ServerData
		wantErr           bool
		mockGetServerData *mockGetServerData
	}{
		{
			name: "error in getting server detail",
			args: args{
				ctx: c,
			},
			wantErr: true,
			mockGetServerData: &mockGetServerData{
				err: errors.New("error in getting response"),
			},
		},
		{
			name: "got server detail",
			args: args{
				ctx: c,
			},
			mockGetServerData: &mockGetServerData{
				resp: &data_controller.GetServerDataResponse{
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
					},
				},
			},
			want: []*models.ServerData{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDataController := mocks.NewMockIDataController(ctrl)
			if tt.mockGetServerData != nil {
				mockDataController.EXPECT().GetServerData(gomock.Any()).Return(tt.mockGetServerData.resp, tt.mockGetServerData.err)
			}
			service := &DataService{
				dataController: mockDataController,
			}
			got, err := service.GetServerDetails(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServerDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

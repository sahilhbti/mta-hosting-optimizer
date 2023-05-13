package mta_optimizer_controller

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"github.com/mta-hosting-optimizer/data_service/models"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail"
	"github.com/mta-hosting-optimizer/mta_hosting_optimizer_service/server_detail/mocks"
)

type mockGetServersDetail struct {
	resp *server_detail.GetServersDetailResponse
	err  error
}

func TestMtaOptimizerController_GetUnderUtilizedHost(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	err := os.Setenv("X", "1")
	if err != nil {
		return
	}
	type fields struct {
		HttpClient    *http.Client
		ServersDetail server_detail.IServersDetail
	}
	type args struct {
		ctx     *gin.Context
		request *GetUnderUtilizedHostRequest
	}
	fmt.Println("x")
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		want                 *GetUnderUtilizedHostResponse
		mockGetServersDetail *mockGetServersDetail
		wantErr              bool
	}{
		{
			name: "error in getting server detail",
			args: args{
				ctx: c,
			},
			mockGetServersDetail: &mockGetServersDetail{
				err: errors.New("error in getting response"),
			},
			wantErr: true,
		},
		{
			name: "got underutilized server",
			args: args{
				ctx: c,
			},
			mockGetServersDetail: &mockGetServersDetail{
				resp: &server_detail.GetServersDetailResponse{
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
				},
			},
			want: &GetUnderUtilizedHostResponse{
				HostNames: []string{
					"mta-prod-1", "mta-prod-3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			serverDetailMock := mocks.NewMockIServersDetail(ctrl)
			m := &MtaOptimizerController{
				HttpClient:    tt.fields.HttpClient,
				ServersDetail: serverDetailMock,
			}
			if tt.mockGetServersDetail != nil {
				serverDetailMock.EXPECT().GetServersDetail(gomock.Any(), gomock.Any()).Return(tt.mockGetServersDetail.resp, tt.mockGetServersDetail.err)
			}
			got, err := m.GetUnderUtilizedHost(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnderUtilizedHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnderUtilizedHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:build unit_test && !integration

package server_detail

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/h2non/gock.v1"

	"github.com/mta-hosting-optimizer/data_service/models"
)

type mockServerDetail struct {
	ResponseStatus int
	Response       []*models.ServerData
}

func TestServersDetail_GetServersDetail(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type fields struct {
		HttpClient *http.Client
	}
	type args struct {
		ctx     *gin.Context
		request *GetServersDetailRequest
	}
	var tests = []struct {
		name             string
		fields           fields
		args             args
		want             *GetServersDetailResponse
		mockServerDetail *mockServerDetail
		wantErr          bool
	}{
		{
			name: "unable to make http request",
			mockServerDetail: &mockServerDetail{
				ResponseStatus: 504,
				Response:       nil,
			},
			args: args{
				ctx: c,
			},
			fields: fields{
				HttpClient: &http.Client{},
			},
			wantErr: true,
		},
		{
			name: "http request successfull",
			mockServerDetail: &mockServerDetail{
				ResponseStatus: 200,
				Response: []*models.ServerData{
					{
						Ip:       "id-1",
						HostName: "host-name",
						Active:   true,
					},
				},
			},
			fields: fields{
				HttpClient: &http.Client{},
			},
			args: args{
				ctx: c,
			},
			want: &GetServersDetailResponse{
				ServerData: []*models.ServerData{
					{
						Ip:       "id-1",
						HostName: "host-name",
						Active:   true,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gock.New("http://localhost").
				Reply(tt.mockServerDetail.ResponseStatus).
				JSON(tt.mockServerDetail.Response)

			s := &ServersDetail{
				HttpClient: tt.fields.HttpClient,
				Url:        "http://localhost",
			}
			got, err := s.GetServersDetail(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServersDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServersDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

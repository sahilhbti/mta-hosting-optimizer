//go:build unit_test && !integration

package data_controller

import (
	"reflect"
	"testing"

	"github.com/mta-hosting-optimizer/data_service/models"
)

func TestDataController_GetServerData(t *testing.T) {
	type args struct {
		request *GetServerDatRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *GetServerDataResponse
		wantErr bool
	}{
		{
			name: "server data",
			args: args{},
			want: &GetServerDataResponse{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DataController{}
			got, err := m.GetServerData(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetServerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetServerData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDataController(t *testing.T) {
	tests := []struct {
		name string
		want *DataController
	}{
		{
			name: "test",
			want: &DataController{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataController() = %v, want %v", got, tt.want)
			}
		})
	}
}

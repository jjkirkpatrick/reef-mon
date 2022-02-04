package temperature

import (
	"reflect"
	"testing"

	"github.com/jjkikrpatrick/reef-mon/system/models"
)


func TestGet(t *testing.T) {
	type args struct {
		monitorConfig models.MonitorConfig
	}
	tests := []struct {
		name string
		args args
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.args.monitorConfig)
		})
	}
}

func Test_sensors(t *testing.T) {
	tests := []struct {
		name    string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sensors()
			if (err != nil) != tt.wantErr {
				t.Errorf("sensors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sensors() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_temperature(t *testing.T) {
	type args struct {
		sensor string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := temperature(tt.args.sensor)
			if (err != nil) != tt.wantErr {
				t.Errorf("temperature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("temperature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListTemperatureDevices(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"ListTemperatureDevices Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListTemperatureDevices()
		})
	}
}

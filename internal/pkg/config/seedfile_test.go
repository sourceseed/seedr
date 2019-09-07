package config

import (
	"reflect"
	"testing"
)

func TestParseSeedfile(t *testing.T) {
	tests := []struct {
		name    string
		filePath    string
		want    *Seedfile
		wantErr bool
	}{
		{
			name: "fail",
			filePath: "./test/testdata/config/Nonexisting",
			want: &Seedfile{},
			wantErr: true,
		},
		{
			name: "success",
			filePath: "../../../test/testdata/config/Sendfile.yml",
			want: &Seedfile{
				Name: "golang",
				Parameters: []ParamOptions{
					{
						Variable: "APPNAME",
						Description: "The name of the application",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSeedfile(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSeedfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSeedfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

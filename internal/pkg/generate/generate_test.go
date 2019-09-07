package generate

import (
	"reflect"
	"testing"
)

func Test_locateFiles(t *testing.T) {
	baseDir := "../../../test/testdata/generate"
	tests := []struct {
		name string
		dir string
		want []string
	}{
		{
			name: "success"	,
			dir: baseDir,
			want: []string{
				baseDir,
				baseDir+"/a",
				baseDir+"/folder",
				baseDir+"/folder/b",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := locateFiles(tt.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("locateFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

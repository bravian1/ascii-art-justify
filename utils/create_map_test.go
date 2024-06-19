package utils

import (
	"os"
	"reflect"
	"testing"
)

func Test_CreateMap(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		name string
		file string // test file
		args
		want    map[rune][]string
		wantErr bool // whether an error is expected
	}{
		{
			name:    "valid file with 95 characters",
			file:    "testdata/standard.txt",
			wantErr: false,
		},

		{
			name:    "empty file",
			file:    "testdata/empty.txt",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "file with less than 95 characters",
			file:    "testdata/fewcharacters.txt",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "File with More Than 95 Characters",
			file:    "testdata/manycharacters.txt",
			want:    nil,
			wantErr: true,
		},

		{
			name:    "file with characters with length less than 8",
			file:    "testdata/charlenght.txt",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMap(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

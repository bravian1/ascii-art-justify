package utils

import (
	"reflect"
	"testing"
)

func Test_CreateMap(t *testing.T) {
	tests := []struct {
		name     string
		filename string // test filename
		want     map[rune][]string
		wantErr  bool
	}{
		{
			name:     "valid filename with 95 characters",
			filename: "testdata/standard",
			wantErr:  false,
		},
		{
			name:     "empty filename",
			filename: "testdata/empty",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "filename with less than 95 characters",
			filename: "testdata/fewcharacters",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "File with More Than 95 Characters",
			filename: "testdata/manycharacters",
			want:     nil,
			wantErr:  true,
		},

		{
			name:     "filename with characters with length less than 8",
			filename: "testdata/charlenght",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMap(tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

package shufflings

import (
	"reflect"
	"sort"
	"testing"
)

func TestProcess(t *testing.T) {
	type args struct {
		data BodyData
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{
				data: BodyData{
					Data: "abc",
				},
			},
			want: []string{
				"abc",
				"acb",
				"bac",
				"bca",
				"cab",
				"cba",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.args.data)

			// sort data of array before assert
			sort.SliceStable(got, func(i, j int) bool {
				return got[i] < got[j]
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}

package odd

import (
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	type args struct {
		numberList []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				numberList: []int{
					1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9, 1,
				},
			},
			want: []int{
				1,
			},
		},
		{
			name: "success case 1",
			args: args{
				numberList: []int{
					7,
				},
			},
			want: []int{
				7,
			},
		},
		{
			name: "success case 2",
			args: args{
				numberList: []int{
					0,
				},
			},
			want: []int{
				0,
			},
		},
		{
			name: "success case 3",
			args: args{
				numberList: []int{
					1, 1, 2,
				},
			},
			want: []int{
				2,
			},
		},
		{
			name: "success case 4",
			args: args{
				numberList: []int{
					0, 1, 0, 1,
				},
			},
			want: []int{
				0,
			},
		},
		{
			name: "success case 5",
			args: args{
				numberList: []int{
					1,2,2,3,3,3,4,3,3,3,2,2,1,
				},
			},
			want: []int{
				4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Process(tt.args.numberList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}

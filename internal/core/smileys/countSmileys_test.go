package smileys

import "testing"

func TestProcess(t *testing.T) {
	type args struct {
		smileysList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				smileysList: []string{":"},
			},
			want: 0,
		},
		{
			name: "success case 1",
			args: args{
				smileysList: []string{":)", ";(", ";}", ":-D"},
			},
			want: 2,
		},
		{
			name: "success case 2",
			args: args{
				smileysList: []string{";D", ":-(", ":-)", ";~)"},
			},
			want: 3,
		},
		{
			name: "success case 3",
			args: args{
				smileysList: []string{"^-^",";^]", ";]", ":[", ";*", ":$", ";-D"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Process(tt.args.smileysList); got != tt.want {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}

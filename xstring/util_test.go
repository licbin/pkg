package xstring

import "testing"

func TestIsCharacterUnique(t *testing.T) {
	type args struct {
		astr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "unique",
			args: args{"abcdef"},
			want: true,
		},
		{
			name: "not unique",
			args: args{"abdda"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCharacterUnique(tt.args.astr); got != tt.want {
				t.Errorf("IsCharacterUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

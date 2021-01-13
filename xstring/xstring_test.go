package xstring

import "testing"

// func BuilderConcatV2(list ...string) string {
// 	b := strings.Builder{}

// 	for _, v := range list {
// 		b.Grow(len(v))
// 		b.WriteString(v)
// 	}

// 	return b.String()
// }
// BenchmarkBuilderConcatV2-8   	15968769	        73.4 ns/op	      24 B/op	       3 allocs/op

// BenchmarkBuilderConcat-8   	31433134	        34.3 ns/op	       8 B/op	       1 allocs/op
func BenchmarkBuilderConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuilderConcat("ni", "hh", "dafa")
	}
}

func TestBuilderConcat(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuilderConcat(tt.args.list...); got != tt.want {
				t.Errorf("BuilderConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}

package xstring

import "testing"

func TestRandomNumber(t *testing.T) {
	got := RandomNumber(10)
	if len(got) != 10 {
		t.Errorf("RandomNumber() = %v, want %v", len(got), 10)
	}
	if testing.Verbose() {
		t.Logf("got:%v\n", got)
	}
}

func BenchmarkRandomNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomNumber(20)
	}
}

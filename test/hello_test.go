package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "Καλημέρα κόσμε; or こんにちは 世界"

	if got != want {
		t.Errorf("got is %s, but want is %s", got, want)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}

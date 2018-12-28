package xrid

import (
	"net/http"
	"testing"
)

func BenchmarkNewWithContext(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = newContextWithID(&http.Request{})
	}
}

func BenchmarkNewID(b *testing.B) {
	b.SetBytes(26)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = newID()
	}
}

func BenchmarkNewXID(b *testing.B) {
	b.SetBytes(20)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = newXID()
	}
}

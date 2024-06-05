package main

import (
	"net/http"
	"testing"
)

func BenchmarkSpringBoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://127.0.0.1:8088/demo/v1/get-name")
	}
}

func BenchmarkGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://127.0.0.1:8089/demo/v1/get-name")
	}
}

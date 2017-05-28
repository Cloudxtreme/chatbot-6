package main

import "testing"

func TestDBSession(t *testing.T) {
	session := DBSession()
	if session == nil {
		t.Error("database session was not created")
	}
}

func benchmarkDBSession(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		DBSession()
	}
}

func BenchmarkDBSession1(b *testing.B)  { benchmarkDBSession(1, b) }
func BenchmarkDBSession2(b *testing.B)  { benchmarkDBSession(2, b) }
func BenchmarkDBSession3(b *testing.B)  { benchmarkDBSession(3, b) }
func BenchmarkDBSession10(b *testing.B) { benchmarkDBSession(10, b) }
func BenchmarkDBSession20(b *testing.B) { benchmarkDBSession(20, b) }
func BenchmarkDBSession40(b *testing.B) { benchmarkDBSession(40, b) }

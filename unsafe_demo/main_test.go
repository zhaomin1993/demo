package main

import "testing"

func BenchmarkStr2Bytes(b *testing.B) {
	var a = "sdjfaerjarlaerae4ra5er3safaera35er43fa5e1r3aef8er34faer4afr5ae313rf1ae3rfa4r1afe3r41a3fer1"
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes(a)
	}
}

func BenchmarkString(b *testing.B) {
	var bs = []byte{12, 12, 1, 31, 32, 32, 4, 32, 2, 3, 132, 4, 234, 34, 2, 3, 2, 31, 21, 3, 1, 32, 43, 3, 23, 1, 21, 2, 13, 2, 3, 121, 3}
	for i := 0; i < b.N; i++ {
		_ = String(bs)
	}
}

func BenchmarkBytes2String2(b *testing.B) {
	var a = "sdjfaerjarlaerae4ra5er3safaera35er43fa5e1r3aef8er34faer4afr5ae313rf1ae3rfa4r1afe3r41a3fer1"
	for i := 0; i < b.N; i++ {
		_ = []byte(a)
	}
}

func BenchmarkString2Bytes2(b *testing.B) {
	var bs = []byte{12, 12, 1, 31, 32, 32, 4, 32, 2, 3, 132, 4, 234, 34, 2, 3, 2, 31, 21, 3, 1, 32, 43, 3, 23, 1, 21, 2, 13, 2, 3, 121, 3}
	for i := 0; i < b.N; i++ {
		_ = string(bs)
	}
}

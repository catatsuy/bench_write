package main_test

import (
	"os"
	"testing"
)

func BenchmarkWriteEveryOpenFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteEveryOpenFile("testtest\n")
	}
}

func WriteEveryOpenFile(s string) {
	file, err := os.OpenFile("tmp", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(s)
}

func BenchmarkWriteOpeningFile(b *testing.B) {
	b.StopTimer()
	file, err := os.OpenFile("tmptmp", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		WriteOpeningFile(file, "testtest\n")
	}

	file.Close()
}

func WriteOpeningFile(f *os.File, s string) {
	f.WriteString(s)
}

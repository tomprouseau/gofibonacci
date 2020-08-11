package main

import (
	"math/big"
	"testing"
)

var fibTests = []struct {
	n        int    // input
	expected string // expected result
}{
	{1, "1"},
	{2, "1"},
	{3, "2"},
	{4, "3"},
	{5, "5"},
	{6, "8"},
	{7, "13"},
	{100, "354224848179261915075"},
	{300, "222232244629420445529739893461909967206666939096499764990979600"},
}

func TestSingleThreaded(t *testing.T) {
	for _, tt := range fibTests {
		actual := calcFibonacci(tt.n)
		var expectedResult big.Int
		expectedResult.SetString(tt.expected, 10)
		if actual.Cmp(&expectedResult) != 0 {
			t.Error("Calculated", actual.String(), "but expected", expectedResult.String())
		}
	}
}

func TestMultiThreaded(t *testing.T) {
	for _, tt := range fibTests {

		fibonacciChannel := make(chan big.Int)
		go calcFibonacciMultithread(tt.n, fibonacciChannel)
		actual := <-fibonacciChannel

		var expectedResult big.Int
		expectedResult.SetString(tt.expected, 10)
		if actual.Cmp(&expectedResult) != 0 {
			t.Error("Calculated", actual.String(), "but expected", expectedResult.String())
		}
	}
}

func BenchmarkSingleThreaded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range fibTests {
			calcFibonacci(tt.n)
		}
	}
}

func BenchmarkMultiThreaded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range fibTests {
			fibonacciChannel := make(chan big.Int)
			go calcFibonacciMultithread(tt.n, fibonacciChannel)
			<-fibonacciChannel
		}
	}
}

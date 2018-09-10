package utils

import (
	"fmt"
	"testing"
)

func TestMyAdd(t *testing.T) {
	tests := []struct{ a, b, c int32 }{
		{1, 2, 3},
		{4, 5, 9},
		//{4, 5, 19},
	}

	fmt.Println(tests)

	for _, test := range tests {
		if actual := MyAdd(test.a, test.b); actual != test.c {
			t.Errorf("%v, %d", test, actual)
		}
	}
}

func ExampleMyAdd() {
	fmt.Println(MyAdd(1, 2))
}

func BenchmarkMyAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MyAdd(1, 2)
	}
}

func BenchmarkMyString(b *testing.B) {
	s := "12;34;5;67;;8"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MyString(s)
	}
}

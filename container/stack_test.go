package container

import (
	"testing"
	"fmt"
)

func TestStack_Push(t *testing.T) {
	var s Stack
	s.Push(1)
	fmt.Println(s)
}
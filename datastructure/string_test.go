package datastructure

import (
	"fmt"
	"testing"
)

func TestSerialize(t *testing.T) {
	s := NewString()
	s.Set([]byte("1234567890"), []byte("1234567890"))
	// s.Set([]byte("2234567890"), []byte("2234567890"))
	// s.Delete([]byte("2234567890"))
	fmt.Printf("%+v\n", s.Serialize())
}

package datastructure

import (
	"fmt"
	"kiDB/utils"
	"testing"
)

func TestSerialize(t *testing.T) {
	s := NewString()
	for i := 0; i < 100000; i++ {
		s.Set([]byte(utils.RandomString(10)), []byte(utils.RandomString(20)))
	}

	// s.Set([]byte("2234567890"), []byte("2234567890"))
	// s.Delete([]byte("2234567890"))
	fmt.Printf("%+vMiB\n", len(s.Serialize())/1024.0/1024.0)
}

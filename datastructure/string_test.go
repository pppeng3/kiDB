package datastructure

import (
	"fmt"
	"kiDB/utils"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestToProtoBufString(t *testing.T) {
	for i := 0; i < 100000; i++ {
		Set(utils.RandomString(10), utils.RandomAlphaString(20))
	}

	s := ToProtoBufString()
	// fmt.Printf("%+v\n", s)
	b, _ := proto.Marshal(s)
	fmt.Printf("%+vKiB\n", len(b)/1024)
}
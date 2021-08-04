package consts

import (
	"fmt"
	"testing"
)

func TestCommand_Serialize(t *testing.T) {
	c := &Command{
		OperationType: Operation_Insert,
		DataType:      String,
		KeySize:       10,
		ValueSize:     10,
		Key:           []byte("1234567890"),
		Value:         []byte("1234567890"),
	}
	// fmt.Printf("%+v\n", string(c.Serialize()))

	// fmt.Printf("%+v\n", len(c.Serialize()))
	fmt.Println(DeSerialize(c.Serialize()))
}

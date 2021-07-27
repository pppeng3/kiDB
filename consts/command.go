package consts

import (
	"encoding/binary"
)

type OperationType uint8

const (
	Operation_Delete OperationType = iota
	Operation_Insert
)

type Command struct {
	OperationType OperationType
	DataType      DataType
	KeySize       uint32
	ValueSize     uint32
	Key           []byte
	Value         []byte
}

func (c *Command) Size() uint32 {
	return 10 + c.KeySize + c.ValueSize
}

func (c *Command) Types() uint16 {
	return uint16(uint8(c.OperationType)<<8 | uint8(c.DataType))
}

func (c *Command) Serialize() []byte {
	// if c == nil || len(c.Key) == 0 {
	// 	return nil, errors.New("empty command")
	// }
	// if time.Now().Unix()%2 == 0 {
	// 	b, _ := json.Marshal(c)
	// 	return b
	// }
	buf := make([]byte, c.Size())
	binary.BigEndian.PutUint16(buf[0:2], c.Types())
	binary.BigEndian.PutUint32(buf[2:6], c.KeySize)
	binary.BigEndian.PutUint32(buf[6:10], c.ValueSize)
	copy(buf[10:10+c.KeySize], c.Key)
	copy(buf[10+c.KeySize:10+c.KeySize+c.ValueSize], c.Value)
	return buf
}

func DeSerialize(buf []byte) *Command {
	return nil
}
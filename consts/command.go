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

//Size OperationType DataType and ValueSize(without Keysize)

func (c *Command) Types() uint16 {
	return uint16(uint8(c.OperationType)<<8 | uint8(c.DataType))
}

func (c *Command) Size() uint32 {
	return 6 + c.ValueSize
}

func (c *Command) Serialize() []byte {
	buf := make([]byte, c.Size())
	binary.BigEndian.PutUint16(buf[0:2], c.Types())
	binary.BigEndian.PutUint32(buf[2:6], c.ValueSize)
	copy(buf[6:6+c.ValueSize], c.Value)
	return buf
}
func DeSerialize(buf []byte) *Command {
	cmd := &Command{}
	types := binary.BigEndian.Uint16(buf[0:2])
	cmd.OperationType = OperationType(types >> 8)
	cmd.DataType = DataType(types & 0x00ff)
	cmd.ValueSize = binary.BigEndian.Uint32(buf[2:6])
	cmd.Value = make([]byte, cmd.ValueSize)
	copy(cmd.Value, buf[6:6+cmd.ValueSize])
	if int(cmd.ValueSize) != len(cmd.Value) {
		panic("value size mismatch")
	}
	return cmd
}

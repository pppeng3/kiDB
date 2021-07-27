package datastructure

import (
	"bytes"
	"encoding/binary"
	"kiDB/consts"
	"kiDB/engine"
)

type String struct {
	engine *engine.SkipList //string全存到一个跳表里
}

func NewString() *String {
	return &String{
		engine.NewDefaultSkipList(),
	}
}

func (s *String) Set(key, value []byte) bool {
	command := &consts.Command{
		OperationType: consts.Operation_Insert,
		DataType:      consts.String,
		KeySize:       uint32(len(key)),
		ValueSize:     uint32(len(value)),
		Key:           key,
		Value:         value,
	}

	s.engine.Set(key, command)
	return true
}

func (s *String) Get(key []byte) ([]byte, bool) { //todo
	node := s.engine.Get(key)
	if node == nil || !bytes.Equal(node.Key(), key) {
		return nil, false
	}

	return node.Value().Value, true
}

func (s *String) Delete(key []byte) bool { //todo
	node := s.engine.Get(key)
	if node == nil || !bytes.Equal(node.Key(), key) {
		return false
	}

	//直接修改
	command := node.Value()
	command.OperationType = consts.Operation_Delete
	command.Value = nil
	command.ValueSize = 0

	return true
}

func (s *String) Trivialize() (commands []*consts.Command, size uint32) { //todo
	commands = make([]*consts.Command, 0, s.engine.Size())
	var cmd *consts.Command
	for it := s.engine.Begin(); it != nil; it = it.Next() {
		cmd = it.Value()
		commands = append(commands, cmd)
		size += cmd.Size() + 4
	}
	return
}

func (s *String) Serialize() []byte {
	cmds, size := s.Trivialize()
	size += 4
	buf := make([]byte, size)
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(cmds)))
	offset := 4
	for _, c := range cmds {
		binary.BigEndian.PutUint32(buf[offset:offset+4], c.Size())
		offset += 4
		copy(buf[offset:offset+int(c.Size())], c.Serialize())
		offset += int(c.Size())
	}
	return buf
}

/*
LSM只有append操作
硬盘里有key1
内存中调用了Delete(key1), 如果Delete方法直接在内存中删除了key1,硬盘里的数据无法得到更新
如果硬盘里存的是command{opType:insert, dataType:string, key: key1}, 且调用Delete(key1)时插入一条command{opType:delete, dataType:string, key: key1},sstable合并的过程中会把硬盘中的key1删除
*/

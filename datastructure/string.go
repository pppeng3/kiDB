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

func (s *String) Trivialize() (commands []*consts.Command) { //todo
	commands = make([]*consts.Command, 0, s.engine.Size())
	var cmd *consts.Command
	for it := s.engine.Begin(); it != nil; it = it.Next() {
		cmd = it.Value()
		commands = append(commands, cmd)
	}
	return
}

func (s *String) Serialize() []byte {
	cmds := s.Trivialize()
	num := len(cmds)
	keys := make([][]byte, 0, num)
	keySizeSum := 0
	for _, cmd := range cmds {
		keys = append(keys, cmd.Key)
		keySizeSum += len(cmd.Key) + int(cmd.Size())
	}
	buf := make([]byte, 4+8*len(keys)+keySizeSum)
	offset := 0
	binary.BigEndian.PutUint32(buf[offset:offset+4], uint32(num))
	offset += 4
	preSumKey := uint32(0) //keySize的前缀和
	for i := 0; i < num; i++ {
		preSumKey += uint32(len(keys[i]))
		binary.BigEndian.PutUint32(buf[offset:offset+4], preSumKey)
		offset += 4
	}
	for i := 0; i < num; i++ {
		copy(buf[offset:offset+len(keys[i])], keys[i])
		offset += len(keys[i])
	}
	preSumCmd := uint32(0) //cmdSize的前缀和
	for i := 0; i < num; i++ {
		preSumCmd += uint32(cmds[i].Size())
		binary.BigEndian.PutUint32(buf[offset:offset+4], preSumCmd)
		offset += 4
	}
	for i := 0; i < num; i++ {
		copy(buf[offset:offset+int(cmds[i].Size())], cmds[i].Serialize())
		offset += int(cmds[i].Size())
	}
	return buf
}

//TODO: DeSerialize
/*
LSM只有append操作
硬盘里有key1
内存中调用了Delete(key1), 如果Delete方法直接在内存中删除了key1,硬盘里的数据无法得到更新
如果硬盘里存的是command{opType:insert, dataType:string, key: key1}, 且调用Delete(key1)时插入一条command{opType:delete, dataType:string, key: key1},sstable合并的过程中会把硬盘中的key1删除
*/

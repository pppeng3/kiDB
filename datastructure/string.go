package datastructure

import (
	"bytes"
	"kiDB/engine"
)

type String struct {
	engine *engine.SkipList //string全存到一个跳表里
}

func (s *String) Set(key, value []byte) bool {
	s.engine.Set(key, value)
	return true
}

func (s *String) Get(key []byte) ([]byte, bool) {
	node := s.engine.Get(key)
	if node == nil || bytes.Equal(node.Key(), key) {
		return nil, false
	}
	return node.Value(), true
}

func (s *String) Delete(key []byte) bool {
	return s.engine.Delete(key)
}

func (s *String) Trivialize() (keys, values [][]byte) {
	keys = make([][]byte, 0, s.engine.Size())
	values = make([][]byte, 0, s.engine.Size())

	for it := s.engine.Begin(); it != nil; it = it.Next() {
		keys = append(keys, it.Key())
		values = append(values, it.Value())
	}
	return
}

// func (s *String) Serialize() []byte {
// 	keys, values := s.Trivialize()
// 	// size := 0
// 	var k, v []byte
// 	for i, _ := range keys {
// 		k = keys[i]
// 		v = values[i]

// 		c := consts.Command{
// 		OperationType: inser,
// 		DataType:      0,
// 		KeySize:       0,
// 		ValueSize:     0,
// 		Key:           []byte{},
// 		Value:         []byte{},
// 	}
// 	}

// }
/*
LSM只有append操作
硬盘里有key1
内存中调用了Delete(key1), 如果Delete方法直接在内存中删除了key1,硬盘里的数据无法得到更新
如果硬盘里存的是command{opType:insert, dataType:string, key: key1}, 且调用Delete(key1)时插入一条command{opType:delete, dataType:string, key: key1},sstable合并的过程中会把硬盘中的key1删除
*/
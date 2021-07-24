package datastructure

import (
	"kiDB/engine"
	kiDB "kiDB/proto/consts"
	kiDB_string "kiDB/proto/string"
	"kiDB/utils"

	"google.golang.org/protobuf/proto"
)

var (
	Engine = engine.NewDefaultSkipList() //string全存到一个跳表里
)

func Set(key, value string) bool {
	cmd := &kiDB.Command{
		Operation: kiDB.Operation_insert,
		DataType:  kiDB.DataType_string,
		Key:       utils.Str2bytes(key),
		Value:     utils.Str2bytes(key),
	}
	b, _ := proto.Marshal(cmd)
	Engine.Set([]byte(key), b)
	return true
}

func Get(key string) (string, bool) {
	it := Engine.Get([]byte(key))
	if it == nil {
		// not found
		return "", false
	}
	cmd := &kiDB.Command{}
	proto.Unmarshal(it.Value(), cmd)
	if cmd.Operation == kiDB.Operation_delete {
		// found but has been deleted
		return "", false
	}

	return utils.Bytes2str(cmd.Value), true
}

// func Delete(key string) bool {
// 	//todo 原地修改减少内存分配
// 	Engine.Set([]byte(key), &kiDB.Command{
// 		Operation: kiDB.Operation_delete,
// 		Key:       []byte(key),
// 		Value:     nil,
// 	})
// 	return true
// }

func trivialize() [][]byte {
	res := make([][]byte, 0, Engine.Size())
	for it := Engine.Begin(); it != nil; it = it.Next() {
		res = append(res, it.Value())
	}
	return res
}

func ToProtoBufString() *kiDB_string.String {
	b := trivialize()
	return &kiDB_string.String{
		Length: int32(len(b)),
		Values: b,
	}
}

package engine

import (
	"kiDB/consts"
	"kiDB/utils"
	"testing"
)

func BenchmarkSkipListSet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), &consts.Command{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), &consts.Command{})
	}
}

func BenchmarkSkipListGet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomString(10)), &consts.Command{})
	}
	k := []byte(utils.RandomString(10))
	sl.Set(k, &consts.Command{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Get(k)
	}
}

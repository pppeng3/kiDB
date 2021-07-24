package engine

import (
	"bytes"
	"testing"

	"kiDB/utils"
)

func TestSkipList(t *testing.T) {
	sl := NewDefaultSkipList()
	a := make(map[string][]byte, 100000)
	for i := 0; i < 100000; i++ {
		a[utils.RandomString(10)] = []byte(utils.RandomString(10))
	}
	for k, v := range a {
		sl.Set([]byte(k), v)
	}
	for k, v := range a {
		if bytes.Equal(sl.Get([]byte(k)).Value(), v) {
			panic("failed")
		}
	}
	for i := 0; i < 100000; i++ {
		key := utils.RandomString(11)
		if sl.Get([]byte(key)) != nil {
			panic("failed")
		}
	}
}

func BenchmarkSkipListSet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), []byte{})
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), []byte{})
	}
}

func BenchmarkSkipListGet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomString(10)), []byte{})
	}
	k := []byte(utils.RandomString(10))
	sl.Set(k, []byte{})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Get(k)
	}
}

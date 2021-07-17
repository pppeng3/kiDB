package engine

import (
	"bytes"
	"math/rand"
)

type SkipList struct {
	maxLevel        int
	length          int
	probability     float64
	probabilityList []float64
	randomGenerator rand.Source
	head            []*node
	previousCache   []*node //每一行key比当前key小的node, 插入时使用
}

type node struct {
	key   []byte
	value interface{}
	next  *node
	pre   *node
	down  *node
}

func NewSkipList(maxLevel int, probability float64, randomGenerator rand.Source) *SkipList {
	if probability >= 1 {
		panic("invalid probability")
	}
	probabilityList := make([]float64, maxLevel)
	probabilityList[0] = 1
	for i := 1; i < maxLevel; i++ {
		probabilityList[i] = probabilityList[i-1] * probability
	}

	return &SkipList{
		maxLevel:        maxLevel,
		length:          0,
		probability:     probability,
		probabilityList: probabilityList,
		randomGenerator: randomGenerator,
		head:            make([]*node, maxLevel),
		previousCache:   make([]*node, maxLevel),
	}
}

func NewDefaultSkipList() *SkipList {
	return NewSkipList(20, float64(1)/3, rand.NewSource(rand.Int63()))
}

// randomLevel return level in [0, maxLevel)
func (sl *SkipList) randomLevel() (level int) {
	r := float64(sl.randomGenerator.Int63()) / (1 << 63)

	level = 1
	for level < sl.maxLevel && r < sl.probabilityList[level] {
		level++
	}
	return level - 1
}

func (sl *SkipList) previousNodes(key []byte) []*node {
	cache := sl.previousCache
	prev := sl.head[sl.maxLevel-1] //prev.key是当前level中小于key的最大的key
	now := prev.next
	for i := sl.maxLevel - 1; i >= 0; i-- {
		for now != nil && bytes.Compare(key, now.key) > 0 {
			prev = now
			now = now.next
		}
		cache[i] = prev
		prev = prev.down
		now = prev.next
	}
	return cache
}

func (sl *SkipList) Set(key []byte, value interface{}) error {
	level := sl.randomLevel()

	for i := 0; i < level; i++ {

	}

	return nil
}

func (sl *SkipList) Get(key []byte) (interface{}, error) {

	return nil, nil
}

func (sl *SkipList) GetRange(begin, end *[]byte) ([]interface{}, error) {
	return nil, nil
}

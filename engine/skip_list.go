package engine

import (
	"math"
	"math/rand"
)

type SkipList struct {
	maxLevel        int
	length          int
	probability     float64
	probabilityList []float64
	randomGenerator rand.Source64
}

func NewSkipList(maxLevel int, probability float64, randomGenerator rand.Source64) *SkipList {
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
	}

}

func (sl *SkipList) Set(key string, value interface{}) error {

	return nil
}

func (sl *SkipList) Get(key string) (interface{}, error) {

	return nil, nil
}

func (sl *SkipList) GetRange(begin, end string) ([]interface{}, error) {
	return nil, nil
}

func (t *SkipList) randomLevel() (level int) {
	r := float64(t.randomGenerator.Uint64()) / math.MaxUint64

	level = 1
	for level < t.maxLevel && r < t.probabilityList[level] {
		level++
	}
	return
}

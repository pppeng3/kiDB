package engine

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/Apale7/common/utils"
)

/*
head 1
head 1 2     6
head 1 2 4 5 6
*/
func TestSkipList_previousNodes(t *testing.T) {
	sl := NewSkipList(3, 0.5, rand.NewSource(time.Now().UnixNano()))
	sl.length = 5
	for i := 0; i < 3; i++ {
		sl.head[i] = &Node{}
	}
	sl.head[0].next = &Node{
		key:   []byte("1"),
		value: 1,
		next: &Node{
			key:   []byte("2"),
			value: 2,
			next: &Node{
				key:   []byte("4"),
				value: 4,
				next: &Node{
					key:   []byte("5"),
					value: 5,
					next: &Node{
						key:   []byte("6"),
						value: 6,
						next:  nil,
					},
				},
			},
		},
	}
	sl.head[1].next = &Node{
		key:   []byte("1"),
		value: 1,
		next: &Node{
			key:   []byte("2"),
			value: 2,
			next: &Node{
				key:   []byte("6"),
				value: 6,
				next:  nil,
				down:  sl.head[0].next.next.next.next.next,
			},
			down: sl.head[0].next.next,
		},
		down: sl.head[0].next,
	}
	sl.head[2].next = &Node{
		key:   []byte("1"),
		value: 1,
		next:  nil,
		down:  sl.head[1].next,
	}
	// fmt.Printf("%+v", sl.head)
	previous := sl.previousNodes([]byte("44"))
	values := make([]int, 0, 3)
	for _, p := range previous {
		values = append(values, p.value.(int))
	}
	if !reflect.DeepEqual(values, []int{4, 2, 1}) {
		t.Logf("want %+v, got %+v", []int{4, 2, 1}, values)
		t.FailNow()
	}
}

func TestSkipList(t *testing.T) {
	sl := NewDefaultSkipList()
	a := make(map[string]int, 100000)
	for i := 0; i < 100000; i++ {
		a[utils.RandomString(10)] = rand.Int()
	}
	for k, v := range a {
		sl.Set([]byte(k), v)
	}
	for k, v := range a {
		if sl.Get([]byte(k)).Value().(int) != v {
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
	sl := NewSkipList(30, 0.5, rand.NewSource(time.Now().UnixNano()))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), rand.Int())
	}
	b.StopTimer()
}

func BenchmarkSkipListGet(b *testing.B) {
	sl := NewDefaultSkipList()
	a := make(map[string]int, b.N)
	cnt := 100000
	for i := 0; i < cnt; i++ {
		a[utils.RandomString(10)] = rand.Int()
	}
	for k, v := range a {
		sl.Set([]byte(k), v)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sl.Get([]byte("123"))
	}
	b.StopTimer()
}

/*
head 1
head 1 2     6
head 1 2 4 5 6
*/
func TestGet(t *testing.T) {
	sl := NewSkipList(3, 0.5, rand.NewSource(time.Now().UnixNano()))
	sl.length = 5
	sl.head[0].next = &Node{
		key:   []byte("1"),
		value: 1,
		next: &Node{
			key:   []byte("2"),
			value: 2,
			next: &Node{
				key:   []byte("4"),
				value: 4,
				next: &Node{
					key:   []byte("5"),
					value: 5,
					next: &Node{
						key:   []byte("6"),
						value: 6,
						next:  nil,
					},
				},
			},
		},
	}
	sl.head[1].next = &Node{
		key:   []byte("1"),
		value: 1,
		next: &Node{
			key:   []byte("2"),
			value: 2,
			next: &Node{
				key:   []byte("6"),
				value: 6,
				next:  nil,
				down:  sl.head[0].next.next.next.next.next,
			},
			down: sl.head[0].next.next,
		},
		down: sl.head[0].next,
	}
	sl.head[2].next = &Node{
		key:   []byte("1"),
		value: 1,
		next:  nil,
		down:  sl.head[1].next,
	}
	fmt.Printf("sl.Get([]byte(\"1\")).Value(): %v\n", sl.Get([]byte("1")).Value())
	fmt.Printf("sl.Get([]byte(\"2\")).Value(): %v\n", sl.Get([]byte("2")).Value())
	fmt.Printf("sl.Get([]byte(\"4\")).Value(): %v\n", sl.Get([]byte("4")).Value())
	fmt.Printf("sl.Get([]byte(\"5\")).Value(): %v\n", sl.Get([]byte("5")).Value())
	fmt.Printf("sl.Get([]byte(\"6\")).Value(): %v\n", sl.Get([]byte("6")).Value())
	sl.Set([]byte("3"), 3)
	fmt.Printf("sl.Get([]byte(\"3\")).Value(): %v\n", sl.Get([]byte("3")).Value())

}

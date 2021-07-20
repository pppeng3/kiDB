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
	//for k, v := range a {
	//	if sl.Get([]byte(k)).Value().(int) != v {
	//		panic("failed")
	//	}
	//}
	//for i := 0; i < 100000; i++ {
	//	key := utils.RandomString(11)
	//	if sl.Get([]byte(key)) != nil {
	//		panic("failed")
	//	}
	//}
}

func BenchmarkSkipListSet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), 1)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sl.Set([]byte(utils.RandomAlphaString(10)), 1)
	}
}

func BenchmarkSkipListGet(b *testing.B) {
	b.StopTimer()
	sl := NewDefaultSkipList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		sl.Set([]byte(utils.RandomString(10)), struct{}{})
	}
	k := []byte(utils.RandomString(10))
	v := 1
	sl.Set(k, v)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Get(k)
	}
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
	//fmt.Printf("sl.Get([]byte(\"6\")).Value(): %v\n", sl.Get([]byte("6")).Value())
	sl.Set([]byte("3"), 3)
	fmt.Printf("sl.Get([]byte(\"3\")).Value(): %v\n", sl.Get([]byte("3")).Value())

}

/*
head 1
head 1 2     6
head 1 2 4 5 6
*/
func TestGetRange(t *testing.T) {
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

	fmt.Printf("sl.GetRange([]byte(\"0\"), []byte(\"3.3\")).Value():")
	a, b := sl.GetRange([]byte("0"), []byte("3.3"))
	fmt.Println(a.Value(), b.Value())

	fmt.Printf("sl.GetRange([]byte(\"1.1\"), []byte(\"5\")).Value():")
	c, d := sl.GetRange([]byte("1.1"), []byte("5"))
	fmt.Println(c.Value(), d.Value())

	fmt.Printf("sl.GetRange([]byte(\"6\"), []byte(\"7\")).Value():")
	e, f := sl.GetRange([]byte("6"), []byte("6"))
	fmt.Println(e.Value(), f)

	fmt.Printf("sl.GetRange([]byte(\"0.1\"), []byte(\"0.2\")).Value():")
	g, h := sl.GetRange([]byte("0.1"), []byte("0.2"))
	fmt.Println(g.Value(), h.Value())
}
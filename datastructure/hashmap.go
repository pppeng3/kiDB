package datastructure

import "bytes"

type Bucket struct {
	hash int
	next *Bucket
	length int
	head *MapNode
}

type Operation string

type MapNode struct {
	key []byte
	value []byte
	operation Operation
	next *MapNode
}

func NewHashMap(length int) []*Bucket {
	var BucketTable []*Bucket
	BucketTable[length - 1] = &Bucket{length + 1, nil,  0, nil}
	for i := length - 2; i > 0; i-- {
		BucketTable[i] = &Bucket{i, BucketTable[i+1], 0, nil}
	}
	return BucketTable
}

func (bk *Bucket) djbHash(key []byte) int {
	length := len(key)
	hash := 5
	for i := 0; i < length; i++ {
		hash = ((hash << 5) + hash) + int(key[i]) //hash(i) = hash(i-1) * 33 + str[i]
	}
	return hash
}

func (bk *Bucket) Set(key , value []byte) bool {
	BucketTable := NewHashMap(16)
	hash := bk.djbHash(key)
	index := hash % len(BucketTable)
	prev := &MapNode{}
	if BucketTable[index].head == nil {
		BucketTable[index].head = &MapNode{key: key, value: value, operation: "Insert"}
		BucketTable[index].length++
		return true
	}else {
		now := BucketTable[index].head
		for now != nil {
			if bytes.Equal(now.key, key){
				now.value = value
				return true
			}
			prev = now
			now = now.next
		}
		prev.next = &MapNode{key: key, value: value, operation: "Insert"}
		BucketTable[index].length++
	}
	return false
}

func (bk *Bucket) Get(key []byte) ([]byte, bool) {
	BucketTable := NewHashMap(16)
	hash := bk.djbHash(key)
	index := hash % len(BucketTable)
	now := BucketTable[index].head
	for now != nil {
		if bytes.Equal(now.key, key) {
			return now.value, true
		}
		now = now.next
	}
	return nil, false
}

//有点冗余，多一个prev就能直接get了删
func (bk *Bucket) Delete(key []byte) bool {
	BucketTable := NewHashMap(16)
	hash := bk.djbHash(key)
	index := hash % len(BucketTable)
	now := BucketTable[index].head
	prev := &MapNode{}
	if bytes.Equal(now.key, key) {
		BucketTable[index].head = now.next
	}
	for now != nil {
		prev = now
		now = now.next
		if bytes.Equal(now.key, key) {
			now.operation = "delete"
			prev.next = now.next
			now.next = nil
			return true
		}
	}
	return false
}

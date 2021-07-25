package engine

import "testing"
import "github.com/stretchr/testify/assert"

func TestAll(t *testing.T) {
	bs := NewBitSet(8)
	assert.Equal(t,len(bs.bytes), 1)
	bs.Set(3)
	assert.Equal(t,bs.Get(3), true)
}

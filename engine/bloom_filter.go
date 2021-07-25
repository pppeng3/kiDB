package engine

type bitSet struct {
	bytes []byte
	size  int
}

// NewBitSet n: the count of bits
func NewBitSet(size int) *bitSet {
	return &bitSet{bytes: make([]byte, (size-1)/8+1), size: size}
}

func (s *bitSet) Set(pos int) {
	idx, offset := pos>>3, pos&7
	s.bytes[idx] |= 1 << offset
}

func (s *bitSet) Get(pos int) bool {
	idx, offset := pos>>3, pos&7
	return s.bytes[idx]&(1<<offset) != 0
}
func (s *bitSet) Size() int { return s.size }

var seeds = [78]uint64{31, 73, 127, 179, 233, 283, 353, 419, 467, 547, 607, 661, 739, 811, 877, 947, 1019, 1063, 1087, 1129, 1153, 1217, 1229, 1289, 1297, 1367, 1381, 1447, 1453, 1499, 1523, 1579, 1597, 1637, 1663, 1723, 1741, 1801, 1823, 1879, 1901, 1979, 1993, 2039, 2063, 2113, 2131, 2207, 2221, 2281, 2293, 2351, 2371, 2417, 2437, 2521, 2539, 2609, 2621, 2683, 2689, 2731, 2749, 2803, 2833, 2897, 2909, 2971, 3001, 3067, 3083, 3169, 3187, 3253, 3259, 3329, 3343, 3407}

var hashFunctions []func([]byte) uint64

func initHashFunctions() { //todo 记得调用初始化函数
	for _, sd := range seeds {
		hashFunctions = append(hashFunctions, func(s []byte) (hash uint64) {
			for _, c := range s {
				hash *= sd
				hash += uint64(c)
			}
			return
		})
	}
}

type BloomFilter struct {
	bitSet
}

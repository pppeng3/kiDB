package consts

type DataType uint8

const (
	String DataType = iota
	List
	HashMap
	HashSet
	ZSet
)

// Serializable Each datastructure has to implement this interface for serializing to hard disk
type Serializable interface {
	Serialize() []byte
}

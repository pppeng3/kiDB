package datastructure

// Serializable Each datastructure has to implement this interface for serializing to hard disk
type Serializable interface {
	Serialize() []byte
}
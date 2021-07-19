package consts

type DataType uint8

const (
	String DataType = iota
	List
	HashMap
	HashSet
	ZSet
)

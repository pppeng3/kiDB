package engine

type Engine interface {
	Set(key []byte, value interface{}) error
	Get(key []byte) (interface{}, error)
	GetRange(begin, end *[]byte) ([]interface{}, error)
}

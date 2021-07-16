package engine

type Engine interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	GetRange(begin, end string) ([]interface{}, error)
}

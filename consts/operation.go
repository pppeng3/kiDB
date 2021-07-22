package consts

type Operation uint8

const (
	insert Operation = iota
)

// Command 所有对数据的操作都会抽象成command
// 例如: set name apale会抽象成插入一条Command{OperationType: insert, key: name, value: apale}
// 而del name会抽象成插入一条Command{OperationType: insert, key: name, value: apale}.
// 这样设计是为了实现LSM Tree
// All kinds of operations for data would be abstracted as a command.
// For example, set name apale would be insert as a Command{OperationType: insert, key: name, value: apale}.
// While del name would be insert as Command{OperationType: delete, key: name, value: }.
// This is designed for LSM Tree
type Command struct {
	OperationType Operation
	ValueType     DataType
	Key           []byte
	Value         []byte
}

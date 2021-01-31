package list

type List interface {
	Push(interface{})
	Pop() (interface{}, error)
	Unshift(interface{})
	Shift() (interface{}, error)
	Get(int) (interface{}, error)
	Set(int, interface{}) error
	Insert(int, interface{}) error
	Remove(int) error
}

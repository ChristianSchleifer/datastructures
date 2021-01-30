package list

type List interface {
	Push(int)
	Pop() (int, error)
	Unshift(int)
	Shift() (int, error)
	Get(int) (int, error)
	Set(int, int) error
	Insert(int, int) error
	Remove(int) error
}

package utils

// Const is a struct that can be set only once
type Const[T any] struct {
	value T
	set   bool
}

func NewConst[T any](v T) *Const[T] {
	return &Const[T]{value: v, set: false}
}

// Set sets the value of the Const
// It returns true if the value was set, false otherwise
func (c *Const[T]) Set(v T) bool {
	if c.set {
		return false
	}
	c.value = v
	c.set = true
	return true
}

func (c *Const[T]) Get() T {
	return c.value
}

func (c *Const[T]) IsSet() bool {
	return c.set
}

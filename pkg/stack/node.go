package stack

type node[T any] struct {
	payload T
	next    *node[T]
}

func newNode[T any](payload T) *node[T] {
	return &node[T]{
		payload: payload,
	}
}

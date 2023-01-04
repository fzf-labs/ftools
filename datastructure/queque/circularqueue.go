package queque

type CircularQueue[T any] struct {
	Data  []T
	Front int
}

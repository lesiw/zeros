package zeros

// Slice is a zero-valueable slice type whose Append method mutates
// in place and returns the updated slice.
type Slice[T any] []T

// Append appends values in place and returns the updated slice.
func (s *Slice[T]) Append(v ...T) Slice[T] {
	*s = append(*s, v...)
	return *s
}

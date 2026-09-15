package zeros

// Is reports whether t is the zero value for its type.
func Is[T comparable](t T) bool {
	var zero T
	return t == zero
}

package result

// Result is the Rustlike interface
type Result[T any] interface {
	IsOk() bool
	IsError() bool
	Unwrap() T
	UnwrapOr(T) T
	UnwrapOrDefault() T
	UnwrapErr() error
}

type result[T any] struct {
	value T
	err   error
}

// New creates a newly created result
func New[T any](value T, err error) result[T] {
	return result[T]{
		value: value,
		err:   err,
	}
}

// Unwrap returns the contained value
// Panics if err is not nil
func (r *result[T]) Unwrap() T {
	if r.IsError() {
		panic(r.err.Error())
	}
	return r.value
}

// UnwrapOr returns the contained value if it has one
// else it will return the given value
func (r *result[T]) UnwrapOr(value T) T {
	if r.IsError() {
		return value
	}
	return r.value
}

// UnwrapOrDefault returns the contained value if it has one
// else it will return the type's default zero value
func (r *result[T]) UnwrapOrDefault() T {
	if r.IsError() {
		var defaultValue T
		return defaultValue
	}
	return r.value
}

// UnwrapErr returns the contained error
// Panics if err is nil
func (r *result[T]) UnwrapErr() error {
	if !r.IsError() {
		panic("error is nil")
	}
	return r.err
}

// IsOk returns true if error is nil
func (r *result[T]) IsOk() bool {
	return r.err == nil
}

// IsError returns true if error is not nil
func (r *result[T]) IsError() bool {
	return r.err != nil
}

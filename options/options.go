package options

type Option[T any] func(target T) error
type OptionConstructor[T any, V any] func(value V) Option[T]
type OptionFunction[T any, V any] func(target T, value V) error

func New[T any, V any](function OptionFunction[T, V]) OptionConstructor[T, V] {
	return func(value V) Option[T] {
		return func(target T) error {
			return function(target, value)
		}
	}
}

func Apply[T any](target T, options []Option[T]) error {
	for _, option := range options {
		if err := option(target); err != nil {
			return err
		}
	}
	return nil
}

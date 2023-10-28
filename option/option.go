package option

type Option[T any] interface {
	IsSome() bool
	GetOrElse(func() T) T
	OrElse(func() Option[T]) Option[T]
}

func NewOption[T any](v *T) Option[T] {
	if v == nil {
		return None[T]()
	}
	return Some[T](v)
}

type Some_[T any] struct {
	Value *T
}

func Some[T any](v *T) Option[T] {
	return Some_[T]{Value: v}
}

type None_[T any] struct{}

func None[T any]() Option[T] {
	return None_[T]{}
}

func (s Some_[T]) IsSome() bool {
	return true
}

func (n None_[T]) IsSome() bool {
	return false
}

func (s Some_[T]) GetOrElse(f func() T) T {
	return *s.Value
}

func (n None_[T]) GetOrElse(f func() T) T {
	return f()
}

func (s Some_[T]) OrElse(f func() Option[T]) Option[T] {
	return s
}

func (n None_[T]) OrElse(f func() Option[T]) Option[T] {
	return f()
}

func Map[T, U any](f func(T) U) func(Option[T]) Option[U] {
	return func(o Option[T]) Option[U] {
		switch v := o.(type) {
		case *Some_[T]:
			tmp := f(*v.Value)
			return Some[U](&tmp)
		case *None_[T]:
			return None[U]()
		default:
			panic("unreachable")
		}
	}
}

func FlatMap[T, U any](f func(T) Option[U]) func(Option[T]) Option[U] {
	return func(o Option[T]) Option[U] {
		switch v := o.(type) {
		case *Some_[T]:
			return f(*v.Value)
		case *None_[T]:
			return None[U]()
		default:
			panic("unreachable")
		}
	}
}

package result

import o "github.com/MFQWKMR4/goutil/option"

type Result[T, E any] interface {
	IsOk() bool
	IsErr() bool
	Ok() o.Option[T]
	Err() o.Option[E]
}

func NewResult[T, E any](v *T, e *E) Result[T, E] {
	if v == nil {
		return Err[T, E](e)
	}
	return Ok[T, E](v)
}

type Ok_[T, E any] struct {
	Value *T
}

func Ok[T, E any](v *T) Result[T, E] {
	return Ok_[T, E]{Value: v}
}

type Err_[T, E any] struct {
	Value *E
}

func Err[T, E any](e *E) Result[T, E] {
	return Err_[T, E]{Value: e}
}

func (o Ok_[T, E]) IsOk() bool {
	return true
}

func (e Err_[T, E]) IsOk() bool {
	return false
}

func (o Ok_[T, E]) IsErr() bool {
	return false
}

func (e Err_[T, E]) IsErr() bool {
	return true
}

func (ok Ok_[T, E]) Ok() o.Option[T] {
	return o.Some[T](ok.Value)
}

func (e Err_[T, E]) Ok() o.Option[T] {
	return o.None[T]()
}

func (ok Ok_[T, E]) Err() o.Option[E] {
	return o.None[E]()
}

func (e Err_[T, E]) Err() o.Option[E] {
	return o.Some[E](e.Value)
}

func Map[T, U, E any](f func(T) U) func(Result[T, E]) Result[U, E] {
	return func(r Result[T, E]) Result[U, E] {
		switch v := r.(type) {
		case *Ok_[T, E]:
			tmp := f(*v.Value)
			return Ok[U, E](&tmp)
		case *Err_[T, E]:
			return Err[U, E](v.Value)
		}
		return nil
	}
}

func FlatMap[T, U, E any](f func(T) Result[U, E]) func(Result[T, E]) Result[U, E] {
	return func(r Result[T, E]) Result[U, E] {
		switch v := r.(type) {
		case *Ok_[T, E]:
			return f(*v.Value)
		case *Err_[T, E]:
			return Err[U, E](v.Value)
		}
		return nil
	}
}

package utilfx

import (
	"reflect"

	"go.uber.org/fx"
)

// Grouped provides constructors whose results belong to the same Fx value
// group. Nil constructors are ignored.
func Grouped(group string, ctors ...any) fx.Option {
	return provide(ctors, func(ctor any) any {
		return fx.Annotate(ctor, fx.ResultTags(GroupTag(group)))
	})
}

// GroupedFor provides constructors in a group composed from group and name.
// Nil constructors are ignored.
func GroupedFor(name, group string, ctors ...any) fx.Option {
	return Grouped(GroupName(group, name), ctors...)
}

// GroupedAs provides constructors as T in the same Fx value group. Each
// constructor result must implement T. Nil constructors are ignored.
func GroupedAs[T any](group string, ctors ...any) fx.Option {
	return provide(ctors, func(ctor any) any {
		return fx.Annotate(
			ctor,
			fx.As(new(T)),
			fx.ResultTags(GroupTag(group)),
		)
	})
}

func provide(ctors []any, annotate func(any) any) fx.Option {
	annotated := make([]any, 0, len(ctors))
	for _, ctor := range ctors {
		if isNil(ctor) {
			continue
		}

		annotated = append(annotated, annotate(ctor))
	}

	if len(annotated) == 0 {
		return fx.Options()
	}

	return fx.Provide(annotated...)
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

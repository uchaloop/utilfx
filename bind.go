package utilfx

import (
	"fmt"
	"reflect"
)

// Bind exposes implementation as contract. It can be passed directly to
// fx.Provide and replaces adapters of the form:
//
//	func(service *Service) Handler { return service }
//
// An incompatible implementation is reported as an Fx constructor error.
func Bind[Contract any, Impl any](impl Impl) (Contract, error) {
	v, ok := any(impl).(Contract)
	if !ok {
		var zero Contract

		return zero, fmt.Errorf(
			"%T does not implement %v",
			impl,
			reflect.TypeFor[Contract](),
		)
	}

	return v, nil
}

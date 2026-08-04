package utilfx

import "fmt"

// NameTag returns an Fx tag for a named dependency or result.
func NameTag(name string) string {
	return fmt.Sprintf(`name:"%s"`, name)
}

// OptionalNameTag returns an Fx tag for an optional named dependency.
func OptionalNameTag(name string) string {
	return fmt.Sprintf(`name:"%s" optional:"true"`, name)
}

// GroupTag returns an Fx tag for a value group.
func GroupTag(group string) string {
	return fmt.Sprintf(`group:"%s"`, group)
}

// GroupName returns group when name is empty. Otherwise it returns a group
// name in the form "group:name".
func GroupName(group, name string) string {
	if len(name) == 0 {
		return group
	}

	return group + ":" + name
}

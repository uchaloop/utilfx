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

// GroupName returns the group a named instance feeds: group itself when name is
// empty - the single default instance - and "group:name" otherwise. It is the
// rule GroupedFor applies, exposed for a module that needs the name without
// providing anything into it.
func GroupName(group, name string) string {
	if len(name) == 0 {
		return group
	}

	return group + ":" + name
}

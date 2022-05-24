package reflectx

import "strings"

func TagValueAndFlags(tag string) (string, map[string]bool) {
	values := strings.Split(tag, ",")
	flags := make(map[string]bool)
	if len(values[0]) > 1 {
		for _, flag := range values[1:] {
			flags[flag] = true
		}
	}
	return values[0], flags
}

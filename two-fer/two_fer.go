// Package twofer for retrieve name and output friendly message
package twofer

import "fmt"

// ShareWith will retrieve name as string and return some line
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

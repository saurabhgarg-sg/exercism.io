// twofer implements second excercise based on theatrical cabling device

package twofer

import "fmt"

// ShareWith implements the name replacement for this excercise
func ShareWith(name string) string {
	if len(name) == 0{
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
	 
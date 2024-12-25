package internal_example

import "github.com/Prashant20nov2003/internal_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}

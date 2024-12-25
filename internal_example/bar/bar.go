package internal_package_example

import "github.com/Prashant20nov2003/internal_example/foo/internal"

func FailUseDoubler(i int) int {
	return internal.Doubler(i)
}

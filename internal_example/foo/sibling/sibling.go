package sibling

import "github.com/Prashant20nov2003/internal_example/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}

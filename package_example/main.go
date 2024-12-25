package main

import (
	"fmt"

	format "github.com/Prashant20nov2003/package_example/do-format"
	"github.com/Prashant20nov2003/package_example/math"
)

func main() {
	num := math.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}

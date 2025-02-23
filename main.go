package main

import (
	"fmt"

	"github.com/shddtr/my-math/double"
	"github.com/shddtr/my-math/formatter"
)

func main() {
	result := double.Double(10)
	fmt.Println(formatter.Format(fmt.Sprintf("%d", result)))
}

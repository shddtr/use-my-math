package main

import (
	"fmt"

	"github.com/shddtr/my-math/v2/pkg/doubler"
	"github.com/shddtr/my-math/v2/pkg/formatter"
)

func main() {
	result := doubler.Double(10)
	fmt.Println(formatter.Format(fmt.Sprintf("%d", result)))
}

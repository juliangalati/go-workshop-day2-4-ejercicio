package main

import (
	"fmt"
)

type Pesos float64

func (p Pesos) String() string {
	return fmt.Sprintf("$%G", p)
}

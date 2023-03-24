package main

import (
	"github.com/jadenHsiao/poscom/src"
	"github.com/jadenHsiao/poscom/src/kernels"
)

func main() {
	a := new(kernels.Gainscha)
	a.A()
	src.ArgsType2String(1, 2, 3)
}

package main

import (
	"factory/impl"
)

func main() {
	factory := impl.NewIDCardFactory()

	idcard1 := factory.Create("Matisse")
	idcard1.Use()

	idcard2 := factory.Create("Sadko")
	idcard2.Use()
}

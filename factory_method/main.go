package main

import (
	"factory/impl"
)

func main() {
	idCardFactory := impl.NewIDCardFactory()

	idcard1 := idCardFactory.Create("Matisse")
	idcard1.Use()

	idcard2 := idCardFactory.Create("Sadko")
	idcard2.Use()
}

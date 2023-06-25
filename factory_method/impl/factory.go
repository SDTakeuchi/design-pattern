package impl

import (
	"fmt"

	"factory/domain"
)

type idCardFactory struct {
	*domain.AbstractFactory
}

var _ domain.Factory = &idCardFactory{}

// 本の内容に従ってPrivateにしている
func (f *idCardFactory) createProduct(owner string) domain.Product {
	return &IDCard{owner: owner}
}

// 本の内容に従って同じくPrivateにしている
func (f *idCardFactory) registerProduct(p domain.Product) {
	fmt.Printf("id card registered for %s\n", p.Owner())
}


func NewIDCardFactory() domain.Factory {
	f := &idCardFactory{
		// 抽象の構造体に具象の構造体に定義したメソッドを注入
		// これをやらないと&idCardFactory.Create()の時にnil pointerエラーになる
		AbstractFactory: domain.NewAbstractFactory(
			(&idCardFactory{}).createProduct,
			(&idCardFactory{}).registerProduct,
		),
	}
	return domain.Factory(f)
}

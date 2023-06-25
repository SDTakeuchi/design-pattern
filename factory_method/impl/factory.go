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

func (f *idCardFactory) registerProduct(p domain.Product) {
	fmt.Printf("card registered for %s\n", p.Owner())
}


func NewIDCardFactory() domain.Factory {
	f := &idCardFactory{
		// 抽象の構造体に具象の構造体に定義したメソッドを注入（注入っていうんか？）
		// これをやらないと&idCardFactory.Create()の時にnil pointerエラーになる
		//
		// 抽象クラスに対してメソッドを渡しているので、(&idCardFactory{}).createProductではなく、createProductだけを定義するのでもいい
		// ただし、今後このパッケージが拡張された時に、インスタンス化されなくてもcreateProductが実行可能になってしまうこともあり得るため、
		// 意図しない利用を避けたり、コードの読みやすさという理由で構造体にメソッドを紐づけている
		AbstractFactory: domain.NewAbstractFactory(
			(&idCardFactory{}).createProduct,
			(&idCardFactory{}).registerProduct,
		),
	}
	// FactoryにキャストすることでPolymorphismを実現できる
	return domain.Factory(f)
}

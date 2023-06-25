package domain

type (
	Product interface {
		Owner() string
		Use()
	}
	Factory interface {
		Create(owner string) Product
	}
	AbstractFactory struct {
		// AbstractFactoryのメソッドを直接叩かれたくないのでPrivate
		// 叩くなら具象クラスの方を叩いてほしい
		createProduct   func(owner string) Product
		registerProduct func(p Product)
	}
)

func NewAbstractFactory(
	createProduct func(owner string) Product,
	registerProduct func(p Product),
) *AbstractFactory {
	return &AbstractFactory{
		createProduct:   createProduct,
		registerProduct: registerProduct,
	}
}

func (f *AbstractFactory) Create(owner string) Product {
	p := f.createProduct(owner)
	f.registerProduct(p)
	return p
}

package impl

import (
	"factory/domain"
	"fmt"
)

type IDCard struct {
	// Privateにしておく。インスタンス生成後に変更されたくない
	owner string
}

// IDCard構造体がProductを満たしていることを担保する
var _ domain.Product = &IDCard{}

func (c *IDCard) Owner() string {
	return c.owner
}

func (c *IDCard) Use() {
	fmt.Printf("id card used: owner name: %s\n", c.Owner())
}

package impl

import (
	"factory/domain"
	"fmt"
)

type IDCard struct {
	// Privateにしておく。インスタンス生成後に変更されたくないため
	owner string
}

// 実際にはこの1文を書くことを忘れてしまわないように運用を明文化しないといけない気がする
// これがあるだけでIDCard構造体がProductを満たしていることを明示しているのできちっと書くようにしておきたい
var _ domain.Product = &IDCard{}

func (c *IDCard) Owner() string {
	return c.owner
}

func (c *IDCard) Use() {
	fmt.Printf("card used: owner name: %s\n", c.Owner())
}

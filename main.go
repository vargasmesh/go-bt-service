package main

import (
	"fmt"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

func main() {
	bt := tree.New()
	bt.Insert(10).Insert(5).Insert(15).Insert(3).Insert(7).Insert(13).Insert(17)
	tree.InOrder(bt.Root, func(n *tree.Node) {
		fmt.Println(n.Value)
	})
}

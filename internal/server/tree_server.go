package server

import (
	"context"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

type TreeServer struct {
	Tree       *tree.Tree
	insertChan chan int
}

func NewTreeServer() *TreeServer {
	return &TreeServer{
		Tree:       tree.New(),
		insertChan: make(chan int),
	}
}

func (s *TreeServer) Insert(value int) {
	s.insertChan <- value
}

func (s *TreeServer) Run(ctx context.Context) {
	for {
		select {
		case value := <-s.insertChan:
			s.Tree.Insert(value)
		case <-ctx.Done():
			return
		}
	}
}

func (s *TreeServer) GetPreOrderTree() []int {
	var flatTree []int
	tree.PreOrder(s.Tree.Root, func(n *tree.Node) {
		flatTree = append(flatTree, n.Value)
	})
	return flatTree
}

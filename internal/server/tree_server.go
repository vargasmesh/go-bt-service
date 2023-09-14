package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

type TreeServer struct {
	Tree        *tree.Tree
	insertChan  chan int
	maxTreeSize int
	currentSize int
}

func NewTreeServer() *TreeServer {
	t := &TreeServer{
		Tree:        tree.New(),
		insertChan:  make(chan int),
		maxTreeSize: 5,
		currentSize: 0,
	}

	_ = os.MkdirAll("./data", os.ModePerm)

	return t
}

func (s *TreeServer) Insert(value int) {
	s.insertChan <- value
}

func (s *TreeServer) Run(ctx context.Context) {
	for {
		select {
		case value := <-s.insertChan:
			s.handleInsert(value)
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

func (s *TreeServer) handleInsert(value int) {
	s.Tree.Insert(value)
	s.currentSize++
	if s.currentSize == s.maxTreeSize {
		f, err := os.Create(fmt.Sprintf("./data/%d.data", time.Now().Unix()))
		if err != nil {
			return
		}
		defer f.Close()

		tree.InOrder(s.Tree.Root, func(n *tree.Node) {
			f.WriteString(fmt.Sprintf("%d\n", n.Value))
		})
		s.Tree = tree.New()
		s.currentSize = 0
	}
}

package server

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

type TreeServer[T any] struct {
	Tree        *tree.Tree[T]
	less        tree.Less[T]
	insertChan  chan T
	maxTreeSize int
	currentSize int
	mu          sync.RWMutex
}

func NewTreeServer[T any](less tree.Less[T]) *TreeServer[T] {
	t := &TreeServer[T]{
		Tree:        tree.New(less),
		insertChan:  make(chan T),
		maxTreeSize: 100,
		currentSize: 0,
	}

	_ = os.MkdirAll("./data", os.ModePerm)

	return t
}

func (s *TreeServer[T]) Insert(value T) {
	s.insertChan <- value
}

func (s *TreeServer[T]) Run(ctx context.Context) {
	for {
		select {
		case value := <-s.insertChan:
			s.handleInsert(value)
		case <-ctx.Done():
			return
		}
	}
}

func (s *TreeServer[T]) GetPreOrderTree() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var flatTree []T
	tree.PreOrder(s.Tree.Root, func(n *tree.Node[T]) {
		flatTree = append(flatTree, n.Value)
	})
	return flatTree
}

func (s *TreeServer[T]) handleInsert(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Tree.Insert(value)
	s.currentSize++
	if s.currentSize == s.maxTreeSize {
		f, err := os.Create(fmt.Sprintf("./data/%d.data", time.Now().Unix()))
		if err != nil {
			return
		}
		defer f.Close()

		tree.InOrder(s.Tree.Root, func(n *tree.Node[T]) {
			f.WriteString(fmt.Sprintf("%d\n", n.Value))
		})
		s.Tree = tree.New(s.less)
		s.currentSize = 0
	}
}

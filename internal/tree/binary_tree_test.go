package tree_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

func TestTree(t *testing.T) {
	binaryTree := tree.New[int](
		func(t1, t2 int) bool {
			return t1 < t2
		},
	)
	for i := 0; i < 500; i++ {
		binaryTree.Insert(rand.Intn(1000))
	}

	result := []int{}
	tree.InOrder(binaryTree.Root, func(n *tree.Node[int]) {
		result = append(result, n.Value)
	})

	if !sort.SliceIsSorted(result, func(i, j int) bool {
		return result[i] < result[j]
	}) {
		t.Errorf("Tree is not sorted: %v", result)
	}
}

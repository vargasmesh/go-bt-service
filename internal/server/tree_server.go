package server

import (
	"context"
	"fmt"

	"github.com/vargasmesh/go-bt-service/internal/tree"
)

type TreeServer struct {
	Tree *tree.Tree
}

func NewTreeServer() *TreeServer {
	return &TreeServer{
		Tree: tree.New(),
	}
}

func (s *TreeServer) Run(ctx context.Context) {
	fmt.Println("Hello from Tree Server")
	<-ctx.Done()
}

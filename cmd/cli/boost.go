package main

import (
	"github.com/americanas-go/boost/internal/pkg/spec"
)

func main() {
	_, err := spec.ParseDir("/Users/joao.faria/Projetos/github.com/americanas-go/boost/examples/simple")
	if err != nil {
		panic(err)
	}

}

package main

import "github.com/americanas-go/log/contrib/rs/zerolog.v1"

func main() {

	zerolog.NewLogger()

	_, err := ParseDir("/Users/joao.faria/Projetos/github.com/americanas-go/boost/examples/simple")
	if err != nil {
		panic(err)
	}
}

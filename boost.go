package main

import (
	"fmt"

	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
	"gopkg.in/yaml.v3"
)

func main() {

	zerolog.NewLogger()

	spec, err := ParseDir("/Users/joao.faria/Projetos/github.com/americanas-go/boost/examples/simple")
	if err != nil {
		log.Error(err.Error())
	}

	j, _ := yaml.Marshal(spec)
	fmt.Println(string(j))
}

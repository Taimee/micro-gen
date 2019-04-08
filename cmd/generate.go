package cmd

import (
	"flag"
	"log"

	"github.com/Taimee/micro-gen/lib/generator"
	"github.com/micro/cli"
)

func Generate(c *cli.Context) {
	flag.Parse()
	args := flag.Args()
	name := args[0]
	action := args[1]
	err := generator.Handler(name, action)
	if err != nil {
		log.Fatal(err)
	}

	err = generator.Interactor(name, action)
	if err != nil {
		log.Fatal(err)
	}

	err = generator.Presenter(name, action)
	if err != nil {
		log.Fatal(err)
	}

	err = generator.Usecase(name, action)
	if err != nil {
		log.Fatal(err)
	}
}

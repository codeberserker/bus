package main

import (
	"log"

	"github.com/codeberserker/bus/pkg/monitor"
)

func main() {
	config, err := monitor.NewConfig()
	if err != nil {
		panic(err)
	}

	m := &monitor.Monitor{Config: config}
	out, err := m.Run()
	if err != nil {
		panic(err)
	}

	for e := range out {
		log.Println(e)
	}
}

package main

import (
	"fmt"
	"github.com/lestoni/go-logtick/pkg/parser"
	"io/ioutil"
)

func main() {
	log, err := ioutil.ReadFile("testdata/git.log")
	if err != nil {
		panic(err)
	}

	content := fmt.Sprintf("%s", log)

	output, err := logtick.Parse(content)
	if err != nil {
		panic(err)
	}

	out, err := output.ToJSON()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", output)
	fmt.Println(out)
}

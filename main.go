package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/emicklei/proto"
)

var (
	memo = map[string]struct{}{}
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	if err := walk(filename); err != nil {
		log.Fatal(err)
	}
}

func walk(filename string) (err error) {
	if _, ok := memo[filename]; ok {
		return nil
	}
	memo[filename] = struct{}{}

	fmt.Println(filename)

	reader, err := os.Open(filename)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return nil
		}
		return err
	}

	defer func() {
		err = reader.Close()
	}()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		return err
	}

	proto.Walk(definition, proto.WithImport(handleImport))

	return nil
}

func handleImport(s *proto.Import) {
	if err := walk(s.Filename); err != nil {
		log.Fatal(err)
	}
}

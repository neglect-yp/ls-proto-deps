package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/emicklei/proto"
)

var (
	baseDir = flag.String("baseDir", "", "base directory of proto files")

	memo = map[string]struct{}{}
)

func main() {
	flag.Parse()
	filenames := flag.Args()

	for _, filename := range filenames {
		if err := walk(filename); err != nil {
			log.Fatal(err)
		}
	}
}

func walk(filename string) (err error) {
	if _, ok := memo[filename]; ok {
		return nil
	}
	memo[filename] = struct{}{}

	reader, err := os.Open(path.Join(*baseDir, filename))
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

func handleImport(i *proto.Import) {
	fmt.Println(i.Filename)

	if err := walk(i.Filename); err != nil {
		log.Fatal(err)
	}
}

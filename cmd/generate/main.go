package main

import (
	"github.com/tyndyll/nidc"
	"github.com/tyndyll/nidc/language"
	"log"
)

func main() {
	tree, err := language.Parse(`MATCH (m:MOVIE {name: "Buried"}) RETURN m.year, m.rating`)
	if err != nil {
		log.Fatalf("Could not parse error: %s", err.Error())
	}

	log.Println(tree.GetText())

	query, err := nidc.NewQuery(tree)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(query.SQL())
}

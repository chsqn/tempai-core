package main

import (
	"fmt"
	"log"

	"github.com/dnovikoff/tempai-core/compact"
	"github.com/dnovikoff/tempai-core/shanten"
)

func main() {
	generator := compact.NewTileGenerator()
	tiles, err := generator.CompactFromString("789m4466678p234s")
	if err != nil {
		log.Fatal(err)
	}
	results := shanten.CalculateTempai(tiles, nil)
	fmt.Printf("Hand is %s\n", tiles.Instances())
	fmt.Printf("Waits are %s\n", results.Waits().Tiles())
}

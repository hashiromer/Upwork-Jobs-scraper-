package main

import (
	"scrapers/upwork"
)

func main() {

	p := upwork.InitPipeline()
	err := p.Run("Shopify")
	if err != nil {
		panic(err)
	}

}

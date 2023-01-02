package main

import (
	"scrapers/upwork"
)

func main() {

	p := upwork.InitPipeline()
	err := p.Run("pdf")
	if err != nil {
		panic(err)
	}

}

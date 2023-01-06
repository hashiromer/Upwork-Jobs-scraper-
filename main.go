package main

import "scrapers/upwork"

func main() {

	p := upwork.InitPipeline()
	err := p.Run("")
	if err != nil {
		panic(err)
	}

}

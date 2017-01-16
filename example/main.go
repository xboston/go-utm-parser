package main

import (
	"log"

	utm "github.com/xboston/go-utm-parser"
)

func main() {

	url := "http://www.example.com/?utm_source=source&utm_campaign=campaign&utm_medium=medium&utm_content=content&utm_term=term&utm_extra1=extra1&utm_extra2=extra2"
	utm, _ := utm.ParseURL(&url)

	log.Println(utm)
}

package main

import (
	"os"

	"github.com/wailovet/easycgo"
)

func as(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	pathname := os.Args[1]
	easycgo.CheckErrorWithLDD(pathname)
}

package main

import (
	"fmt"
	"log"

	kbbi "github.com/dimassfeb-09/kbbi/kbbi"
)

func main() {
	if result, err := kbbi.Search("baju"); err != nil {
		log.Println(err)
	} else {
		fmt.Println(result)
	}
}

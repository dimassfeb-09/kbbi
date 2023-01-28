package main

import (
	"fmt"
	kbbi "github.com/dimassfeb-09/kbbi.git/kbbi"
	"log"
)

func main() {
	if result, err := kbbi.Search("baju"); err != nil {
		log.Println(err)
	} else {
		fmt.Println(result)
	}
}

package main

import (
	"fmt"

	"github.com/AravindTadi/cryptit/decrypt"
	"github.com/AravindTadi/cryptit/encrypt"
)

func main() {
	d := encrypt.Encmsg("Hi this is Aravind")
	fmt.Println(d)

	k := decrypt.Decmsg(d)
	fmt.Println(k)

}

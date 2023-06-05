package main

import (
	"github.com/valthern/paquetes_go/greet"
	"fmt"
	"rsc.io/quote"
)

func main()  {
	fmt.Println(greet.Italian())
	fmt.Println(greet.Russian())
	fmt.Println(quote.Hello())
}

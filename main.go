package main

import (
	"github.com/valthern/paquetes_go/greet"
	"fmt"
	"rsc.io/quote/v3"
)

func main()  {
	fmt.Println(greet.Spanish())
	fmt.Println(greet.Russian())
	fmt.Println(quote.HelloV3())
	fmt.Println(quote.Concurrency())
}

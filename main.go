package main

import (
	"github.com/valthern/paquetes_go/greet"
	"fmt"
	"rsc.io/quote"
	quotev3 "rsc.io/quote/v3"
)

func main()  {
	fmt.Println(greet.Italian())
	fmt.Println(greet.Russian())
	fmt.Println(quote.Hello())
	fmt.Println(quotev3.Concurrency())
	fmt.Println(quotev3.HelloV3())
}

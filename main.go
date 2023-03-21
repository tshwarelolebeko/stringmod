package main

import (
	"fmt"

	qr "github.com/tshwarelolebeko/stringmod/quotes"
	str "github.com/tshwarelolebeko/stringmod/strings"
)

func main() {

	e, o := str.CountOddEven("227374")
	fmt.Println(e, o)

	fmt.Println(qr.GetEmoji("turtle"))

}

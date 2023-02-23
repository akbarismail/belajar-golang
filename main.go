package main

import (
	"fmt"
	"math"

	mathakbar "github.com/akbarismail/belajargolang/math"
	"github.com/akbarismail/belajargolang/person"
	"github.com/bojanz/currency"
)

func main() {
	fmt.Println(mathakbar.Divide(10, 2))
	fmt.Println(math.Abs(-10))
	fmt.Println(mathakbar.Multiple(2, 4))

	budi := person.Person{Name: "Budi"}
	fmt.Println(budi)

	amount, err := currency.NewAmount("120", "USD")
	if err != nil {
		panic(err)
	}

	formatter := currency.NewFormatter(currency.NewLocale("en_US"))
	fmt.Println(formatter.Format(amount))

	dollarToIdr, err := amount.Convert("IDR", "14000")
	if err != nil {
		panic(err)
	}
	fmt.Println(dollarToIdr)
}

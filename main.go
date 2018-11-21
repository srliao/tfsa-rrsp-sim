package main

import (
	"log"

	"github.com/shopspring/decimal"
)

func main() {

	log.Println(tax(decimal.New(100000, 0)))
}

func tax(income decimal.Decimal) decimal.Decimal {
	//
	return decimal.New(0, 0)
}

func fedtax(ti decimal.Decimal) decimal.Decimal {

}

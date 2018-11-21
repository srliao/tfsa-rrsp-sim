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

	//rates
	r := []decimal.Decimal{
		decimal.New(150, -3),
		decimal.New(205, -3),
		decimal.New(250, -3),
		decimal.New(290, -3),
		decimal.New(330, -3),
	}
	//brackets
	b := []decimal.Decimal{
		decimal.New(0, 0),
		decimal.New(46605, 0),
		decimal.New(93208, 0),
		decimal.New(144489, 0),
		decimal.New(205842, 0),
	}

	tax := decimal.New(0, 0)

	for i := 0; i < (len(b) - 1); i++ {
		if ti.Cmp(b[i]) != -1 {
			inc := decimal.Min(ti.Sub(b[i]), b[i+1].Sub(b[i]))
			tax = tax.Add(inc.Mul(r[i]))
		}
	}

	if ti.Cmp(b[len(b)-1]) != -1 {
		tax = tax.Add(ti.Sub(b[len(b)-1]).Mul(r[len(r)-1]))
	}

	return tax
}

func provtax(ti decimal.Decimal) decimal.Decimal {

	//rates
	r := []decimal.Decimal{
		decimal.New(505, -4),
		decimal.New(915, -4),
		decimal.New(1116, -4),
		decimal.New(1216, -4),
		decimal.New(1316, -4),
	}
	//brackets
	b := []decimal.Decimal{
		decimal.New(0, 0),
		decimal.New(42960, 0),
		decimal.New(85923, 0),
		decimal.New(150000, 0),
		decimal.New(220000, 0),
	}

	tax := decimal.New(0, 0)

	for i := 0; i < (len(b) - 1); i++ {
		if ti.Cmp(b[i]) != -1 {
			inc := decimal.Min(ti.Sub(b[i]), b[i+1].Sub(b[i]))
			tax = tax.Add(inc.Mul(r[i]))
		}
	}

	if ti.Cmp(b[len(b)-1]) != -1 {
		tax = tax.Add(ti.Sub(b[len(b)-1]).Mul(r[len(r)-1]))
	}

	return tax
}

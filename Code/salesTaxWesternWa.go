package main

import (
	"fmt"
	"log"
	"os"
)

func logging() {
	l, _ := os.OpenFile("salesTax.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(l)
}

func calcSeattleSalesTax(r float64) float64 {
	s := .1035
	var se *float64 = &s
	log.Println("Mem address of Seattle is:", se)
	return s * r
}

func calcKingCountySalesTax(r float64) float64 {
	k := .1010
	var ki = &k
	log.Println("Mem address of King County is:", *ki)
	return k * r
}

func calcEverettSalesTax(r float64) float64 {
	e := .0990
	ev := &e
	log.Println("Mem address of Everett is:", ev)
	return e * r
}

func calcSnahomishCountySalesTax(r float64) float64 {
	s := .1035
	var sn *float64 = &s
	log.Println("Mem address of Snahomish County is:", sn)
	return s * r
}

func calcTacomaSalesTax(r float64) float64 {
	t := .1030
	var ta = &t
	log.Println("Mem address of Tacoma is:", ta)
	return t * r
}

func calcPierceCountySalesTax(r float64) float64 {
	p := .1025
	pi := &p
	log.Println("Mem address of Pierce county is:", pi)
	return p * r
}

func calcOlympiaSalesTax(r float64) float64 {
	o := .0985
	var ol *float64 = &o
	log.Println("Mem address of Olympia is:", ol)
	return o * r
}

func calcKitsapCountySalesTax(r float64) float64 {
	k := .0925
	var kit = &k
	log.Println("Mem address of Kitsap county is:", kit)
	return k * r
}

func thankYou() {
	fmt.Print("Thank you for your input. Please follow this app for additional features as they are added. \n")
}
func printCost(cost float64, calcTax func(r float64) float64) {
	taxCollected := calcTax(cost)
	fmt.Print("The tax on your purchase is $", taxCollected, "\n")
}

func taxFunction(location int) func(r float64) float64 {
	locationTax := map[int]func(r float64) float64{
		1: calcSeattleSalesTax,
		2: calcKingCountySalesTax,
		3: calcEverettSalesTax,
		4: calcSnahomishCountySalesTax,
		5: calcTacomaSalesTax,
		6: calcPierceCountySalesTax,
		7: calcOlympiaSalesTax,
		8: calcKitsapCountySalesTax,
	}
	return locationTax[location]
}

func main() {
	var cost float64
	var location int

	logging()
	defer thankYou()

	fmt.Print("Please enter the cost of your purchase, in dollars and cents: ")
	fmt.Scanf("%f\n", &cost)
	fmt.Printf("Please enter the location you want to calculate sales tax for on a purchase \n 1 - Seattle \n 2 - King \n 3 - Everett \n 4 - Snahomish \n 5 - Tacoma \n 6 - Pierce \n 7 - Olympia \n 8 - Kitsap: ")
	fmt.Scanf("%d\n", &location)

	if location <= 8 {
		printCost(cost, taxFunction(location))
	} else {
		fmt.Println("Not a valid location, please select one of the provided six choices.")
	}

}

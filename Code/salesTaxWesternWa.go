package main

import "fmt"

func calcSeattleSalesTax(r float64) float64 {
	return .1035 * r
}

func calcKingCountySalesTax(r float64) float64 {
	return .1010 * r
}

func calcEverettSalesTax(r float64) float64 {
	return .0990 * r
}

func calcSnahomishCountySalesTax(r float64) float64 {
	return .1035 * r
}

func calcTacomaSalesTax(r float64) float64 {
	return .1030 * r
}

func calcPierceCountySalesTax(r float64) float64 {
	return .1025 * r
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
	}
	return locationTax[location]
}

func main() {
	var cost float64
	var location int

	defer thankYou()

	fmt.Print("Please enter the cost of your purchase, in dollars and cents: ")
	fmt.Scanf("%f\n", &cost)
	fmt.Printf("Please enter the location you want to calculate sales tax for on a purchase \n 1 - Seattle \n 2 - King \n 3 - Everett \n 4 - Snahomish \n 5- Tacoma \n 6 - Pierce: ")
	fmt.Scanf("%d\n", &location)

	if location <= 6 {
		printCost(cost, taxFunction(location))
	} else {
		fmt.Println("Not a valid location, please select one of the provided six choices.")
	}

}

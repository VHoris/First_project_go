package main

import (
	"fmt"
	"math"
)

const ApplePrice = 5.99
const PearPrice = 7.
const AllMоnеу = 32.

func main() {
	SumPriceAP := ApplePrice*9 + PearPrice*8
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? - ", SumPriceAP, "грн.")

	SumPriceP := (AllMоnеу - math.Mod(AllMоnеу, PearPrice)) / PearPrice
	fmt.Println("2. Скільки груш ми можемо купити? - ", SumPriceP)

	SumPriceA := (AllMоnеу - math.Mod(AllMоnеу, ApplePrice)) / ApplePrice
	fmt.Println("3. Скільки яблук ми можемо купити? - ", SumPriceA)

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - ", 2*(ApplePrice+PearPrice) <= AllMоnеу)
}

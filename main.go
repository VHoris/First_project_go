package main

import (
	"fmt"
	"math"
)

const applePrice = 5.99
const pearPrice = 7.
const allMоnеу = 32.

func main() {
	
	applesToBuy := 9
	pearsToBuy := 8
	
	sumPriceAP := applePrice*float64(applesToBuy) + pearPrice*float64(pearsToBuy)
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? - ", sumPriceAP, "грн.")

	sumPriceP := (allMоnеу - math.Mod(allMоnеу, pearPrice)) / pearPrice
	fmt.Println("2. Скільки груш ми можемо купити? - ", sumPriceP)

	sumPriceA := (allMоnеу - math.Mod(allMоnеу, applePrice)) / applePrice
	fmt.Println("3. Скільки яблук ми можемо купити? - ", sumPriceA)
	
	applesToBuy = 2
	pearsToBuy = 2


	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - ", (float64(applesToBuy)*applePrice+float64(pearsToBuy)*pearPrice) <= allMоnеу)
}

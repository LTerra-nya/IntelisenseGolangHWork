package main

import (
	"fmt"
)

var (
	applePrice float32 = 5.99
	pearPrice  float32 = 7
	money      float32 = 23
	sum        float32 = (applePrice * 2) + (pearPrice * 2)
)

func main() {
	fmt.Println("В нас є", money, "грн.", "яблука коштують", applePrice, "а груші коштують", pearPrice)
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Printf("%.2f", (applePrice*9)+(pearPrice*8))
	fmt.Println(" Грн")
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Println(int(money / pearPrice))
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Println(int(money / applePrice))
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	if sum <= money {
		fmt.Println("Так")
	} else {
		fmt.Println("Ні")
	}
}

package gobasics

import "fmt"

func Functions() {
	printAllFromZeroToNumber(10)
	printSmthIfZeroElseHello(100)
	printAllStrings([]string{
		"A", "B", "C",
	})
}

func printAllFromZeroToNumber(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
}

func printSmthIfZeroElseHello(number int) {
	if number == 0 {
		fmt.Print("Smth")
	} else {
		fmt.Print("Hello")
	}
}

func printAllStrings(strings []string) {
	for _, str := range strings {
		fmt.Print(str)
	}
}

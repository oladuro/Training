package main

import "fmt"

func main() {
	for i := 0; i <= 5000; i++ {
		fmt.Printf("%d \t %b \t %0#X \t %q	\n", i, i, i, i)
	}
}

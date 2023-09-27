package main

import "fmt"

type Base struct {

}

func (b *Base) purchase(cost int) {
	fmt.Printf("Purchasing using base account for %d\n", cost)
}
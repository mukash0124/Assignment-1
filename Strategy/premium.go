package main

import "fmt"

type Premium struct {

}

func (p *Premium) purchase(cost int) {
	fmt.Printf("Purchasing using premium account for %f\n", float32(cost) * 0.8)
}